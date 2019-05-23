package mobile

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/segmentio/ksuid"
	"github.com/textileio/go-textile/core"
	"github.com/textileio/go-textile/keypair"
	"github.com/textileio/go-textile/pb"
)

var cafesTestVars = struct {
	cafePath    string
	cafe        *core.Textile
	cafeApiPort string
	mobilePath  string
	mobile      *Mobile
}{
	cafePath:    "./testdata/.textile3",
	cafeApiPort: "5000",
	mobilePath:  "./testdata/.textile4",
}

func TestMobile_SetupCafes(t *testing.T) {
	var err error
	cafesTestVars.mobile, err = createAndStartMobile(
		cafesTestVars.mobilePath, false, &testHandler{}, &testMessenger{})
	if err != nil {
		t.Fatal(err)
	}

	// start a cafe
	_ = os.RemoveAll(cafesTestVars.cafePath)
	err = core.InitRepo(core.InitConfig{
		Account:     keypair.Random(),
		RepoPath:    cafesTestVars.cafePath,
		CafeApiAddr: "0.0.0.0:" + cafesTestVars.cafeApiPort,
		CafeOpen:    true,
	})
	if err != nil {
		t.Fatal(err)
	}
	cafesTestVars.cafe, err = core.NewTextile(core.RunConfig{
		RepoPath: cafesTestVars.cafePath,
		Debug:    true,
	})
	if err != nil {
		t.Fatal(err)
	}
	err = cafesTestVars.cafe.Start()
	if err != nil {
		t.Fatal(err)
	}

	<-cafesTestVars.mobile.OnlineCh()
	<-cafesTestVars.cafe.OnlineCh()
}

func TestMobile_RegisterCafe(t *testing.T) {
	// create a token
	token, err := cafesTestVars.cafe.CreateCafeToken("", true)
	if err != nil {
		t.Fatal(err)
	}

	// register with cafe
	url := "http://127.0.0.1:" + cafesTestVars.cafeApiPort
	err = cafesTestVars.mobile.RegisterCafe(url, token)
	if err != nil {
		t.Fatal(err)
	}

	// add some data
	err = addTestData(cafesTestVars.mobile)
	if err != nil {
		t.Fatal(err)
	}

	//flush := func() error {
	//	count, err := flushCafeRequest(10)
	//	if err != nil {
	//		return err
	//	}
	//	if count > 0 {
	//		return flush()
	//	}
	//}

	_, err = flushCafeRequest(100)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func addTestData(m *Mobile) error {
	thrd, err := addTestThread(m, &pb.AddThreadConfig{
		Key:  ksuid.New().String(),
		Name: "test",
		Schema: &pb.AddThreadConfig_Schema{
			Preset: pb.AddThreadConfig_Schema_MEDIA,
		},
		Type:    pb.Thread_PRIVATE,
		Sharing: pb.Thread_INVITE_ONLY,
	})
	if err != nil {
		return err
	}

	_, err = m.addFiles([]string{"../mill/testdata/image.jpeg"}, thrd.Id, "hi")
	if err != nil {
		return err
	}

	_, err = m.AddMessage(thrd.Id, "hi")
	if err != nil {
		return err
	}

	hash, err := m.addFiles([]string{"../mill/testdata/image.png"}, thrd.Id, "hi")
	if err != nil {
		return err
	}

	_, err = m.AddComment(hash.B58String(), "nice")
	if err != nil {
		return err
	}

	hash, err = m.addFiles([]string{"../mill/testdata/image.jpeg", "../mill/testdata/image.png"}, thrd.Id, "hi")
	if err != nil {
		return err
	}

	_, err = m.AddLike(hash.B58String())
	if err != nil {
		return err
	}

	_, err = m.AddMessage(thrd.Id, "bye")
	if err != nil {
		return err
	}

	return nil
}

/*
Handle the request queue.
  1. List some groups
  2. List get the HTTP request list for each of those groups
  3. Handle them (set to pending, send to cafe)
  4. Delete failed (reties not handled here)
  5. Set successful to complete
*/
func flushCafeRequest(limit int) (int, error) {
	var count int
	res, err := cafesTestVars.mobile.CafeRequests(limit)
	if err != nil {
		return count, err
	}
	groups := new(pb.Strings)
	err = proto.Unmarshal(res, groups)
	if err != nil {
		return count, err
	}
	count = len(groups.Values)

	for _, g := range groups.Values {
		printGroupStatus(g)
	}

	// write the req for each group
	reqs := make(map[string]*pb.CafeHTTPRequest)
	for _, g := range groups.Values {
		res, err = cafesTestVars.mobile.WriteCafeHTTPRequest(g)
		if err != nil {
			return count, err
		}
		req := new(pb.CafeHTTPRequest)
		err = proto.Unmarshal(res, req)
		if err != nil {
			return count, err
		}
		reqs[g] = req
	}

	// mark each as pending (new loops for clarity)
	for g := range reqs {
		err = cafesTestVars.mobile.SetCafeRequestPending(g)
		if err != nil {
			return count, err
		}
	}

	// handle each
	for g, req := range reqs {
		res, err := handleReq(req)
		if err != nil {
			return count, err
		}
		if res.StatusCode >= 300 {
			fmt.Printf("got bad status: %d\n", res.StatusCode)
			err = cafesTestVars.mobile.SetCafeRequestFailed(g)
		} else {
			err = cafesTestVars.mobile.SetCafeRequestComplete(g)
		}
		if err != nil {
			return count, err
		}
		res.Body.Close()
		printGroupStatus(g)
	}
	return count, nil
}

func printGroupStatus(group string) {
	res, err := cafesTestVars.mobile.CafeRequestGroupStatus(group)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	status := new(pb.CafeRequestSyncGroupStatus)
	err = proto.Unmarshal(res, status)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(">>> " + group)
	fmt.Println(fmt.Sprintf("num. pending: %d", status.NumPending))
	fmt.Println(fmt.Sprintf("num. complete: %d", status.NumComplete))
	fmt.Println(fmt.Sprintf("num. total: %d", status.NumTotal))
	fmt.Println(fmt.Sprintf("size pending: %d", status.SizePending))
	fmt.Println(fmt.Sprintf("size complete: %d", status.SizeComplete))
	fmt.Println(fmt.Sprintf("size total: %d", status.SizeTotal))
	fmt.Println("<<<")
}

func handleReq(r *pb.CafeHTTPRequest) (*http.Response, error) {
	f, err := os.Open(r.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	req, err := http.NewRequest(r.Type.String(), r.Url, f)
	if err != nil {
		return nil, err
	}
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	return client.Do(req)
}

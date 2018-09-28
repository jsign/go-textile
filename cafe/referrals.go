package cafe

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/textileio/textile-go/cafe/dao"
	"github.com/textileio/textile-go/cafe/models"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func (c *Cafe) createReferral(g *gin.Context) {
	// cheap way to lock down this endpoint
	if c.ReferralKey != g.GetHeader("X-Referral-Key") {
		g.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	// how many should we make?
	count := 1
	// how many times should we be able to use them
	limit := 1
	// simple reference to who requested the new code
	requestedBy := ""
	params := g.Request.URL.Query()
	if params["count"] != nil && len(params["count"]) > 0 {
		tmp, err := strconv.ParseInt(params["count"][0], 10, 64)
		if err == nil {
			count = int(tmp)
		}
	}
	if params["limit"] != nil {
		tmp, err := strconv.ParseInt(params["limit"][0], 10, 64)
		if err == nil {
			limit = int(tmp)
		}
	}
	if params["requested_by"] != nil {
		requestedBy = params["requested_by"][0]
	}

	// hodl 'em
	codes := make([]string, count)
	for i := range codes {
		code, err := createReferralModel(limit, requestedBy)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		codes[i] = code
	}

	// ship it
	g.JSON(http.StatusCreated, models.ReferralResponse{
		RefCodes: codes,
	})
}

func (c *Cafe) listReferrals(g *gin.Context) {
	// cheap way to lock down this endpoint
	if g.GetHeader("X-Referral-Key") != c.ReferralKey {
		g.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	// get 'em
	refs, err := dao.Dao.ListUnusedReferrals()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	codes := make([]string, len(refs))
	for i, r := range refs {
		codes[i] = r.Code
	}

	// ship it
	g.JSON(http.StatusOK, models.ReferralResponse{
		RefCodes: codes,
	})
}

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func createReferralModel(limit int, requester string) (string, error) {
	code := randString(5)
	ref := models.Referral{
		ID:        bson.NewObjectId(),
		Code:      code,
		Created:   time.Now(),
		Remaining: limit,
		Requester: requester,
	}
	if err := dao.Dao.InsertReferral(ref); err != nil {
		return "", err
	}
	return code, nil
}

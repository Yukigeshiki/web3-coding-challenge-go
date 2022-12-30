package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	l "robothouse.ui/web3-coding-challenge/lib/log"
	"runtime/debug"
)

// Success responds with HTTP status 200 and returns a payload.
func Success(ctx *gin.Context, reqID string, payload interface{}) {
	s := gin.H{"requestId": reqID, "status": "success", "payload": payload}
	log(s)
	ctx.JSON(http.StatusOK, s)
}

// FailBadRequest responds with a StatusBadRequest error code and returns a message.
func FailBadRequest(ctx *gin.Context, reqID string, message any) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, fail(reqID, message))
}

// FailInternalServerError responds with a StatusInternalServerError error code and returns a message.
func FailInternalServerError(ctx *gin.Context, reqID string, message any) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, fail(reqID, message))
}

// log the response as an INFO statement.
func log(res gin.H) {
	f, _ := l.Format(res)
	l.Info(f, nil)
}

// fail logs an error then creates, logs and returns a Gin map
func fail(reqID string, message any) gin.H {
	l.Error(fmt.Sprintf("requestID: %s %s", reqID, debug.Stack()), &reqID)
	f := gin.H{"requestId": reqID, "status": "failed", "message": message}
	log(f)
	return f
}

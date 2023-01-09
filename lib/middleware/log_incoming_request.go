package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"robothouse.io/web3-coding-challenge/lib/log"
	res "robothouse.io/web3-coding-challenge/lib/response"
)

// LogIncomingRequest is used to log the incoming request data.
func LogIncomingRequest(ctx *gin.Context) {
	var logText string
	reqID := uuid.New().String()
	ctx.Set("requestID", reqID)

	method := ctx.Request.Method
	if method == "POST" || method == "PUT" {
		mapBody := make(map[string]any)

		// get the body as bytes and then add it back onto the request for future use
		byteBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			res.FailBadRequest(ctx, reqID, err.Error())
			return
		}
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(byteBody))

		// convert the bytes to a map and format into a clean JSON string
		if err = json.Unmarshal(byteBody, &mapBody); err != nil {
			res.FailBadRequest(ctx, reqID, err.Error())
			return
		}
		JSONString, err := log.Format(mapBody)
		if err != nil {
			res.FailBadRequest(ctx, reqID, err.Error())
			return
		}

		logText = fmt.Sprintf("%s to %s requestBody: %s", method, ctx.Request.URL, JSONString)
	} else {
		logText = fmt.Sprintf("%s to %s ", method, ctx.Request.URL)
	}

	log.Info(logText, &reqID)
}

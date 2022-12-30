package handler

import (
	"github.com/gin-gonic/gin"
	res "robothouse.ui/web3-coding-challenge/lib/response"
	tf "robothouse.ui/web3-coding-challenge/repository/transfers"
	"strconv"
	"strings"
)

const (
	bDec    int = 10
	bitSize int = 64
)

type Handler struct{}

func (h *Handler) CheckApiHealth(ctx *gin.Context) {
	res.Success(ctx, ctx.MustGet("requestID").(string), map[string]string{"message": "I'm running just fine, thanks!"})
}

// GetTransfers fetches an array of transfer logs
func (h *Handler) GetTransfers(ctx *gin.Context) {
	var (
		from, to     string
		above, below int64
	)

	reqID := ctx.MustGet("requestID").(string)

	// check "from" and "to" params - must validate these beforehand
	f, okF := ctx.GetQuery("from")
	if okF {
		from = strings.ToLower(f)
	}
	t, okT := ctx.GetQuery("to")
	if okT {
		to = strings.ToLower(t)
	}
	// fetching all is disabled
	if !okF && !okT {
		res.FailBadRequest(ctx, reqID, "Please supply at least one address")
		return
	}

	// check "above" and "below" params - must validate these beforehand
	if a, ok := ctx.GetQuery("above"); ok {
		above, _ = strconv.ParseInt(a, bDec, bitSize)
	}
	if b, ok := ctx.GetQuery("below"); ok {
		below, _ = strconv.ParseInt(b, bDec, bitSize)
	}

	logs, err := tf.GetLogs(&tf.FilterOpts{From: from, To: to, Above: above, Below: below})
	if err != nil {
		res.FailInternalServerError(ctx, reqID, err.Error())
		return
	}

	res.Success(ctx, reqID, logs)
}

package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	res "robothouse.io/web3-coding-challenge/lib/response"
	val "robothouse.io/web3-coding-challenge/lib/validation"
	tf "robothouse.io/web3-coding-challenge/repository/transfers"
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
	p := new(params)
	reqID := ctx.MustGet("requestID").(string)
	if err := ctx.Bind(p); err != nil {
		res.FailBadRequest(ctx, reqID, err.Error())
		return
	}
	if errs := val.Validate(ctx, p); errs != nil {
		res.FailBadRequest(ctx, reqID, errs)
		return
	}
	fOpts, err := p.toFilterOpts()
	if err != nil {
		res.FailBadRequest(ctx, reqID, err.Error())
		return
	}

	res.Success(ctx, reqID, tf.GetLogs(fOpts, &reqID))
}

// params are the GetTransfers request query parameters
type params struct {
	From  string `form:"from" validate:"omitempty,eth_addr"`
	To    string `form:"to" validate:"omitempty,eth_addr"`
	Above string `form:"above" validate:"omitempty,number"`
	Below string `form:"below" validate:"omitempty,number"`
}

// toFilterOpts maps GetTransfers query params to *transfers.FilterOpts
func (p *params) toFilterOpts() (*tf.FilterOpts, error) {
	var (
		from, to     string
		above, below int64
	)

	f, t, a, b := p.From, p.To, p.Above, p.Below

	// check and convert "from" and "to" params
	if f != "" {
		from = strings.ToLower(f)
	}
	if t != "" {
		to = strings.ToLower(t)
	}
	// fetching all is disabled
	if f == "" && t == "" {
		return nil, errors.New("please supply at least one address")
	}
	// check and convert "above" and "below" params
	if a != "" {
		above, _ = strconv.ParseInt(a, bDec, bitSize)
	}
	if b != "" {
		below, _ = strconv.ParseInt(b, bDec, bitSize)
	}

	return &tf.FilterOpts{From: from, To: to, Above: above, Below: below}, nil
}

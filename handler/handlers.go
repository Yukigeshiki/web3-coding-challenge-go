package handler

import (
	"github.com/gin-gonic/gin"
	res "robothouse.ui/web3-coding-challenge/lib/response"
)

type Handler struct{}

func (h *Handler) CheckApiHealth(ctx *gin.Context) {
	res.Success(ctx, ctx.MustGet("requestID").(string), map[string]string{"message": "I'm running just fine, thanks!"})
}

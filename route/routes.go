package route

import (
	"github.com/gin-gonic/gin"
	"robothouse.io/web3-coding-challenge/handler"
)

// Routes provides client request routing
func Routes(api *gin.RouterGroup) {

	h := new(handler.Handler)

	api.GET("/healthcheck", h.CheckApiHealth)

	api.GET("/transfers", h.GetTransfers)
}

package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"robothouse.ui/web3-coding-challenge/lib/middleware"
	"robothouse.ui/web3-coding-challenge/route"
	"time"
)

// InitGinRouterEngine is used to initialise gin routing
func InitGinRouterEngine() (*gin.Engine, error) {

	router := gin.New()

	// setup CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("erc20/contract")

	// add incoming request logging middleware
	api.Use(middleware.LogIncomingRequest)

	// supply routes
	route.Routes(api)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})

	return router, nil
}

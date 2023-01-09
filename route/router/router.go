package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"robothouse.io/web3-coding-challenge/config"
	eth "robothouse.io/web3-coding-challenge/lib/ethereum"
	"robothouse.io/web3-coding-challenge/lib/middleware"
	repo "robothouse.io/web3-coding-challenge/repository"
	"robothouse.io/web3-coding-challenge/route"
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

	// get ERC20 contract instances
	instHTTP, err := eth.GetContractInstance(config.EthClientURLHTTP, config.ERC20ContractAddr)
	if err != nil {
		return nil, err
	}
	instWS, err := eth.GetContractInstance(config.EthClientURLWS, config.ERC20ContractAddr)
	if err != nil {
		return nil, err
	}

	/*
		Initialise repository (populate in-memory cache)
		This can be done in a separate thread, but the trade-off between having a long start-up time and data consistency
		rather than a short start-up time and data inconsistency seems worthwhile if your application is not going to be
		restarting often, i.e. you're not running in a serverless environment.
		Caching of this type could also be done using Redis in which case a separate service would be responsible for
		populating and updating the data.
	*/
	err = repo.InitialiseRepo(instHTTP)
	if err != nil {
		return nil, err
	}
	/*
		Initialise repo live update (live update in-memory cache)
		This would also be the responsibility of the same separate service as mentioned above if using Redis.
	*/
	go repo.InitialiseRepoLiveUpdate(instWS)

	// add validator to gin context
	api.Use(func(ctx *gin.Context) { ctx.Set("validator", validator.New()) })

	// add incoming request logging middleware
	api.Use(middleware.LogIncomingRequest)

	// supply routes
	route.Routes(api)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	})

	return router, nil
}

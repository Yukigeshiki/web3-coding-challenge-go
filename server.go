package main

import (
	"robothouse.ui/web3-coding-challenge/config"
	repo "robothouse.ui/web3-coding-challenge/repository"
	"robothouse.ui/web3-coding-challenge/route/router"
	//"robothouse.io/roll-coding-challenge/lib/log"
	//"robothouse.io/roll-coding-challenge/route/router"
)

func main() {

	// initialise new gin engine
	e, err := router.InitGinRouterEngine()
	if err != nil {
		panic(err)
	}

	// determine HTTP port
	p := config.Port
	if p == "" {
		p = "8080"
		//log.Debug("main: defaulting to port "+p, nil)
	}

	defer repo.CloseSubscriptionChannels()

	// start the server
	_ = e.Run(":" + p)
}

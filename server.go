package main

import (
	"robothouse.io/web3-coding-challenge/config"
	"robothouse.io/web3-coding-challenge/lib/log"
	"robothouse.io/web3-coding-challenge/route/router"
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
		log.Debug("main: defaulting to port "+p, nil)
	}

	// start the server
	_ = e.Run(":" + p)
}

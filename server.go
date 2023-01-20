package main

import (
	"robothouse.io/web3-coding-challenge/config"
	"robothouse.io/web3-coding-challenge/lib/log"
	"robothouse.io/web3-coding-challenge/route/router"
)

func main() {

	e, err := router.InitGinServer()
	if err != nil {
		panic(err)
	}

	p := config.Port
	if p == "" {
		p = "8080"
		log.Debug("main: defaulting to port "+p, nil)
	}

	// start the server
	_ = e.Run(":" + p)
}

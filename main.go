package main

import (
	"github.com/Yamiyo/account/db"
	"github.com/Yamiyo/account/glob/config"
	"github.com/Yamiyo/account/utils/log"
	"github.com/Yamiyo/account/utils"
	"github.com/Yamiyo/account/server"

	//_ "github.com/Yamiyo/account/glob/init"
)

func main() {
	defer func() {
		db.Close()
		log.Error("Server shutdown...")
		if err := recover(); err != nil {
			log.Errorf("error: %v", err)
		}
	}()

	db.Init(config.Config.DatabaseConfig)

	utils.ErrExit(server.Run())
}

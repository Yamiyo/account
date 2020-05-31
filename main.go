package main

import (
	"github.com/Yamiyo/account/db"
	"github.com/Yamiyo/account/glob/config"

	//_ "github.com/Yamiyo/account/glob/init"
)

func main() {
	config.InitConfig()

	db.Init(config.Config.DatabaseConfig)
}
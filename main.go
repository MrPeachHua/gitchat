package main

import (
	router "ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitRedis()
	r := router.Router()
	r.Run(":8081")
}

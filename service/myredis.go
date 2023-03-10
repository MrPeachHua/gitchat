package service

import (
	"fmt"
	"ginchat/utils"
)

func getBykey(key string) string {
	val, err := utils.Client.Get(key).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	return val
}

func setBykey(key string, value string) {
	err := utils.Client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

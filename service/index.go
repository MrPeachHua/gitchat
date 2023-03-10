package service

import (
	"github.com/gin-gonic/gin"
)

// CreateUser
// @Tags 创建用户
// @Sucess 200 {string} welcom
// @Router /index [get]
func GetIndex(c *gin.Context) {
	setBykey("name", "abcdefg")
	nm := getBykey("name")

	c.JSON(200, gin.H{
		"message": nm,
	})
}

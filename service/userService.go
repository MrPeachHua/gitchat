package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// CreateUser
// @Summary 创建用户
// @Tags 用户模块
// @param name query string false "名称"
// @Sucess 200 {string} 创建首页
// @Router /createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")

	// models.CreateUser(user)

	c.JSON(200, gin.H{
		"message": user.Name,
	})
}

type ConfigParam struct {
	UserId int    `json:"userId" example:"7"`
	Name   string `json:"name" example:"sdfdsf"`
}

// DeleteUser
// @Summary 删除用户
// @Tags 删除用户
// @param body body ConfigParam true  "JSON数据"
// @Accept       json
// @Produce      json
// @Sucess 200 {string} 删除
// @Router /deleteUser [post]
func DeleteUser(c *gin.Context) {
	json := ConfigParam{}
	c.BindJSON(&json)
	c.JSON(200, gin.H{
		"message": json.Name,
	})
}

package models

import (
	"ginchat/utils"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name         string
	PassWord     string
	Phone        string
	Identity     string
	ClientIp     string
	ClientPort   string
	LoginTime    uint64
	HeartbeaTime uint64
	LoginOutTime uint64
	DeviceInfo   string
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func (table *UserBasic) TableName() string {
	return "zhangsan"
}

package models

import (
	"time"
)

type SystemAdmin struct {
	ID         string    `gorm:"primary_key"`
	AdminPhone string    `json:"admin_phone" gorm:"type:varchar(20);unique;not null"` //注册账号
	VerifyCode string    `json:"verify_code" gorm:"type:varchar(20);not null;default:'147258'"`
	CreateTime time.Time `json:"createTime" ` //创建时间gorm:"not null"
	UpdateTime time.Time `json:"updateTime" ` // 修改时间
}

//登录实体
type SystemAdminLogin struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verify_code"`
}

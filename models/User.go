package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Phone        string    `json:"phone" gorm:"type:varchar(12);unique;not null"`
	DeviceId     int       `json:"device_id" gorm:"size:11;` //绑定的设备id
	Email        string    `json:"email" gorm:"type:varchar(64)`
	MpHash       string    `json:"mp_hash" gorm:"type:varchar(64)`        //登录密码
	EncriptedPin string    `json:"encripted_pin" gorm:"type:varchar(128)` //
	SecretFactor string    `json:"secret_factor" gorm:"type:varchar(16)`  //
	Salt         string    `json:"salt" gorm:"type:varchar(60)"`          //盐
	CreateTime   time.Time `json:"create_time" `                          //创建时间gorm:"not null"
	UpdateTime   time.Time `json:"update_time" `                          // 修改时间
	LoginTime    time.Time `json:"login_time" `
}

type GetVerifyCode struct {
	Phone      string `json:"phone"`
	VerifyType string `json:"verify_type"` //注册 "register"、登录 "login"、恢复数据 "restore_data"、重置Pin "reset_pin"

}

//注册
type RegisterUser struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verify_code"`
	MpHash     string `json:"mp_hash"`
}

//生成盐和密码
func (user *User) AddSaltedPassword() error {

	salt := uuid.New().String()

	passwordBytes := []byte(user.MpHash + salt)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.MpHash = string(hash[:])
	user.Salt = salt

	return nil
}

//校验密码
func (user *User) ComparePasswords(pwd string) bool {
	byteHash := []byte(pwd + user.Salt)
	plainHash := []byte(user.MpHash)
	err := bcrypt.CompareHashAndPassword(plainHash, byteHash)
	if err != nil {
		return false
	}
	return true

}

package models

import (
	"time"
)

type Good struct {
	ID         int       `json:"id" gorm:"primary_key;unique"`
	GoodsName  string    `json:"goods_name" gorm:"type:varchar(20)`      //商品名称
	Price      string    `json:"price" gorm:"type:varchar(20);not null"` //商品价钱
	Picture    string    `json:"picture " gorm:"type:varchar(60)`        //图片路径
	CreateTime time.Time `json:"create_time" `                           //创建时间gorm:"not null"
	UpdateTime time.Time `json:"update_time"`
}

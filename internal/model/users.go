package model

import (
	"github.com/zhufuyi/sponge/pkg/ggorm"
)

type Users struct {
	ggorm.Model `gorm:"embedded"` // embed id and time

	Name         string `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
	Email        string `gorm:"column:email;type:varchar(255);NOT NULL" json:"email"`
	MobileNumber string `gorm:"column:mobile_number;type:varchar(20);NOT NULL" json:"mobileNumber"`
	Age          int    `gorm:"column:age;type:int(11);NOT NULL" json:"age"`
}

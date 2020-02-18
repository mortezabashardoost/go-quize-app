package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	UserId		uint64		`gorm:"primary_key;auto_increment" json:"userid"`
	Name	    string    	`gorm:"size:255;not null" json:"name"`
	email	    string    	`gorm:"size:255;not null" json:"email"`
	password    string    	`gorm:"size:255;not null" json:"password"`
	CreatedOn 	time.Time 	`gorm:"default:CURRENT_TIMESTAMP" json:"createdon"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Model(&User{}).Create(&u).error
	if err != nil {
		return &User{},err
	}
	return u, nil
}

package models

import (
	"errors"
	//"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Quize struct {
	QuizeId		uint64		`gorm:"primary_key;auto_increment" json:"quizeid"`
	Title 		string		`gorm:"size:255;not null" json:"title"`
	UserId 		uint64		`gorm:"not null" json:"userid"`
	CreatedOn 	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"createdon"`
	CreatedBy 	uint64		`gorm:"not null" json:"createdby"`
}

func (q *Quize) SaveQuize(db *gorm.DB) (*Quize, error) {
	var err error
	err = db.Debug().Model(&Quize{}).Create(&q).Error
	if err != nil {
		return &Quize{},err
	}
	// TODO: check user id
	// if q.QuizeId != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?",q.UserId)
	// }
	return q, nil
}
package models

import (
	"mcblog/utils/errno"

	"gorm.io/gorm"
)

type Wish struct {
	ID     uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name   string `gorm:"type:varchar(40);not null" json:"name"`
	Status bool   `gorm:"type:bool;default:false" json:"status"`
}

// 新增心愿
func CreateWish(wish *Wish) error {
	err = db.Create(wish).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 分页查询心愿列表
func GetWishes(pageSize, pageNum int) ([]Wish, int64, error) {
	var wishs []Wish
	var total int64
	err := db.Find(&wishs).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}
	return wishs, total, nil
}

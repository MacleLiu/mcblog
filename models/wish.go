package models

import (
	"fmt"
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

// 查询单个心愿信息
func GetWish(id int) (Wish, error) {
	var wish Wish
	err := db.Where("id = ?", id).First(&wish).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return wish, errno.New(errno.ERROR, err)
	}
	return wish, nil
}

// 分页查询心愿列表
func GetWishes(pageSize, pageNum int) ([]Wish, int64, error) {
	var wishs []Wish
	var total int64
	err := db.Model(&wishs).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&wishs).Error
	fmt.Println(len(wishs))
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}
	return wishs, total, nil
}

// 编辑心愿信息
func EditWish(id int, data Wish) error {
	var wish Wish
	dataMap := make(map[string]any)
	dataMap["name"] = data.Name
	dataMap["status"] = data.Status

	err = db.Model(&wish).Where("id = ?", id).Updates(dataMap).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 删除心愿
func DeleteWish(id int) error {
	var wish Wish
	err := db.Where("id = ?", id).Delete(&wish).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

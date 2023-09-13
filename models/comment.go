package models

import (
	"mcblog/utils/errno"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Article Article `gorm:"foreignKey:Aid"`
	Name    string  `gorm:"type:varchar(20);not null" json:"name"`
	Content string  `gorm:"type:longtext" json:"content"`
	Aid     int     `gorm:"type:int;not null" json:"artid"`
	Toid    int     `gorm:"type:int;" json:"toid"`
	Toname  string  `gorm:"type:varchar(20)" json:"toname"`
}

// 新增评论
func CreateComment(comt *Comment) (uint, error) {
	err = db.Create(comt).Error
	if err != nil {
		return 0, errno.New(errno.ERROR, err)
	}
	return comt.ID, nil
}

// 获取文章评论
func GetCommentByArt(id int) ([]Comment, int64, error) {
	var comts []Comment
	var total int64
	err := db.Where("aid = ?", id).Order("id").Find(&comts).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}

	return comts, total, nil
}

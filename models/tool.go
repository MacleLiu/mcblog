package models

import (
	"mcblog/utils/errno"

	"gorm.io/gorm"
)

type Tool struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(40);not null" json:"name"`
	Url  string `gorm:"type:varchar(100);not null" json:"url"`
}

// 新增工具
func CreateTool(tool *Tool) error {
	err = db.Create(tool).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 分页查询工具列表
func GetTools(pageSize, pageNum int) ([]Tool, int64, error) {
	var tools []Tool
	var total int64
	err := db.Find(&tools).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}
	return tools, total, nil
}

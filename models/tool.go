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

// 查询单个工具信息
func GetTool(id int) (Tool, error) {
	var tool Tool
	err := db.Where("id = ?", id).First(&tool).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tool, errno.New(errno.ERROR, err)
	}
	return tool, nil
}

// 分页查询工具列表
func GetTools(pageSize, pageNum int) ([]Tool, int64, error) {
	var tools []Tool
	var total int64
	err := db.Model(&tools).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&tools).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}
	return tools, total, nil
}

// 编辑工具信息
func EditTool(id int, data Tool) error {
	var tool Tool
	dataMap := make(map[string]any)
	dataMap["name"] = data.Name
	dataMap["url"] = data.Url

	err = db.Model(&tool).Where("id = ?", id).Updates(dataMap).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 删除工具
func DeleteTool(id int) error {
	var tool Tool
	err := db.Where("id = ?", id).Delete(&tool).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

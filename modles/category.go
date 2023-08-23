package modles

import (
	"errors"
	"mcblog/utils/errno"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 检查分类是否存在
func CheckCategory(id int) error {
	var cate Category
	err := db.Select("name").Where("id = ?", id).First(&cate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.New(errno.ERROR_CATE_NOT_EXIST, err)
		}
		return errno.New(errno.ERROR, err)
	}

	return nil
}

// 检查分类名是否已存在
func CheckCategoryName(name string) error {
	var cate Category
	err := db.Select("id").Where("name=?", name).First(&cate).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.New(errno.ERROR_CATE_NOT_EXIST, err)
		}
		return errno.New(errno.ERROR, err)
	}

	if cate.ID > 0 {
		return nil
	}
	return errno.New(errno.ERROR, err)
}

// 新增分类
func CreateCategory(cate *Category) error {
	err = db.Create(cate).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 查询单个分类信息
func GetCategory(id int) (Category, error) {
	var cate Category
	err := db.Where("id = ?", id).First(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return cate, errno.New(errno.ERROR, err)
	}
	return cate, nil
}

// 查询分类列表
func GetCategories(pageSize, pageNum int) ([]Category, int64, error) {
	var cates []Category
	var total int64
	err := db.Find(&cates).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}
	return cates, total, nil
}

// 编辑分类信息
func EditCategory(id int, data Category) error {
	var cate Category
	dataMap := make(map[string]any)
	dataMap["name"] = data.Name

	err = db.Model(&cate).Where("id = ?", id).Updates(dataMap).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 删除分类
func DeleteCategory(id int) error {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

package models

import (
	"errors"
	"mcblog/utils/errno"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

type CateStat struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
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

// 分页查询分类列表
func GetCategories(pageSize, pageNum int) ([]Category, int64, error) {
	var cates []Category
	var total int64
	err := db.Model(&cates).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}
	return cates, total, nil
}

// 查询分类统计信息（分类下文章数量）
func GetCateStat() ([]CateStat, error) {
	var cateStat []CateStat
	err := db.Raw("SELECT category.id, r.name, r.count FROM category JOIN (SELECT c.name, count(a.id) AS count FROM category c LEFT JOIN article a ON a.deleted_at IS NULL AND c.id = a.cid GROUP BY c.name) AS r ON category.name = r.name").Scan(&cateStat).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errno.New(errno.ERROR, err)
	}
	return cateStat, nil
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

// 检查分类是否被使用
func CheckCateUsed(id int) error {
	var art Article
	err := db.Where("cid = ?", id).First(&art).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errno.New(errno.ERROR, err)
	}

	if art.ID > 0 {
		return errno.New(errno.ERROR_CATE_DELETE_USED, err)
	}
	return errno.New(errno.ERROR, err)
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

// 分页查询分类下的文章
func GetCateArticles(cid int, pageSize, pageNum int) ([]Article, int64, error) {
	var arts []Article
	var total int64
	err := db.Preload("Category").Where("cid = ?", cid).Model(&arts).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at DESC").Find(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}

	return arts, total, nil
}

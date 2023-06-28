package modles

import (
	"fmt"
	"mcblog/utils/errno"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 检查分类是否存在
func CheckCategory(name string) error {
	var cate Category
	db.Select("id").Where("name=?", name).First(&cate)

	if cate.ID > 0 {
		return errno.New(errno.ERROR_CATENAME_USED, err)
	}
	return nil
}

// 新增分类
func CreateCategory(cate *Category) error {
	err = db.Create(cate).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 查询分类列表
func GetCategories(pageSize, pageNum int) ([]Category, error) {
	var cates []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errno.New(errno.ERROR, err)
	}
	return cates, nil
}

// 编辑分类信息
func EditCategory(id int, data Category) error {
	var cate Category
	dataMap := make(map[string]any)
	dataMap["name"] = data.Name

	fmt.Println("dataMap: ", dataMap)

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

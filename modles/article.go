package modles

import (
	"errors"
	"mcblog/utils/errno"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `grom:"type:varchar(100)" json:"img"`
}

// 检查文章是否存在
func CheckArticleExist(id int) error {
	var art Article
	if err := db.Where("id = ?", id).First(&art).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.New(errno.ERROR_ARTICLE_NOT_EXIST, err)
		}
		return errno.New(errno.ERROR, err)
	}

	return nil
}

// 新增文章
func CreateArticle(art *Article) error {
	err = db.Create(art).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 查询单个文章
func GetArticle(id int) (Article, error) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Article{}, errno.New(errno.ERROR_ARTICLE_NOT_EXIST, err)
		}
		return Article{}, errno.New(errno.ERROR, err)
	}
	return art, nil
}

// 查询文章列表
func GetArticles(pageSize, pageNum int, title string) ([]Article, int64, error) {
	var arts []Article
	var total int64
	var err error
	if title == "" {
		err = db.Preload("Category").Find(&arts).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	} else {
		err = db.Where("title LIKE ?", title+"%").Find(&arts).Count(&total).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}

	return arts, total, nil
}

// 查询分类下的文章
func GetCateArticles(cid int, pageSize, pageNum int) ([]Article, int64, error) {
	var arts []Article
	var total int64
	err := db.Preload("Category").Where("cid = ?", cid).Find(&arts).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}

	return arts, total, nil
}

// 编辑文章
func EditArticle(id int, data Article) error {
	var art Article
	dataMap := make(map[string]any)
	dataMap["title"] = data.Title
	dataMap["cid"] = data.Cid
	dataMap["desc"] = data.Desc
	dataMap["content"] = data.Content
	dataMap["img"] = data.Img

	err = db.Model(&art).Where("id = ?", id).Updates(dataMap).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 删除文章
func DeleteArticle(id int) error {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

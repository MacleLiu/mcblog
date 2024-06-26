package models

import (
	"errors"
	"mcblog/utils/errno"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignKey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null;index:idx_fulltext, class:FULLTEXT, option:WITH PARSER ngram" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200);index:idx_fulltext, class:FULLTEXT, option:WITH PARSER ngram" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
	Winnow   bool     `gorm:"type:bool;default:false" json:"winnow"`
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
func CreateArticle(art *Article) (uint, error) {
	err = db.Create(art).Error
	if err != nil {
		return 0, errno.New(errno.ERROR, err)
	}
	return art.ID, nil
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

// 查询全部文章的基础信息
func GetAllArtInfo() ([]Article, error) {
	var arts []Article
	err := db.Select("id", "created_at", "title").Find(&arts).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return arts, errno.New(errno.ERROR_ARTICLE_NOT_EXIST, err)
		}
		return arts, errno.New(errno.ERROR, err)
	}
	return arts, nil
}

// 查询文章列表(分页或关键词全文索引查询)
func GetArticles(pageSize, pageNum int, keyword string) ([]Article, int64, error) {
	var arts []Article
	var total int64
	var err error
	if keyword == "" {
		err = db.Preload("Category").Model(&arts).Select("id", "created_at", "title", "cid", "desc", "img", "winnow").Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at DESC").Find(&arts).Error
	} else {
		err = db.Preload("Category").Model(&arts).Select("id", "created_at", "title", "cid", "desc", "img", "winnow").Where("MATCH(`title`, `desc`) AGAINST(? IN NATURAL LANGUAGE MODE)", keyword).Order("created_at DESC").Find(&arts).Count(&total).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, errno.New(errno.ERROR, err)
	}

	return arts, total, nil
}

// 查询文章总数(主要给统计信息使用)
func GetArticCuont() (int64, error) {
	var c int64
	err := db.Table("article").Count(&c).Error
	if err != nil {
		return 0, errno.New(errno.ERROR, err)
	}
	return c, nil
}

// 查询精选文章列表
func GetWinnowArticles() ([]Article, error) {
	var arts []Article
	err := db.Preload("Category").Model(&arts).Select("id", "created_at", "title", "cid", "desc", "img", "winnow").Where("winnow = ?", true).Find(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errno.New(errno.ERROR, err)
	}

	return arts, nil
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

// 设置精选文章
func SetWinnow(id int) error {
	err := db.Exec("UPDATE article SET winnow = NOT winnow Where id = ?", id).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return err
}

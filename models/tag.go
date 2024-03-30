package models

import (
	"errors"
	"mcblog/utils/errno"

	"gorm.io/gorm"
)

type Tag struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 文章和标签关联表
type ArticleTag struct {
	ID      uint    `gorm:"primary_key;auto_increment" json:"id"`
	Tag     Tag     `gorm:"foreignKey:Tag_id"`
	Article Article `gorm:"foreignKey:Art_id "`
	Art_id  uint    `gorm:"type:int;not null" json:"aid"`
	Tag_id  uint    `gorm:"type:int;not null" json:"tid"`
}

// 检查标签名是否已存在
func CheckTagName(name string) error {
	var tag Tag
	err := db.Select("id").Where("name=?", name).First(&tag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.New(errno.ERROR_TAG_NOT_EXIST, err)
		}
		return errno.New(errno.ERROR, err)
	}

	if tag.ID > 0 {
		return nil
	}
	return errno.New(errno.ERROR, err)
}

// 检查标签是否存在
func CheckTag(id int) error {
	var tag Tag
	err := db.Select("name").Where("id = ?", id).First(&tag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.New(errno.ERROR_TAG_NOT_EXIST, err)
		}
		return errno.New(errno.ERROR, err)
	}

	return nil
}

// 新增标签
func CreateTag(tag *Tag) error {
	err = db.Create(tag).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 查询一个标签
func GetTag(tid int) (Tag, error) {
	var tag Tag
	err := db.Where("id = ?", tid).Find(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, errno.New(errno.ERROR, err)
	}
	return tag, nil
}

// 查询标签列表
func GetTags() ([]Tag, int64, error) {
	var tags []Tag
	var total int64
	err := db.Find(&tags).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tags, 0, errno.New(errno.ERROR, err)
	}
	return tags, total, nil
}

// // 分页查询标签列表
// func GetTags(pageSize, pageNum int) ([]Tag, int64, error) {
// 	var tags []Tag
// 	var total int64
// 	err := db.Find(&tags).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, 0, errno.New(errno.ERROR, err)
// 	}
// 	return tags, total, nil
// }

// 编辑标签信息
func EditTag(id int, data Tag) error {
	var tag Tag
	dataMap := make(map[string]any)
	dataMap["name"] = data.Name

	err = db.Model(&tag).Where("id = ?", id).Updates(dataMap).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 检查标签是否被使用
func CheckTagUsed(id int) error {
	var artTag ArticleTag
	err := db.Where("tag_id = ?", id).First(&artTag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return errno.New(errno.ERROR, err)
	}

	if artTag.ID > 0 {
		return errno.New(errno.ERROR_TAG_DELETE_USED, err)
	}
	return errno.New(errno.ERROR, err)
}

// 删除标签
func DeleteTag(id int) error {
	var tag Tag
	err := db.Where("id = ?", id).Delete(&tag).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

// 分页查询标签下的文章信息
func GetTagArticles(tid int, pageSize, pageNum int) ([]Article, int64, error) {
	var arts []Article
	var total int64
	// 将文章标签关联表和标签表结合成一个新的衍生表
	// query := db.Table("article_tag").Select("art_id", "tag_id").Joins("join tag on article_tag.tag_id = tag.id").Where("tag.id = ?", tid)
	// err := db.Preload("Tag").Model(&arts).Select("id", "created_at", "title", "cid", "desc").Joins("join (?) q on article.id = q.art_id", query).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
	err := db.Preload("Category").Model(&arts).Select("article.id", "created_at", "title", "cid", "desc").Joins("join article_tag on article.id = article_tag.art_id AND article_tag.tag_id = ?", tid).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at DESC").Find(&arts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return arts, 0, errno.New(errno.ERROR, err)
	}
	return arts, total, nil
}

// 查询文章标签
func GetArticleTags(aid int) ([]Tag, error) {
	var tags []Tag
	err := db.Model(&tags).Select("tag.id", "name").Joins("join article_tag on tag.id = article_tag.tag_id AND article_tag.art_id = ?", aid).Scan(&tags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tags, errno.New(errno.ERROR, err)
	}
	return tags, nil
}

// 更新文章标签
func UpdateArticleTags(aid int, tags []Tag) error {
	var atag ArticleTag
	var atags []ArticleTag
	// 先删除全部标签再批量插入标签
	err := db.Where("art_id = ?", aid).Delete(&atag).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}

	if len(tags) != 0 {
		for _, v := range tags {
			item := ArticleTag{Art_id: uint(aid), Tag_id: uint(v.ID)}
			atags = append(atags, item)
		}
		err = db.Create(&atags).Error
		if err != nil {
			return errno.New(errno.ERROR, err)
		}
	}
	return nil
}

// 删除文章标签
func DeleteArticleTags(aid int) error {
	var atag ArticleTag
	err := db.Where("art_id = ?", aid).Delete(&atag).Error
	if err != nil {
		return errno.New(errno.ERROR, err)
	}
	return nil
}

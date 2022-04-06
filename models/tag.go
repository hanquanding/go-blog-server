/**
 * @author: hqd
 * @description: tag model
 * @file: tag
 * @date: 2021-02-12 18:26
 */

package models

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/hqd8080/go-blog-server/models/form"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	TagID     uint   `gorm:"primary_key" json:"tag_id"`
	TagName   string `gorm:"column:tag_name;   not null; type:varchar(100); default:''; comment:'标签名称';"`
	TagStatus uint   `gorm:"column:tag_status; not null; default:0; comment:'标签状态（0：禁用,1：启用）';"`
}

type TagModel struct {
	DB *gorm.DB
}

func NewTagModel(db *gorm.DB) *TagModel {
	return &TagModel{DB: db}
}

func (t *TagModel) Create(tag *Tag) (uint, error) {
	if err := t.DB.Model(&Tag{}).Create(tag).Error; err != nil {
		return 0, err
	}
	return tag.TagID, nil
}

func (t *TagModel) Update(tagID uint, data map[string]interface{}) error {
	return t.DB.Model(&Tag{}).Where("id = ?", tagID).Update(data).Error
}

func (t *TagModel) Delete(ids []int) error {
	return t.DB.Model(&Tag{}).Delete(Tag{}, ids).Error
}

func (t *TagModel) Get(tagID uint) (*Tag, error) {
	var tag Tag
	err := t.DB.Model(&Tag{}).Where("id = ? AND is_deleted = ?", tagID, 0).First(&tag).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (t *TagModel) ExistByTagName(name string, ID uint) bool {
	var tag Tag
	err := t.DB.Model(&Tag{}).Unscoped().Where("tag_name = ?", name).Where("id != ?", ID).Find(&tag).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (t *TagModel) List(params form.TagQueryParam, opts form.TagQueryOptions) (*form.PaginationResult, error) {
	db := t.DB.Model(&Tag{})
	if params.TagName != "" {
		db = db.Where("tag_name LIKE ?", fmt.Sprintf("%%%s%%", params.TagName))
	}
	if params.TagStatus == "all" {
		db = db.Where("tag_status in(?)", []int{0, 1})
	} else {
		db = db.Where("tag_status = ?", params.TagStatus)
	}
	db = db.Order("id desc")

	var tags form.Tags
	total, err := queryBuilder(db, opts.PageParam, &tags)
	if err != nil {
		logs.Info("page query err:%s", err.Error())
		return nil, err
	}
	if total == 0 {
		return nil, nil
	}
	result := &form.PaginationResult{
		List:  tags.ToFormTags(),
		Total: total,
	}
	return result, nil

}

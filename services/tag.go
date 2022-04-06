/**
 * @author: hqd
 * @description: tag service
 * @file: tag
 * @date: 2021-02-12 19:47
 */

package services

import (
	"github.com/hqd8080/go-blog-server/models"
	"github.com/hqd8080/go-blog-server/models/form"
)

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Create(tag form.AddTag) (uint, error) {
	newTag := &models.Tag{
		TagName:   tag.TagName,
		TagStatus: tag.TagStatus,
	}
	return models.NewTagModel(dbMaster).Create(newTag)
}

func (t *Tag) Update(tag form.UpdateTag) error {
	data := map[string]interface{}{
		"tag_name":   tag.TagName,
		"tag_status": tag.TagStatus,
	}
	return models.NewTagModel(dbMaster).Update(tag.ID, data)
}

func (t *Tag) Delete(ids []int) error {
	return models.NewTagModel(dbMaster).Delete(ids)
}

func (t *Tag) Get(tagID uint) (*form.Tag, error) {
	data, err := models.NewTagModel(dbSlave).Get(tagID)
	if err != nil {
		return nil, err
	}
	return &form.Tag{
		ID:        data.TagID,
		TagName:   data.TagName,
		TagStatus: data.TagStatus,
		CreatedAt: data.CreatedAt.Format(models.TimeFormat),
		UpdatedAt: data.UpdatedAt.Format(models.TimeFormat),
	}, err
}

func (t *Tag) List(params form.TagQueryParam, opts form.TagQueryOptions) (*form.PaginationResult, error) {
	return models.NewTagModel(dbSlave).List(params, opts)
}

func (t *Tag) ExistByTagName(TagName string, ID uint) bool {
	return models.NewTagModel(dbSlave).ExistByTagName(TagName, ID)
}

func (t *Tag) UpdateTagStatus(tagID uint) error {
	return nil
}

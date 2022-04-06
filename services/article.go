/**
 * @author: hqd
 * @description: article service
 * @file: article
 * @date: 2021-02-12 21:15
 */

package services

import (
	"github.com/hqd8080/go-blog-server/models"
	"github.com/hqd8080/go-blog-server/models/form"
)

type Article struct{}

func NewArticle() *Article {
	return &Article{}
}

func (a *Article) Create(req form.AddArticle) (uint, error) {
	data := models.Article{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		CoverImgURL: req.CoverImgURL,
		Sort:        req.Sort,
		Status:      req.Status,
	}
	return models.NewArticleModel(dbMaster).Create(data)
}

func (a *Article) Update(article form.UpdateArticle) error {
	data := map[string]interface{}{
		"title":         article.Title,
		"description":   article.Description,
		"content":       article.Content,
		"cover_img_url": article.CoverImgURL,
		"sort":          article.Sort,
		"status":        article.Status,
	}
	return models.NewArticleModel(dbMaster).Update(article.ID, data)
}

func (a *Article) Delete(IDs []int) error {
	return models.NewArticleModel(dbMaster).Delete(IDs)
}

func (a *Article) List(params form.ArticleQueryParam, opts form.ArticleQueryOptions) (*form.PaginationResult, error) {
	return models.NewArticleModel(dbSlave).List(params, opts)
}

func (a *Article) Get(ID uint) (*form.Article, error) {
	data, err := models.NewArticleModel(dbSlave).Get(ID)
	if err != nil {
		return nil, err
	}
	return &form.Article{
		ID:          data.ArticleID,
		Title:       data.Title,
		Description: data.Description,
		Content:     data.Content,
		CoverImgURL: data.CoverImgURL,
		Sort:        data.Sort,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt.Format(models.TimeFormat),
		UpdatedAt:   data.UpdatedAt.Format(models.TimeFormat),
	}, err
}

func (a *Article) ExistByTitle(title string, ID uint) bool {
	return models.NewArticleModel(dbSlave).ExistByTitle(title, ID)
}

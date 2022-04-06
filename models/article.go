/**
 * @author: hqd
 * @description: article model
 * @file: article
 * @date: 2021-02-12 21:19
 */

package models

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/hqd8080/go-blog-server/models/form"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	ArticleID   uint   `gorm:"primary_key" json:"article_id"`
	Title       string `gorm:"column:title;  not null; type:varchar(150); default:''; comment:'文章标题';"`
	Description string `gorm:"column:description; not null; type:varchar(1000); default:''; comment:'文章描述';"`
	Content     string `gorm:"column:content; not null; default:''; comment:'文章内容';"`
	CoverImgURL string `gorm:"column:cover_img_url; not null; type:varchar(200); default:''; comment:'文章封面图';"`
	Visits      int    `gorm:"column:visits; not null; default:0; comment:'文章浏览次数';"`
	Sort        int    `gorm:"column:sort; not null; default:0; comment:'文章排序';"`
	Status      uint   `gorm:"column:status; not null; default:0; comment:'文章状态（0：禁用,1：启用）';"`
}

type ArticleModel struct {
	DB *gorm.DB
}

func NewArticleModel(db *gorm.DB) *ArticleModel {
	return &ArticleModel{DB: db}
}

func (a *ArticleModel) ExistByTitle(title string, ID uint) bool {
	var article Article
	err := a.DB.Model(&Article{}).Unscoped().Where("title = ?", title).Where("id != ?", ID).Find(&article).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (a *ArticleModel) Create(article Article) (uint, error) {
	if err := a.DB.Model(&Article{}).Create(&article).Error; err != nil {
		return 0, err
	}
	return article.ArticleID, nil
}

func (a *ArticleModel) Update(ID uint, data map[string]interface{}) error {
	return a.DB.Model(&Article{}).Where("id = ?", ID).Update(data).Error
}

func (a *ArticleModel) Delete(IDs []int) error {
	return a.DB.Model(&Article{}).Delete(Article{}, IDs).Error
}

func (a *ArticleModel) List(params form.ArticleQueryParam, opts form.ArticleQueryOptions) (*form.PaginationResult, error) {
	db := a.DB
	db = db.Model(&Article{}).Select("id, title, description, cover_img_url, content, sort, status, is_deleted, created_at, updated_at")
	if params.Title != "" {
		db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", params.Title))
	}
	if params.Description != "" {
		db = db.Where("description LIKE ?", fmt.Sprintf("%%%s%%", params.Description))
	}
	if params.Status == "all" {
		db = db.Where("status in(?)", []int{0, 1})
	} else {
		db = db.Where("status = ?", params.Status)
	}
	db = db.Order("sort desc")

	var fa form.Articles
	total, err := queryBuilder(db, opts.PageParam, &fa)
	if err != nil {
		logs.Info("page query err:%s", err.Error())
		return nil, err
	}
	if total == 0 {
		return nil, nil
	}

	result := &form.PaginationResult{
		List:  fa.ToFormArticles(),
		Total: total,
	}
	return result, nil
}

func (a *ArticleModel) Get(ID uint) (*Article, error) {
	var article Article
	err := a.DB.Model(&Article{}).Where("id = ? AND is_deleted = ?", ID, 0).First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

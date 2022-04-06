/**
 * @author: hqd
 * @description: website model
 * @file: website
 * @date: 2021-02-07 13:13
 */

package models

import (
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	"github.com/hqd8080/go-blog-server/models/form"
	"github.com/jinzhu/gorm"
)

type Website struct {
	Model
	WebID          uint   `gorm:"primary_key" json:"web_id"`
	WebName        string `gorm:"column:web_name; not null; type:varchar(150); default:''; comment:'网站名称';"`
	WebURL         string `gorm:"column:web_url;  not null; type:varchar(200); default:''; comment:'网站地址';"`
	WebDescription string `gorm:"column:web_description; not null; type:varchar(1000); comment:'网站描述';"`
	WebSort        int    `gorm:"column:web_sort;   not null; default:0; comment:'列表排序';"`
	WebStatus      uint   `gorm:"column:web_status; not null; default:0; comment:'网站状态（0：不显示,1：显示）';"`
}

type WebsiteModel struct {
	DB *gorm.DB
}

func NewWebsiteModel(db *gorm.DB) *WebsiteModel {
	return &WebsiteModel{DB: db}
}

func (w *WebsiteModel) Create(web *Website) (uint, error) {
	if err := w.DB.Model(&Website{}).Create(web).Error; err != nil {
		return 0, err
	}
	return web.WebID, nil
}

func (w *WebsiteModel) Update(websiteID uint, data map[string]interface{}) error {
	return w.DB.Model(&Website{}).Where("web_id = ?", websiteID).Update(data).Error
}

func (w *WebsiteModel) Get(websiteID uint) (*Website, error) {
	var web Website
	err := w.DB.Model(&Website{}).Where("web_id = ? AND is_deleted = ?", websiteID, 0).First(&web).Error
	if err != nil {
		return nil, err
	}
	return &web, nil
}

func (w *WebsiteModel) Delete(ids []int) error {
	return w.DB.Model(&Website{}).Delete(Website{}, ids).Error
}

func (w *WebsiteModel) ExistByWebName(webName string, ID ...uint) bool {
	var web Website
	err := w.DB.Model(&Website{}).Unscoped().Where("web_name = ?", webName).Where("web_id != ?", ID).Find(&web).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (w *WebsiteModel) ExistByWebURL(webURL string, ID uint) bool {
	var web Website
	err := w.DB.Model(&Website{}).Unscoped().Where("web_url = ?", webURL).Where("web_id != ?", ID).Find(&web).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (w *WebsiteModel) List(params form.WebsiteQueryParam, opts form.WebsiteQueryOptions) (*form.PaginationResult, error) {
	db := w.DB
	db = db.Model(&Website{}).Select("web_id, web_name, web_url, web_description, web_sort, web_status, is_deleted, created_at, updated_at")
	if params.WebName != "" {
		db = db.Where("web_name LIKE ?", fmt.Sprintf("%%%s%%", params.WebName))
	}
	if params.WebURL != "" {
		db = db.Where("web_url = ?", params.WebURL)
	}
	if params.WebStatus == "all" {
		db = db.Where("web_status in(?)", []int{0, 1})
	} else {
		db = db.Where("web_status = ?", params.WebStatus)
	}
	db = db.Order("web_sort desc")

	var web form.Websites
	total, err := queryBuilder(db, opts.PageParam, &web)
	if err != nil {
		logs.Info("page query err:%s", err.Error())
		return nil, err
	}
	if total == 0 {
		return nil, nil
	}
	result := &form.PaginationResult{
		List:  web.ToFormWebsites(),
		Total: total,
	}
	return result, nil
}

func (w *WebsiteModel) UpdateWebStatus(websiteID uint) error {
	result, err := w.Get(websiteID)
	if err != nil {
		return err
	}
	attrs := make(map[string]interface{})
	if result.WebStatus == 0 {
		attrs["web_status"] = 1
	} else {
		attrs["web_status"] = 0
	}
	return w.DB.Model(&Website{}).Where("web_id = ?", websiteID).Update(attrs).Error
}

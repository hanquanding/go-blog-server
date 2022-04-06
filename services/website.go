/**
 * @author: hqd
 * @description: website service
 * @file: website
 * @date: 2021-02-07 21:31
 */

package services

import (
	"github.com/hqd8080/go-blog-server/models"
	"github.com/hqd8080/go-blog-server/models/form"
)

type Website struct{}

func NewWebsite() *Website {
	return &Website{}
}

func (w *Website) Create(fw form.AddWebsiteReq) (uint, error) {
	model := models.NewWebsiteModel(dbMaster)
	web := &models.Website{
		WebName:        fw.WebName,
		WebURL:         fw.WebURL,
		WebDescription: fw.WebDescription,
		WebSort:        fw.WebSort,
		WebStatus:      fw.WebStatus,
	}
	return model.Create(web)
}

func (w *Website) Update(req form.UpdateWebsiteReq) error {
	model := models.NewWebsiteModel(dbMaster)
	data := map[string]interface{}{
		"web_name":        req.WebName,
		"web_url":         req.WebURL,
		"web_description": req.WebDescription,
		"web_sort":        req.WebSortNum,
		"web_status":      req.WebStatus,
	}
	return model.Update(req.WebID, data)
}

func (w *Website) UpdateWebStatus(websiteID uint) error {
	return models.NewWebsiteModel(dbMaster).UpdateWebStatus(websiteID)
}

func (w *Website) Get(websiteID uint) (*form.Website, error) {
	data, err := models.NewWebsiteModel(dbSlave).Get(websiteID)
	if err != nil {
		return nil, err
	}
	return &form.Website{
		WebID:          data.WebID,
		WebName:        data.WebName,
		WebURL:         data.WebURL,
		WebDescription: data.WebDescription,
		WebSort:        data.WebSort,
		WebStatus:      data.WebStatus,
		IsDeleted:      data.IsDeleted,
		CreatedAt:      data.CreatedAt.Format(models.TimeFormat),
		UpdatedAt:      data.UpdatedAt.Format(models.TimeFormat),
	}, err
}

func (w *Website) ExistByWebName(webName string, ID ...uint) bool {
	return models.NewWebsiteModel(dbSlave).ExistByWebName(webName, ID...)
}

func (w *Website) ExistByWebURL(webURL string, ID uint) bool {
	return models.NewWebsiteModel(dbSlave).ExistByWebURL(webURL, ID)
}

func (w *Website) Delete(ids []int) error {
	return models.NewWebsiteModel(dbMaster).Delete(ids)
}

func (w *Website) List(params form.WebsiteQueryParam, opts form.WebsiteQueryOptions) (*form.PaginationResult, error) {
	return models.NewWebsiteModel(dbSlave).List(params, opts)
}

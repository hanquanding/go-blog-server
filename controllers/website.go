/**
 * @author: hqd
 * @description: website controller
 * @file: website
 * @date: 2021-02-07 10:20
 */

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	"github.com/hqd8080/go-blog-server/models/form"
	"github.com/hqd8080/go-blog-server/services"
)

type WebsiteController struct {
	BaseController
}

// List 网站信息列表
func (w *WebsiteController) List() {
	params := form.WebsiteQueryParam{
		WebName:   w.GetString("web_name"),
		WebURL:    w.GetString("web_url"),
		WebStatus: w.GetString("web_status", "all"),
	}
	srv := services.NewWebsite()
	result, err := srv.List(params, form.WebsiteQueryOptions{w.GetPaginationParam()})
	if err != nil {
		w.ErrorCode(1, http.StatusInternalServerError, err.Error())
	}
	w.Correct(result)
}

// Create 添加网站信息
func (w *WebsiteController) Create() {
	logs.Info("%v", string(w.Ctx.Input.RequestBody))
	params := form.AddWebsiteReq{}
	err := json.Unmarshal(w.Ctx.Input.RequestBody, &params)
	if err != nil {
		logs.Error("json unmarshal website params err:%s", err.Error())
		w.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}

	err = w.ValidBody(&params)
	if err != nil {
		w.ErrorCode(1, http.StatusForbidden, err.Error())
	}

	exists := services.NewWebsite().ExistByWebName(params.WebName)
	if exists {
		logs.Error("create website err:%s", "website name already exists")
		w.ErrorCode(0, http.StatusOK, fmt.Sprintf("website name `%s` already exists", params.WebName))
	}
	exists = services.NewWebsite().ExistByWebURL(params.WebURL, 0)
	if exists {
		logs.Error("create website err:%s", "website url already exists")
		w.ErrorCode(0, http.StatusOK, fmt.Sprintf("website url `%s` already exists", params.WebURL))
	}

	ID, err := services.NewWebsite().Create(params)
	if err != nil {
		logs.Error("create website err:%s", err.Error())
		w.ErrorCode(1, http.StatusInternalServerError, "create website fail")
	}
	var InsertID = struct {
		WebID uint `json:"web_id"`
	}{WebID: ID}
	w.Correct(InsertID)
}

// Update 修改网站信息
func (w *WebsiteController) Update() {
	params := form.UpdateWebsiteReq{}
	err := json.Unmarshal(w.Ctx.Input.RequestBody, &params)
	if err != nil {
		logs.Error("json unmarshal update website params err:%s", err.Error())
		w.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}

	err = w.ValidBody(&params)
	if err != nil {
		w.ErrorCode(1, http.StatusForbidden, err.Error())
	}

	if params.WebName != "" {
		exists := services.NewWebsite().ExistByWebName(params.WebName, params.WebID)
		if exists {
			logs.Error("update website err:%s", "website name already exists")
			w.ErrorCode(0, http.StatusOK, fmt.Sprintf("website name `%s` already exists", params.WebName))
		}
	}
	if params.WebURL != "" {
		exists := services.NewWebsite().ExistByWebURL(params.WebURL, params.WebID)
		if exists {
			logs.Error("update website err:%s", "website url already exists")
			w.ErrorCode(0, http.StatusOK, fmt.Sprintf("website url `%s` already exists", params.WebURL))
		}

	}

	err = services.NewWebsite().Update(params)
	if err != nil {
		logs.Error("update website err:%s", err.Error())
		w.ErrorCode(1, http.StatusInternalServerError, "update website fail")
	}
	w.Correct(nil)
}

// Get 网站信息详情
func (w *WebsiteController) Get() {
	websiteID, _ := w.GetInt(":id")
	if websiteID == 0 {
		w.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}
	result, err := services.NewWebsite().Get(uint(websiteID))
	if err != nil {
		logs.Error("website get err:%s", err.Error())
		w.ErrorCode(1, http.StatusInternalServerError, err.Error())
	}
	w.Correct(result)
}

// Delete  删除网站信息
func (w *WebsiteController) Delete() {
	ids := w.GetString("ids")
	logs.Info("ids:%s", ids)
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 && idsArr[0] == "" {
		w.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}

	list := make([]int, len(idsArr))
	for i, v := range idsArr {
		ID, err := strconv.Atoi(v)
		if err != nil {
			w.ErrorCode(1, http.StatusBadRequest, "invalid param")
		}
		list[i] = ID
	}
	err := services.NewWebsite().Delete(list)
	if err != nil {
		logs.Error("delete website err:%s", err.Error())
		w.ErrorCode(1, http.StatusInternalServerError, "delete website fail")
	}
	w.Correct(nil)
}

// UpdateWebStatus 更新网站状态
func (w *WebsiteController) UpdateWebStatus() {
	websiteID, _ := w.GetInt(":id")
	if websiteID == 0 {
		w.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}
	err := services.NewWebsite().UpdateWebStatus(uint(websiteID))
	if err != nil {
		logs.Error("update website err:%s", err.Error())
		w.ErrorCode(1, http.StatusInternalServerError, "update website fail")
	}
	w.Correct(nil)
}

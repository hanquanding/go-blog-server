/**
 * @author: hqd
 * @description: article controller
 * @file: article
 * @date: 2021-02-12 20:59
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

type ArticleController struct {
	BaseController
}

// Create  添加文章信息
func (a *ArticleController) Create() {
	article := form.AddArticle{}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Error("json unmarshal add article err:%s", err.Error())
		a.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}

	err = a.ValidBody(&article)
	if err != nil {
		if err != nil {
			a.ErrorCode(1, http.StatusForbidden, err.Error())
		}
	}
	exists := services.NewArticle().ExistByTitle(article.Title, 0)
	if exists {
		logs.Error("add article err:%s", "title already exists")
		a.ErrorCode(0, http.StatusOK, fmt.Sprintf("title`%s` already exists", article.Title))
	}

	ID, err := services.NewArticle().Create(article)
	if err != nil {
		logs.Error("add article err:%s", err.Error())
		a.ErrorCode(1, http.StatusInternalServerError, "add article err")
	}
	var InsertID = struct {
		ID uint `json:"id"`
	}{ID: ID}

	a.Correct(InsertID)
}

// Update 更新文章信息
func (a *ArticleController) Update() {
	article := form.UpdateArticle{}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Error("json unmarshal update article err:%s", err.Error())
		a.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}

	err = a.ValidBody(&article)
	if err != nil {
		if err != nil {
			a.ErrorCode(1, http.StatusForbidden, err.Error())
		}
	}
	exists := services.NewArticle().ExistByTitle(article.Title, article.ID)
	if exists {
		logs.Error("update article err:%s", "title already exists")
		a.ErrorCode(0, http.StatusOK, fmt.Sprintf("title`%s` already exists", article.Title))
	}
	err = services.NewArticle().Update(article)
	if err != nil {
		logs.Error("update article err:%s", err.Error())
		a.ErrorCode(1, http.StatusInternalServerError, "update article err")
	}
	a.Correct("")
}

// Delete 删除文章信息
func (a *ArticleController) Delete() {
	article := form.DeleteArticle{}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Error("json unmarshal delete article err:%s", err.Error())
		a.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}
	err = a.ValidBody(&article)
	if err != nil {
		a.ErrorCode(1, http.StatusForbidden, err.Error())
	}

	ids := strings.Split(article.ID, ",")
	if len(ids) == 1 && ids[0] == "" {
		a.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}
	l := make([]int, len(ids))
	for i, v := range ids {
		ID, err := strconv.Atoi(v)
		if err != nil {
			a.ErrorCode(1, http.StatusBadRequest, "invalid param")
		}
		l[i] = ID
	}
	err = services.NewArticle().Delete(l)
	if err != nil {
		logs.Error("delete article info err:%s", err.Error())
		a.ErrorCode(1, http.StatusInternalServerError, "delete article err")
	}
	a.Correct("")
}

// List 文章信息列表
func (a *ArticleController) List() {
	params := form.ArticleQueryParam{
		Title:       a.GetString("title"),
		Description: a.GetString("description"),
		Status:      a.GetString("status", "all"),
	}
	result, err := services.NewArticle().List(params, form.ArticleQueryOptions{a.GetPaginationParam()})
	if err != nil {
		a.ErrorCode(1, http.StatusInternalServerError, err.Error())
	}
	a.Correct(result)
}

// Get 获取文章详情
func (a *ArticleController) Get() {
	ID, _ := a.GetInt(":id")
	if ID == 0 {
		a.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}
	result, err := services.NewArticle().Get(uint(ID))
	if err != nil {
		logs.Error("article  get err:%s", err.Error())
		a.ErrorCode(1, http.StatusInternalServerError, err.Error())
	}
	a.Correct(result)
}

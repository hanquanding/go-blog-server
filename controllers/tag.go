/**
 * @author: hqd
 * @description: tag controller
 * @file: tag
 * @date: 2021-02-12 19:39
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

type TagController struct {
	BaseController
}

// List 标签列表
func (t *TagController) List() {
	params := form.TagQueryParam{
		TagName:   t.GetString("tag_name"),
		TagStatus: t.GetString("tag_status", "all"),
	}
	result, err := services.NewTag().List(params, form.TagQueryOptions{t.GetPaginationParam()})
	if err != nil {
		t.ErrorCode(1, http.StatusInternalServerError, err.Error())
	}
	t.Correct(result)
}

// Get  获取标签详情
func (t *TagController) Get() {
	tagID, _ := t.GetInt(":id")
	if tagID == 0 {
		t.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}

	result, err := services.NewTag().Get(uint(tagID))
	if err != nil {
		logs.Error("tag get err:%s", err.Error())
		t.ErrorCode(1, http.StatusInternalServerError, err.Error())
	}
	t.Correct(result)
}

// Create  添加标签信息
func (t *TagController) Create() {
	tag := form.AddTag{}
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &tag)
	if err != nil {
		logs.Error("json unmarshal add tag err:%s", err.Error())
		t.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}

	err = t.ValidBody(&tag)
	if err != nil {
		t.ErrorCode(1, http.StatusForbidden, err.Error())
	}
	exists := services.NewTag().ExistByTagName(tag.TagName, 0)
	if exists {
		logs.Error("add tag info err:%s", "tag name already exists")
		t.ErrorCode(0, http.StatusOK, fmt.Sprintf("tag name `%s` already exists", tag.TagName))
	}
	TagID, err := services.NewTag().Create(tag)
	if err != nil {
		logs.Error("add tag err:%s", err.Error())
		t.ErrorCode(1, http.StatusInternalServerError, "add tag fail")
	}
	var InsertID = struct {
		ID uint `json:"id"`
	}{ID: TagID}

	t.Correct(InsertID)
}

// Update  更新标签信息
func (t *TagController) Update() {
	tag := form.UpdateTag{}
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &tag)
	if err != nil {
		logs.Error("json unmarshal update tag err:%s", err.Error())
		t.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}

	err = t.ValidBody(&tag)
	if err != nil {
		t.ErrorCode(1, http.StatusForbidden, err.Error())
	}

	exists := services.NewTag().ExistByTagName(tag.TagName, tag.ID)
	if exists {
		logs.Error("update tag info err:%s", "tag name already exists")
		t.ErrorCode(0, http.StatusOK, fmt.Sprintf("tag name `%s` already exists", tag.TagName))
	}
	err = services.NewTag().Update(tag)
	if err != nil {
		logs.Error("update tag info err:%s", err.Error())
		t.ErrorCode(1, http.StatusInternalServerError, "update tag fail")
	}

	t.Correct("")
}

// Delete 删除标签
func (t *TagController) Delete() {
	tag := form.DeleteTag{}
	err := json.Unmarshal(t.Ctx.Input.RequestBody, &tag)
	if err != nil {
		logs.Error("json unmarshal delete tag err:%s", err.Error())
		t.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}
	err = t.ValidBody(&tag)
	if err != nil {
		t.ErrorCode(1, http.StatusForbidden, err.Error())
	}

	ids := strings.Split(tag.ID, ",")
	if len(ids) == 1 && ids[0] == "" {
		t.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}
	l := make([]int, len(ids))
	for i, v := range ids {
		ID, err := strconv.Atoi(v)
		if err != nil {
			t.ErrorCode(1, http.StatusBadRequest, "invalid param")
		}
		l[i] = ID
	}
	err = services.NewTag().Delete(l)
	if err != nil {
		logs.Error("delete tag info err:%s", err.Error())
		t.ErrorCode(1, http.StatusInternalServerError, "delete tag fail")
	}
	t.Correct("")
}

// UpdateTagStatus 更新标签状态
func (t *TagController) UpdateTagStatus() {
	tagID, _ := t.GetInt(":id")
	if tagID == 0 {
		t.ErrorCode(1, http.StatusBadRequest, "invalid param")
	}
	err := services.NewTag().UpdateTagStatus(uint(tagID))
	if err != nil {
		logs.Error("update tag info err:%s", err.Error())
		t.ErrorCode(1, http.StatusInternalServerError, "update tag fail")
	}

	t.Correct("")
}

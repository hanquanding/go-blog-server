/**
 * @author: hqd
 * @description: base controller
 * @file: base
 * @date: 2021-02-07 10:17
 */

package controllers

import (
	"errors"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	"github.com/beego/beego/v2/server/web"
	"github.com/hqd8080/go-blog-server/models/form"
)

type BaseController struct {
	web.Controller
	valid validation.Validation
}

func (base *BaseController) Prepare() {
	logs.Info("prepare.")
}

func (base *BaseController) ErrorCode(code, statusCode int, msg string) {
	var result = struct {
		Ret int    `json:"ret"`
		Msg string `json:"msg"`
	}{code, msg}

	base.Data["json"] = result
	base.Ctx.ResponseWriter.WriteHeader(statusCode)
	base.ServeJSON()
	base.StopRun()
}

func (base *BaseController) Correct(data interface{}) {
	var ok bool
	var ret map[string]interface{}
	if ret, ok = data.(map[string]interface{}); !ok {
		ret = make(map[string]interface{})
		ret["data"] = data
	}
	ret["ret"] = 0
	ret["msg"] = "success"
	base.Data["json"] = ret
	base.ServeJSON()
	base.StopRun()
}

func (base *BaseController) ValidBody(i interface{}) error {
	ok, err := base.valid.Valid(i)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(base.valid.Errors[0].Key + base.valid.Errors[0].Message)
	}
	return nil
}

func (base *BaseController) GetPaginationParam() *form.PaginationParam {
	return &form.PaginationParam{
		PageNum:  base.GetPageNum(),
		PageSize: base.GetPageSize(),
	}
}

func (base *BaseController) GetPageNum() int {
	var pageNum = 1
	i, err := base.GetInt("pagenum")
	if err != nil {
		return pageNum
	}
	return i
}

func (base *BaseController) GetPageSize() int {
	var pageSize = 10
	i, err := base.GetInt("pagesize")
	if err != nil {
		return pageSize
	}
	return i
}

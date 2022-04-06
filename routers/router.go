/**
 * @author: hqd
 * @description: router
 * @file: router
 * @date: 2021-02-07 10:05
 */

package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/hqd8080/go-blog-server/controllers"
)

const API string = "/api/v1"

func init() {
	web.Router(API+"/website", &controllers.WebsiteController{}, "get:List")
	web.Router(API+"/website", &controllers.WebsiteController{}, "post:Create")
	web.Router(API+"/website", &controllers.WebsiteController{}, "put:Update")
	web.Router(API+"/website", &controllers.WebsiteController{}, "delete:Delete")
	web.Router(API+"/website/:id", &controllers.WebsiteController{}, "get:Get")
	web.Router(API+"/website/status/:id", &controllers.WebsiteController{}, "put:UpdateWebStatus")

	web.Router(API+"/tag", &controllers.TagController{}, "get:List")
	web.Router(API+"/tag", &controllers.TagController{}, "post:Create")
	web.Router(API+"/tag", &controllers.TagController{}, "put:Update")
	web.Router(API+"/tag", &controllers.TagController{}, "delete:Delete")
	web.Router(API+"/tag/:id", &controllers.TagController{}, "get:Get")
	web.Router(API+"/tag/status/:id", &controllers.TagController{}, "put:UpdateTagStatus")

	web.Router(API+"/category", &controllers.CategoryController{}, "get:List")
	web.Router(API+"/category", &controllers.CategoryController{}, "post:Create")
	web.Router(API+"/category", &controllers.CategoryController{}, "put:Update")
	web.Router(API+"/category", &controllers.CategoryController{}, "delete:Delete")
	web.Router(API+"/category/:id", &controllers.CategoryController{}, "get:Get")

	web.Router(API+"/article", &controllers.ArticleController{}, "get:List")
	web.Router(API+"/article", &controllers.ArticleController{}, "post:Create")
	web.Router(API+"/article", &controllers.ArticleController{}, "put:Update")
	web.Router(API+"/article", &controllers.ArticleController{}, "delete:Delete")
	web.Router(API+"/article/:id", &controllers.ArticleController{}, "get:Get")
}

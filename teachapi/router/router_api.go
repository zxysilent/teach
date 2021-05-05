package router

import (
	"teach/control"

	"github.com/labstack/echo/v4"
)

// apiRouter 通用访问
func apiRouter(api *echo.Group) {
	api.GET("/sys/info", control.SysInfo)
	// auth
	api.POST("/auth/login", control.AuthLogin)
	// upload
	api.POST("/upload/file", control.UploadFile)
	api.POST("/upload/image", control.UploadImage)
	// user
	api.GET("/user/get", control.UserGet)
	api.GET("/user/page", control.UserPage)
	// class
	api.GET("/class/all", control.ClassAll)
	api.GET("/class/get", control.ClassGet)
	api.GET("/class/page", control.ClassPage)
	// article
	api.GET("/article/get", control.ArticleGet)
	api.GET("/article/page", control.ArticlePage)
	api.GET("/article/edit/hits", control.ArticleEditHits)
}

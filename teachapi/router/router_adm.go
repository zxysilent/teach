package router

import (
	"teach/control"

	"github.com/labstack/echo/v4"
)

// admRouter 需要登录才能访问
func admRouter(adm *echo.Group) {
	// auth
	adm.GET("/auth/get", control.AuthGet)
	adm.POST("/auth/edit/info", control.AuthEditInfo)
	adm.POST("/auth/edit/passwd", control.AuthEditPasswd)
	// user
	adm.POST("/user/add", control.UserAdd)
	adm.POST("/user/edit", control.UserEdit)
	adm.POST("/user/drop", control.UserDrop)
	// class
	adm.POST("/class/add", control.ClassAdd)
	adm.POST("/class/edit", control.ClassEdit)
	adm.POST("/class/drop", control.ClassDrop)
	// article
	adm.POST("/article/add", control.ArticleAdd)
	adm.POST("/article/edit", control.ArticleEdit)
	adm.POST("/article/drop", control.ArticleDrop)
}

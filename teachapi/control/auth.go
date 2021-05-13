package control

import (
	"teach/conf"
	"teach/internal/hmt"
	"teach/model"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zxysilent/logs"
	"github.com/zxysilent/utils"
)

// Login doc
// @Tags auth
// @Summary 登陆
// @Accept mpfd
// @Param num formData string true "账号" default(admin)
// @Param passwd formData string true "密码" default(12348765)
// @Success 200 {object} model.Reply{data=string} "返回数据"
// @Router /api/auth/login [post]
func AuthLogin(ctx echo.Context) error {
	ipt := struct {
		Num    string `json:"num" form:"num"`
		Passwd string `json:"passwd" form:"passwd"`
	}{}
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("请输入用户名和密码", err.Error()))
	}
	if ipt.Num == "" && len(ipt.Num) > 18 {
		return ctx.JSON(utils.ErrIpt("请输入正确的用户名"))
	}
	mod, err := model.UserLogin(ipt.Num)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("用户名输入错误"))
	}
	if mod.Passwd != ipt.Passwd {
		return ctx.JSON(utils.Fail("密码输入错误"))
	}
	auth := hmt.HmtAuth{
		Id:    mod.Id,
		Num:   mod.Num,
		ExpAt: time.Now().Add(time.Hour * time.Duration(conf.App.HmtExp)).Unix(),
	}
	token := auth.Encode(conf.App.HmtKey)
	logs.Info("登陆成功", auth.Id, token)
	return ctx.JSON(utils.Succ("登陆成功", token))
}

// Auth doc
// @Tags auth
// @Summary 登录信息
// @Param token query string false "凭证"
// @Param Authorization header string false "凭证"
// @Success 200 {object} model.Reply{data=model.User} "返回数据"
// @Router /adm/auth/get [get]
func AuthGet(ctx echo.Context) error {
	mod, _ := model.UserGet(ctx.Get("uid").(int))
	return ctx.JSON(utils.Succ("succ", mod))
}

// AuthEditInfo doc
// @Tags auth
// @Summary 修改自己的信息
// @Param token query string false "凭证"
// @Param Authorization header string false "凭证"
// @Success 200 {object} model.Reply "返回数据"
// @Param body body model.User true "request"
// @Router /adm/auth/edit/info [post]
func AuthEditInfo(ctx echo.Context) error {
	ipt := &model.User{}
	err := ctx.Bind(&ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("输入数据有误", err.Error()))
	}
	ipt.Id = ctx.Get("uid").(int)
	if err := model.UserEdit(ipt); err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// AuthEditPasswd doc
// @Tags auth
// @Summary 修改自己的密码
// @Param token query string false "凭证"
// @Param Authorization header string false "凭证"
// @Param opass formData string true "旧密码"
// @Param npass formData string true "新密码"
// @Success 200 {object} model.Reply "返回数据"
// @Router /adm/auth/edit/passwd [post]
func AuthEditPasswd(ctx echo.Context) error {
	ipt := &struct {
		Opass string `json:"opass" form:"opass"`
		Npass string `json:"npass" form:"npass"`
	}{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("输入数据有误", err.Error()))
	}
	mod, err := model.UserGet(ctx.Get("uid").(int))
	if err != nil {
		return ctx.JSON(utils.Fail("输入数据有误,请重试"))
	}
	if mod.Passwd != ipt.Opass {
		return ctx.JSON(utils.Fail("原始密码输入错误,请重试"))
	}
	mod.Passwd = ipt.Npass
	if err := model.UserEdit(&mod); err != nil {
		return ctx.JSON(utils.Fail("密码修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

package control

import (
	"teach/model"

	"github.com/labstack/echo/v4"
	"github.com/zxysilent/utils"
)

// UserGet doc
// @Tags user
// @Summary 通过id获取user信息
// @Param id query int true "id"
// @Success 200 {object} model.Reply{data=model.User} "成功数据"
// @Router /api/user/get [get]
func UserGet(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.UserGet(ipt.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到用户信息", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ", mod))
}

// UserPage doc
// @Tags user
// @Summary 分页数据
// @Param pi query int true "分页数"  default(1)
// @Param ps query int true "每页条数[5,30]" default(8)
// @Success 200 {object} model.Reply{data=[]model.User} "成功数据"
// @Router /api/user/page [get]
func UserPage(ctx echo.Context) error {
	// cid, err := strconv.Atoi(ctx.Param("cid"))
	// if err != nil {
	// 	return ctx.JSON(utils.ErrIpt("数据输入错误", err.Error()))
	// }
	ipt := &model.Page{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	err = ipt.Stat()
	if err != nil {
		return ctx.JSON(utils.ErrIpt("分页信息输入错误", err.Error()))
	}
	count, _ := model.UserCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据", " count < 1"))
	}
	mods, err := model.UserPage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("查询数据错误", err.Error()))
	}
	if len(mods) < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据", "len(mods) < 1"))
	}
	return ctx.JSON(utils.Page("succ", mods, int(count)))
}

// UserAdd doc
// @Tags user
// @Summary 添加user信息
// @Param token query string true "凭证"
// @Param body body model.User true "request"
// @Success 200 {object} model.Reply "成功数据"
// @Router /adm/user/add [post]
func UserAdd(ctx echo.Context) error {
	ipt := &model.User{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	err = model.UserAdd(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// UserEdit doc
// @Tags user
// @Summary 修改user信息
// @Param token query string true "凭证"
// @Param body body model.User true "request"
// @Success 200 {object} model.Reply "成功数据"
// @Router /adm/user/edit [post]
func UserEdit(ctx echo.Context) error {
	ipt := &model.User{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	err = model.UserEdit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// UserDrop doc
// @Tags user
// @Summary 删除user信息
// @Param body body model.IptId true "请求数据"
// @Param token query string true "凭证"
// @Success 200 {object} model.Reply "成功数据"
// @Router /adm/user/drop [post]
func UserDrop(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.UserGet(ipt.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", "不存在当前数据"))
	}
	err = model.UserDrop(mod.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

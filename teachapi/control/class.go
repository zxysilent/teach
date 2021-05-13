package control

import (
	"teach/model"

	"github.com/labstack/echo/v4"
	"github.com/zxysilent/utils"
)

// ClassGet doc
// @Tags class
// @Summary 通过id获取class信息
// @Param id query int true "id"
// @Success 200 {object} model.Reply{data=model.Class} "返回数据"
// @Router /api/class/get [get]
func ClassGet(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.ClassGet(ipt.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到分类信息", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ", mod))
}

// ClassPage doc
// @Tags class
// @Summary 分页数据
// @Param pi query int true "分页数"  default(1)
// @Param ps query int true "每页条数[5,30]" default(8)
// @Success 200 {object} model.Reply{data=[]model.Class} "返回数据"
// @Router /api/class/page [get]
func ClassPage(ctx echo.Context) error {
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
	count, _ := model.ClassCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据", " count < 1"))
	}
	mods, err := model.ClassPage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("查询数据错误", err.Error()))
	}
	if len(mods) < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据", "len(mods) < 1"))
	}
	return ctx.JSON(utils.Page("succ", mods, int(count)))
}

// ClassAll doc
// @Tags class
// @Summary 所有分类
// @Success 200 {object} model.Reply{data=[]model.Class} "返回数据"
// @Router /api/class/all [get]
func ClassAll(ctx echo.Context) error {
	mods, err := model.ClassAll()
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到数据", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ", mods))
}

// ClassAdd doc
// @Tags class
// @Summary 添加class信息
// @Param token query string true "凭证"
// @Param body body model.Class true "请求数据"
// @Success 200 {object} model.Reply "返回数据"
// @Router /adm/class/add [post]
func ClassAdd(ctx echo.Context) error {
	ipt := &model.Class{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	err = model.ClassAdd(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// ClassEdit doc
// @Tags class
// @Summary 修改class信息
// @Param token query string true "凭证"
// @Param body body model.Class true "请求数据"
// @Success 200 {object} model.Reply "返回数据"
// @Router /adm/class/edit [post]
func ClassEdit(ctx echo.Context) error {
	ipt := &model.Class{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	err = model.ClassEdit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// ClassDrop doc
// @Tags class
// @Summary 删除class信息
// @Param body body model.IptId true "请求数据"
// @Param token query string true "凭证"
// @Success 200 {object} model.Reply "返回数据"
// @Router /adm/class/drop [post]
func ClassDrop(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.ClassGet(ipt.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", "不存在当前数据"))
	}
	err = model.ClassDrop(mod.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

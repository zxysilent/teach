package control

import (
	"teach/model"

	"github.com/labstack/echo/v4"
	"github.com/zxysilent/utils"
)

// ArticleGet doc
// @Tags article
// @Summary 通过id获取article信息
// @Param id query int true "id"
// @Success 200 {object} model.Reply{data=model.Article} "返回数据"
// @Router /api/article/get [get]
func ArticleGet(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.ArticleGet(ipt.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("未查询到用户信息", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ", mod))
}

// ArticlePage doc
// @Tags article
// @Summary 分页数据
// @Param pi query int true "分页数"  default(1)
// @Param ps query int true "每页条数[5,30]" default(8)
// @Success 200 {object} model.Reply{data=[]model.Article} "返回数据"
// @Router /api/article/page [get]
func ArticlePage(ctx echo.Context) error {
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
	count, _ := model.ArticleCount()
	if count < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据", " count < 1"))
	}
	mods, err := model.ArticlePage(ipt.Pi, ipt.Ps)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("查询数据错误", err.Error()))
	}
	if len(mods) < 1 {
		return ctx.JSON(utils.ErrOpt("未查询到数据", "len(mods) < 1"))
	}
	return ctx.JSON(utils.Page("succ", mods, int(count)))
}

// ArticleAdd doc
// @Tags article
// @Summary 添加article信息
// @Param token query string true "凭证"
// @Param body body model.Article true "请求数据"
// @Success 200 {object} model.Reply "返回数据"
// @Router /adm/article/add [post]
func ArticleAdd(ctx echo.Context) error {
	ipt := &model.Article{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	err = model.ArticleAdd(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("添加失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// ArticleEdit doc
// @Tags article
// @Summary 修改article信息
// @Param token query string true "凭证"
// @Param body body model.Article true "请求数据"
// @Success 200 {object} model.Reply "返回数据"
// @Router /adm/article/edit [post]
func ArticleEdit(ctx echo.Context) error {
	ipt := &model.Article{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	err = model.ArticleEdit(ipt)
	if err != nil {
		return ctx.JSON(utils.Fail("修改失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

// ArticleEditHits doc
// @Tags article
// @Summary 修改article信息hits
// @Param id query int true "id"
// @Success 200 {object} model.Reply{data=int} "返回数据"
// @Router /api/article/edit/hits [get]
func ArticleEditHits(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.ArticleGet(ipt.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("修改失败", "不存在当前数据"))
	}
	err = model.ArticleEditHits(mod.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("修改失败", err.Error()))
	}
	mod.Hits++
	return ctx.JSON(utils.Succ("succ", mod.Hits))
}

// ArticleDrop doc
// @Tags article
// @Summary 删除article信息
// @Param body body model.IptId true "请求数据"
// @Param token query string true "凭证"
// @Success 200 {object} model.Reply "返回数据"
// @Router /adm/article/drop [post]
func ArticleDrop(ctx echo.Context) error {
	ipt := &model.IptId{}
	err := ctx.Bind(ipt)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	mod, err := model.ArticleGet(ipt.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", "不存在当前数据"))
	}
	err = model.ArticleDrop(mod.Id)
	if err != nil {
		return ctx.JSON(utils.ErrOpt("删除失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("succ"))
}

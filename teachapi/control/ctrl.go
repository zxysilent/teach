package control

import (
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"teach/conf"
	"teach/model"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nfnt/resize"
	"github.com/zxysilent/utils"
	_ "golang.org/x/image/bmp"
)

// UploadFile doc
// @Tags ctrl
// @Summary 上传文件
// @Accept  mpfd
// @Param file formData file true "file"
// @Success 200 {object} model.Reply "成功数据"
// @Router /api/upload/file [post]
func UploadFile(ctx echo.Context) error {
	f, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(utils.Fail("上传失败", err.Error()))
	}
	src, err := f.Open()
	if err != nil {
		return ctx.JSON(utils.Fail("文件打开失败", err.Error()))

	}
	dir := time.Now().Format("200601/02")
	os.MkdirAll("./static/upload/"+dir[:6], 0666)
	name := "static/upload/" + dir + utils.RandStr(10) + path.Ext(f.Filename)
	dst, err := os.Create(name)
	if err != nil {
		return ctx.JSON(utils.Fail("文件打创建文件失败", err.Error()))

	}
	_, err = io.Copy(dst, src)
	if err != nil {
		return ctx.JSON(utils.Fail("文件保存失败", err.Error()))
	}
	src.Close()
	dst.Close()
	return ctx.JSON(utils.Succ("上传成功", "/"+name))
}

// UploadImage doc
// @Tags ctrl
// @Summary 上传图片
// @Accept  mpfd
// @Param file formData file true "file"
// @Success 200 {object} model.Reply "成功数据"
// @Router /api/upload/image [post]
func UploadImage(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(utils.ErrIpt(`未发现文件,请重试`, err.Error()))
	}
	if !strings.Contains(file.Header.Get("Content-Type"), "image") {
		return ctx.JSON(utils.NewErrIpt("请选择图片文件", file.Header.Get("Content-Type")))
	}
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(utils.ErrIpt(`文件打开失败,请重试`, err.Error()))
	}
	defer src.Close()
	dir := time.Now().Format("200601/02")
	os.MkdirAll("./static/upload/"+dir[:6], 0666)
	name := "static/upload/" + dir + utils.RandStr(10) + ".jpg"
	dst, err := os.Create(name)
	if err != nil {
		return ctx.JSON(utils.ErrIpt(`目标文件创建失败,请重试`, err.Error()))
	}
	defer dst.Close()
	imgSrc, _, err := image.Decode(src)
	// 图片解码
	if err != nil {
		return ctx.JSON(utils.ErrIpt(`读取图片失败,请重试`, err.Error()))
	}
	if conf.App.ImageCut { //图片裁剪
		// 宽度>指定宽度 防止负调整
		dx := imgSrc.Bounds().Dx()
		if dx > conf.App.ImageWidth {
			imgSrc = resize.Resize(uint(conf.App.ImageWidth), 0, imgSrc, resize.Lanczos3)
		}
		//高度>指定高度 防止负调整
		dy := imgSrc.Bounds().Dy()
		if dy > conf.App.ImageHeight {
			imgSrc = resize.Resize(0, uint(conf.App.ImageHeight), imgSrc, resize.Lanczos3)
		}
	}
	err = jpeg.Encode(dst, imgSrc, nil)
	if err != nil {
		return ctx.JSON(utils.ErrIpt(`文件写入失败,请重试`, err.Error()))
	}
	return ctx.JSON(utils.Succ(`文件上传成功`, "/"+name))
}

// SysInfo doc
// @Tags ctrl
// @Summary 系统信息
// @Success 200 {object} model.Reply{data=model.SysInfo} "成功数据"
// @Router /api/sys/info [get]
func SysInfo(ctx echo.Context) error {
	state := struct {
		ARCH    string `json:"arch"`
		OS      string `json:"os"`
		Version string `json:"version"`
		NumCPU  int    `json:"num_cpu"`
	}{
		ARCH:    runtime.GOARCH,
		OS:      runtime.GOOS,
		Version: runtime.Version(),
		NumCPU:  runtime.NumCPU(),
	}
	return ctx.JSON(utils.Succ(`系统信息`, state))
}

// BannerAll doc
// @Tags banner
// @Summary 所有轮播
// @Success 200 {object} model.Reply{data=[]model.Banner} "成功数据"
// @Router /api/banner/all [get]
func BannerAll(ctx echo.Context) error {
	mods := make([]model.Banner, 0, 4)
	for idx, val := range _arr {
		mods = append(mods, model.Banner{
			Id:    idx + 1,
			Title: "测试图片",
			Url:   val,
			Cunix: time.Now().Unix(),
		})
	}
	return ctx.JSON(utils.Succ("succ", mods))
}

var _arr = []string{"/static/imgs/0.jpg", "/static/imgs/1.jpg", "/static/imgs/2.jpg", "/static/imgs/3.jpg", "/static/imgs/4.jpg", "/static/imgs/5.jpg"}

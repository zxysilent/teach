package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/zxysilent/logs"
)

type appconf struct {
	Title   string `toml:"title"`
	Explain string `toml:"explain"`
	Mode    string `toml:"mode"`
	Addr    string `toml:"addr"`
	Srv     string `toml:"srv"`
	HmtKey  string `toml:"hmt_key"`
	HmtExp  int    `toml:"hmt_exp"`
	Author  struct {
		Name    string `toml:"name"`
		Website string `toml:"website"`
	} `toml:"author"`
	Wechat struct {
		Appid  string `toml:"appid"`
		Secret string `toml:"secret"`
	} `toml:"wechat"`
	ImageCut    bool `toml:"image_cut"`
	ImageWidth  int  `toml:"image_width"`
	ImageHeight int  `toml:"image_height"`
	PageMin     int  `toml:"page_min"`
	PageMax     int  `toml:"page_max"`
	DbIdle      int  `toml:"db_idle"`
	DbOpen      int  `toml:"db_open"`
}

func (app *appconf) IsProd() bool {
	return app.Mode == "prod"
}
func (app *appconf) IsDev() bool {
	return app.Mode == "dev"
}

var (
	App       *appconf
	defConfig = "./conf/conf.toml"
)

func Init() {
	var err error
	App, err = initConf()
	if err != nil {
		logs.Fatal("config init error : ", err.Error())
	}
	logs.Debug("conf init")
}

func initConf() (*appconf, error) {
	app := &appconf{}
	_, err := toml.DecodeFile(defConfig, &app)
	if err != nil {
		return nil, err
	}
	return app, nil
}

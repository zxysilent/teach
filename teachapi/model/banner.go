package model

// Banner 文章
type Banner struct {
	Id    int    `json:"id"`    //主键
	Title string `json:"title"` //标题
	Url   string `json:"url"`   //图片
	Cunix int64  `json:"cunix"` //创建时间
}

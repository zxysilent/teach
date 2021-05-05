package model

// Article 文章
type Article struct {
	Id      int    `json:"id"`           //主键
	Cid     int    `json:"cid"`          //cls id
	Class   *Class `db:"-" json:"class"` //栏目
	Title   string `json:"title"`        //标题
	Author  string `json:"author"`       //作者
	Hits    int    `json:"hits"`         //点击数
	Content string `json:"content"`      //详细
	Uunix   int64  `json:"uunix"`        //修改时间
	Cunix   int64  `json:"cunix"`        //创建时间
}

// ArticleGet 查询一条新闻
func ArticleGet(id int) (*Article, error) {
	mod := &Article{}
	err := Db.Get(mod, "SELECT * FROM Article WHERE id = ?", id)
	return mod, err
}

// ArticleAll 查询所有新闻
func ArticleAll() ([]Article, error) {
	mods := make([]Article, 0, 4)
	err := Db.Select(&mods, "SELECT * FROM Article")
	for idx := range mods {
		mods[idx].Class, _ = ClassGet(mods[idx].Cid)
	}
	return mods, err
}

// ArticlePage 分页查询新闻
func ArticlePage(pi int, ps int) ([]Article, error) {
	mods := make([]Article, 0, 4)
	err := Db.Select(&mods, "SELECT * FROM Article LIMIT ? OFFSET ?", ps, (pi-1)*ps)
	for idx := range mods {
		mods[idx].Class, _ = ClassGet(mods[idx].Cid)
	}
	return mods, err
}

// ArticleCount 查询新闻总数
func ArticleCount() (int, error) {
	var count int
	err := Db.Get(&count, "SELECT count(id) as count FROM Article")
	return count, err
}

// ArticleAdd 添加新闻
func ArticleAdd(mod *Article) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.NamedExec("INSERT INTO Article(cid,title,author,hits,content,uunix,cunix) VALUES(:cid,:title,:author,:hits,:content,:uunix,:cunix)", mod)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// ArticleEdit 编辑新闻
func ArticleEdit(mod *Article) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.NamedExec("UPDATE Article SET cid=:cid,title=:title,author=:author,hits=:hits,content=:content,uunix=:uunix WHERE id=:id", mod)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// ArticleEditHits 编辑新闻点击量
func ArticleEditHits(id int) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.Exec("UPDATE article SET hits=hits+1 WHERE id=?", id)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// ArticleDrop 删除新闻
func ArticleDrop(id int) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.Exec(`DELETE FROM Article WHERE id=?`, id)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

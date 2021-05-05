package model

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// Class 类别
type Class struct {
	Id   int    `json:"id"`   //主键
	Name string `json:"name"` //类别名称
	Url  string `json:"url"`  //图片url
}

// ClassGet 查询一条数据
func ClassGet(id int) (*Class, error) {
	mod := &Class{}
	err := Db.Get(mod, "SELECT * FROM class WHERE id = ?", id)
	return mod, err
}

// ClassAll 查询所有数据
func ClassAll() ([]Class, error) {
	mods := make([]Class, 0, 4)
	err := Db.Select(&mods, "SELECT * FROM class")
	return mods, err
}

// ClassPage 分页查询数据
func ClassPage(pi int, ps int) ([]Class, error) {
	mods := make([]Class, 0, 4)
	err := Db.Select(&mods, "SELECT * FROM class LIMIT ? OFFSET ?", ps, (pi-1)*ps)
	return mods, err
}

// ClassCount 查询数据总数
func ClassCount() (int, error) {
	var count int
	err := Db.Get(&count, "SELECT count(id) as count FROM class")
	return count, err
}

// ClassAdd 添加数据
func ClassAdd(mod *Class) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.NamedExec("INSERT INTO class(name,url) VALUES(:name,:url)", mod)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// ClassEdit 编辑数据
func ClassEdit(mod *Class) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.NamedExec("UPDATE class SET name=:name,url=:url WHERE id=:id", mod)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// ClassDrop 删除数据
func ClassDrop(id int) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.Exec(`DELETE FROM class WHERE id=?`, id)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// ClassIds 通过ids返回数据集合
func ClassIds(ids []int) (map[int]*Class, error) {
	mods := make([]Class, 0, 4)
	query, args, err := sqlx.In("SELECT * FROM class WHERE id IN (?);", ids)
	if err != nil {
		return nil, err
	}
	query = Db.Rebind(query)
	Db.Select(&mods, query, args...)
	if len(mods) > 0 {
		mapMods := make(map[int]*Class, len(mods))
		for idx := range mods {
			mapMods[mods[idx].Id] = &mods[idx]
		}
		return mapMods, nil
	}
	return nil, errors.New("app: no rows in result SET")
}

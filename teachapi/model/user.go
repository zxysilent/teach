package model

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// User 管理用户
type User struct {
	Id     int    `json:"id"`    //主键
	Num    string `json:"num"`   //账号
	Name   string `json:"name"`  //名称
	Passwd string `json:"-"`     //密码
	Phone  string `json:"phone"` //手机
	Email  string `json:"email"` //邮箱
	Cunix  int64  `json:"cunix"` //创建时间
}

// UserGet 查询一条用户
func UserGet(id int) (User, error) {
	mod := User{}
	err := Db.Get(&mod, "SELECT * FROM user WHERE id = ?", id)
	return mod, err
}

// UserLogin 用户登录
func UserLogin(num string) (User, error) {
	mod := User{}
	err := Db.Get(&mod, "SELECT * FROM user WHERE num = ?", num)
	return mod, err
}

// UserAll 查询所有用户
func UserAll() ([]User, error) {
	mods := make([]User, 0, 4)
	err := Db.Select(&mods, "SELECT * FROM user")
	return mods, err
}

// UserPage 分页查询用户
func UserPage(pi int, ps int) ([]User, error) {
	mods := make([]User, 0, 4)
	err := Db.Select(&mods, "SELECT * FROM user LIMIT ? OFFSET ?", ps, (pi-1)*ps)
	return mods, err
}

// UserCount 查询用户总数
func UserCount() (int, error) {
	var count int
	err := Db.Get(&count, "SELECT COUNT(id) as count FROM user")
	return count, err
}

// UserAdd 添加用户
func UserAdd(mod *User) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.NamedExec("INSERT INTO user(num,name,passwd,phone,email,cunix) VALUES(:num,:name,:passwd,:phone,:email,:cunix)", mod)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// UserEdit 编辑用户
func UserEdit(mod *User) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.NamedExec("UPDATE user SET num=:num,name=:name,phone=:phone,email=:email WHERE id=:id", mod)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// UserEdit 编辑用户密码
func UserEditPasswd(mod *User) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.NamedExec("UPDATE user SET passwd=:passwd WHERE id=:id", mod)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// UserDrop 删除用户
func UserDrop(id int) error {
	Tx, err := Db.Beginx()
	if err != nil {
		return err
	}
	_, err = Tx.Exec(`DELETE FROM user WHERE id=?`, id)
	if err != nil {
		Tx.Rollback()
		return err
	}
	Tx.Commit()
	return nil
}

// UserIds 通过ids返回用户集合
func UserIds(ids []int) (map[int]*User, error) {
	mods := make([]User, 0, 4)
	query, args, err := sqlx.In("SELECT * FROM user WHERE id IN (?);", ids)
	if err != nil {
		return nil, err
	}
	query = Db.Rebind(query)
	Db.Select(&mods, query, args...)
	if len(mods) > 0 {
		mapMods := make(map[int]*User, len(mods))
		for idx := range mods {
			mapMods[mods[idx].Id] = &mods[idx]
		}
		return mapMods, nil
	}
	return nil, errors.New("app: no rows in result SET")
}

package model

import (
	"blog/service"
)

type UserModel struct {}

var (
	User = &UserModel {}
	db = service.GetIns().GetMysqlDB()
)

/**
* 查询用户
*/
func (u *UserModel) FindUser() (map[int]map[string]string) {
	// 获取db实例
	res := db.Table("users").Where("username = ?", "admin").Find(&User)
	rows, _ := res.Rows()
	return service.GetAllRows(rows)
}
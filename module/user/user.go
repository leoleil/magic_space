package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"github.com/leoleil/magic_space/common/database"
)

type MsSysUser struct {
	Id int64 `json:"id"`
	GmtCreate string `json:"gmt_create"`
	Psw string `json:"psw"`
	GmtModified string `json:"gmt_modified"`
	Username string `json:"username"`
	Key sql.NullString `json:"key"`
	Email sql.NullString `json:"email"`
}
func QueryUserById(id int64)(msSysUser MsSysUser,err error)  {
	handle := database.GetHandle()
	db,ok := handle.InitDB()
	if !ok{
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	err = db.QueryRow("SELECT `id`,`Gmt_Create`,`Psw`,`Gmt_Modified`,`Username`,`Key`,`Email` FROM ms_sys_user WHERE id = ?", id).Scan(&msSysUser.Id, &msSysUser.GmtCreate, &msSysUser.Psw, &msSysUser.GmtModified, &msSysUser.Username, &msSysUser.Key, &msSysUser.Email)
	if err != nil{
		fmt.Println("查询出错了")
	}
	return
}
func QueryUserByUsername(username string)(msSysUser MsSysUser,err error){
	handle := database.GetHandle()
	db,ok := handle.InitDB()
	if !ok{
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	err = db.QueryRow("SELECT `id`,`Gmt_Create`,`Psw`,`Gmt_Modified`,`Username`,`Key`, `Email` FROM ms_sys_user WHERE username = ?", username).Scan(&msSysUser.Id, &msSysUser.GmtCreate, &msSysUser.Psw, &msSysUser.GmtModified, &msSysUser.Username, &msSysUser.Key, &msSysUser.Email)
	if err != nil{
		fmt.Println("查询出错了")
	}
	return
}

func UpdateKeyByUserId(id int64,key string)(err error) {
	handle := database.GetHandle()
	_,ok := handle.InitDB()
	if !ok{
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	ok = handle.Update("UPDATE ms_sys_user SET `key` = ? WHERE `id` = ?",key,id)
	if !ok{
		err = errors.New("更新失败失败")
		return
	}
	return
}

func QueryUserByKey(key string)(msSysUser MsSysUser,err error)  {
	handle := database.GetHandle()
	db,ok := handle.InitDB()
	if !ok{
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	err = db.QueryRow("SELECT `id`,`Gmt_Create`,`Psw`,`Gmt_Modified`,`Username`,`Key`,`Email` FROM ms_sys_user WHERE `key` = ?", key).Scan(&msSysUser.Id, &msSysUser.GmtCreate, &msSysUser.Psw, &msSysUser.GmtModified, &msSysUser.Username, &msSysUser.Key, &msSysUser.Email)
	if err != nil{
		log.Println(err)
		fmt.Println("查询出错了")
	}
	return
}

func InsertUser(username,password,email string)(err error) {
	handle := database.GetHandle()
	_,ok := handle.InitDB()
	if !ok{
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	//todo 新增一列为是否进行邮箱校验
	if handle.Insert("INSERT INTO ms_sys_user(`GMT_CREATE`, `GMT_MODIFIED`, `USERNAME`, `PSW`, `EMAIL`)VALUES(now(),now(),?,?,?)",username,password,email){
		return nil
	}else{
		return errors.New("插入出错")
	}
}
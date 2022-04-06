package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/magic_space/common/database"
	"log"
)

type MsSysUser struct {
	Id          int64          `json:"id"`
	GmtCreate   string         `json:"gmt_create"`
	Psw         string         `json:"psw"`
	GmtModified string         `json:"gmt_modified"`
	Username    string         `json:"username"`
	Key         sql.NullString `json:"key"`
	Email       sql.NullString `json:"email"`
	EmailCnf    bool           `json:"email_cnf"`
}

func QueryUserById(id int64) (msSysUser MsSysUser, err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	err = db.QueryRow("SELECT `id`,`Gmt_Create`,`Psw`,`Gmt_Modified`,`Username`,`Key`,`Email`, `email_cnf` FROM ms_sys_user WHERE id = ?", id).Scan(&msSysUser.Id, &msSysUser.GmtCreate, &msSysUser.Psw, &msSysUser.GmtModified, &msSysUser.Username, &msSysUser.Key, &msSysUser.Email, &msSysUser.EmailCnf)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return
}

// 查询验证用户
func QueryUserByUsername(username string) (msSysUser MsSysUser, err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	err = db.QueryRow("SELECT `id`,`Gmt_Create`,`Psw`,`Gmt_Modified`,`Username`,`Key`, `Email`, `email_cnf` FROM ms_sys_user WHERE username = ? or email = ? and email_cnf = 1", username, username).Scan(&msSysUser.Id, &msSysUser.GmtCreate, &msSysUser.Psw, &msSysUser.GmtModified, &msSysUser.Username, &msSysUser.Key, &msSysUser.Email, &msSysUser.EmailCnf)
	return
}
func QueryUserByEmail(email string) (msSysUser MsSysUser, err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	err = db.QueryRow("SELECT `id`,`Gmt_Create`,`Psw`,`Gmt_Modified`,`Username`,`Key`, `Email`, `email_cnf` FROM ms_sys_user WHERE Email = ?", email).Scan(&msSysUser.Id, &msSysUser.GmtCreate, &msSysUser.Psw, &msSysUser.GmtModified, &msSysUser.Username, &msSysUser.Key, &msSysUser.Email, &msSysUser.EmailCnf)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return
}

func UpdateKeyByUserId(id int64, key string) (err error) {
	handle := database.GetHandle()
	_, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	ok = handle.Update("UPDATE ms_sys_user SET `key` = ? WHERE `id` = ?", key, id)
	if !ok {
		err = errors.New("更新失败失败")
		return
	}
	return
}

func QueryUserByKey(key string) (msSysUser MsSysUser, err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	err = db.QueryRow("SELECT `id`,`Gmt_Create`,`Psw`,`Gmt_Modified`,`Username`,`Key`,`Email` FROM ms_sys_user WHERE `key` = ?", key).Scan(&msSysUser.Id, &msSysUser.GmtCreate, &msSysUser.Psw, &msSysUser.GmtModified, &msSysUser.Username, &msSysUser.Key, &msSysUser.Email)
	if err != nil {
		log.Println(err)
		fmt.Println("查询出错了")
	}
	return
}

func InsertUser(username, password, email string) (err error) {
	handle := database.GetHandle()
	_, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	if handle.Insert("INSERT INTO ms_sys_user(`GMT_CREATE`, `GMT_MODIFIED`, `USERNAME`, `PSW`, `EMAIL`)VALUES(now(),now(),?,?,?)", username, password, email) {
		return nil
	} else {
		return errors.New("插入出错")
	}
}

func QueryUserConfirmByEmail(email string) (rowsNum int, confirmByEmail bool, err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		return 0, false, errors.New("数据库连接失败")
	}
	defer handle.CloseDB()
	sqlColm := `SELECT EMAIL_CNF FROM ms_sys_user WHERE EMAIL = ?`
	rows, err := db.Query(sqlColm, email)
	if err != nil {
		return 0, false, errors.New("查询失败")
	}
	for rows.Next() { //能够取值就读取下一行
		rowsNum++
		err = rows.Scan(&confirmByEmail)
		if err != nil {
			fmt.Println("[ConfirmByEmail]出错了 ", err)
			return 0, false, errors.New("[ConfirmByEmail]出错了")
		}
	}
	return rowsNum, confirmByEmail, nil
}

func UpdateUserConfirmByEmail(email string) (err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		return errors.New("数据库连接失败")
	}
	defer handle.CloseDB()
	sqlStr := `UPDATE ms_sys_user SET EMAIL_CNF=1 WHERE EMAIL = ?`
	_, err = db.Exec(sqlStr, email)
	if !ok {
		return errors.New("更新失败失败")
	}
	return nil
}

func QueryUserConfirmByUser(user string) (rowsNum int, confirmByEmail bool, err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		return 0, false, errors.New("数据库连接失败")
	}
	defer handle.CloseDB()
	sqlColm := `SELECT EMAIL_CNF FROM ms_sys_user WHERE USERNAME = ?`
	rows, err := db.Query(sqlColm, user)
	if err != nil {
		return 0, false, errors.New("查询失败")
	}
	for rows.Next() { //能够取值就读取下一行
		rowsNum++
		err = rows.Scan(&confirmByEmail)
		if err != nil {
			fmt.Println("[ConfirmByEmail]出错了 ", err)
			return 0, false, errors.New("[ConfirmByEmail]出错了")
		}
	}
	return rowsNum, confirmByEmail, nil
}

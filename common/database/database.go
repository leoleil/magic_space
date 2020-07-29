package database

import (
	"database/sql"
	"github.com/leoleil/magic_space/common/database/mysql"
)

type DataBase interface {
	InitDB()(*sql.DB,bool)
	Insert(sqlStr string,arg ... interface{})bool
	Delete(sqlStr string,arg ... interface{})bool
	Update(sqlStr string,arg ... interface{})bool
	CloseDB()
}

func GetHandle()(handle DataBase)  {
	handle = &mysql.Handle{}
	return handle
}
package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/leoleil/magic_space/common/config"

	_ "github.com/go-sql-driver/mysql"
)
type Handle struct {
	db *sql.DB
}

func (h *Handle)InitDB()(DB *sql.DB, ok bool) {
	path := strings.Join([]string{config.AppHandle.Mysql.User, ":", config.AppHandle.Mysql.Pwd, "@tcp(",config.AppHandle.Mysql.Host, ":", config.AppHandle.Mysql.Port, ")/", config.AppHandle.Mysql.Dbname, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		return nil,false
	}
	h.db = DB
	ok = true
	return
}

func (h *Handle)Insert(sqlStr string,arg ... interface{})bool{
	//开启事务
	tx, err := h.db.Begin()
	if err != nil{
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare(sqlStr)
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(arg ...)
	if err != nil{
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func (h *Handle)Delete(sqlStr string,arg ... interface{})bool{
	//开启事务
	tx, err := h.db.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare(sqlStr)
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(arg ... )
	if err != nil{
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}

func (h *Handle)Update(sqlStr string,arg ... interface{})bool{
	//开启事务
	tx, err := h.db.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare(sqlStr)
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(arg ... )
	if err != nil{
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
func (h *Handle)CloseDB(){
	h.db.Close()
}



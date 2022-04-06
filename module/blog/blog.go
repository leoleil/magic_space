package blog

import (
	"errors"
	"fmt"
	"github.com/magic_space/common/database"
)

type BlogEntity struct {
	Id          int64  `json:"id"`
	GmtCreate   string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	Nice        int64  `json:"nice"`
	PV          int64  `json:"pv"`
	UserName    string `json:"user_name"`
	UserId      int64  `json:"user_id"`
}

func InsertBlog(b BlogEntity) (err error) {
	handle := database.GetHandle()
	_, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	if handle.Insert("INSERT INTO ms_service_blog(`GMT_CREATE`, `GMT_MODIFIED`, `TITLE`, `USER_ID`, `USERNAME`, `BODY`, `NICE`, `PV`)VALUES(now(),now(),?,?,?,?,?,?)", b.Title, b.UserId, b.UserName, b.Body, b.Nice, b.PV) {
		return err
	} else {
		return errors.New("blog 插入数据库出错")
	}
}
func SelectBlogById(id int64) (b BlogEntity, err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return b, err
	}
	defer handle.CloseDB()
	err = db.QueryRow("SELECT `ID`, `GMT_CREATE`, `GMT_MODIFIED`, `TITLE`, `USER_ID`, `USERNAME`, `BODY`, `NICE`, `PV` FROM ms_service_blog WHERE id = ?", id).Scan(&b.Id, &b.GmtCreate, &b.GmtModified, &b.Title, &b.UserId, &b.UserName, &b.Body, &b.Nice, &b.PV)
	if err != nil {
		fmt.Println("查询出错了")
	}
	return b, err
}
func SelectBlogListByPage(page, limit int64) (b []BlogEntity, err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return b, err
	}
	defer handle.CloseDB()
	rows, err := db.Query("SELECT `ID`, `GMT_CREATE`, `GMT_MODIFIED`, `TITLE`, `USER_ID`, `USERNAME`, `BODY`, `NICE`, `PV` FROM ms_service_blog ORDER BY GMT_MODIFIED DESC LIMIT ? OFFSET ?", limit, (page-1)*limit)
	if err != nil {
		fmt.Println("查询出错了")
		return b, err
	}
	for rows.Next() {
		blog := BlogEntity{}
		err = rows.Scan(&blog.Id, &blog.GmtCreate, &blog.GmtModified, &blog.Title, &blog.UserId, &blog.UserName, &blog.Body, &blog.Nice, &blog.PV)
		if err != nil {
			continue
		}
		b = append(b, blog)
	}
	return b, err
}
func CountBlog() (n int64, err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return n, err
	}
	defer handle.CloseDB()
	err = db.QueryRow("SELECT COUNT(`ID`) FROM ms_service_blog").Scan(&n)
	return n, err
}

func UpdateBlogById(b BlogEntity) (err error) {
	handle := database.GetHandle()
	_, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	if !handle.Update("UPDATE ms_service_blog SET `TITLE` = ? , `BODY` = ? , `GMT_MODIFIED` = now() WHERE `id` = ? ", b.Title, b.Body, b.Id) {
		err = errors.New("更新blog 出错")
	}
	return err
}

func DeleteBlogById(id int64) (err error) {
	handle := database.GetHandle()
	_, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	if !handle.Delete("DELETE FROM ms_service_blog WHERE ID = ?", id) {
		err = errors.New("删除失败")
	}
	return err
}
func DeleteBlogByIdAndUserId(id, userId int64) (err error) {
	handle := database.GetHandle()
	_, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	if !handle.Delete("DELETE FROM ms_service_blog WHERE ID = ? and USER_ID = ?", id, userId) {
		err = errors.New("删除失败")
	}
	return err
}

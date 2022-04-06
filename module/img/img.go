package img

import (
	"errors"
	"fmt"
	"github.com/magic_space/common/database"
)

type ImgEntity struct {
	Id          int64  `json:"id"`
	GmtCreate   string `json:"gmt_create"`
	GmtModified string `json:"gmt_modified"`
	BlogId      int64  `json:"blog_id"`
	Path        string `json:"path"`
	Url         string `json:"url"`
}

func InsertImg(blogId int64, path, url string) (err error) {
	handle := database.GetHandle()
	_, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	if handle.Insert("INSERT INTO ms_service_img(`GMT_CREATE`, `GMT_MODIFIED`, `BLOG_ID`, `PATH`, `URL`)VALUES(now(),now(),?,?,?)", blogId, path, url) {
		return err
	} else {
		return errors.New("img 插入数据库出错")
	}
}

func SelectImgByByUrl(url string) (i ImgEntity, err error) {
	handle := database.GetHandle()
	db, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return i, err
	}
	defer handle.CloseDB()
	err = db.QueryRow("SELECT `ID`, `GMT_CREATE`, `GMT_MODIFIED`, `BLOG_ID`, `PATH`, `URL` FROM ms_service_img WHERE URL = ?", url).Scan(&i.Id, &i.GmtCreate, &i.GmtModified, &i.BlogId, &i.Path, &i.Url)
	if err != nil {
		fmt.Println("img 查询出错了")
	}
	return i, err
}

func DeleteImgByBLogId(blogId int64) (err error) {
	handle := database.GetHandle()
	_, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	if !handle.Delete("DELETE FROM ms_service_img WHERE BLOG_ID = ? ", blogId) {
		err = errors.New("删除失败")
	}
	return err
}
func DeleteImgByPath(path string) (err error) {
	handle := database.GetHandle()
	_, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	if !handle.Delete("DELETE FROM ms_service_img WHERE 	PATH = ? ", path) {
		err = errors.New("删除失败")
	}
	return err
}
func DeleteImgByUrl(url string) (err error) {
	handle := database.GetHandle()
	_, ok := handle.InitDB()
	if !ok {
		err = errors.New("数据库连接失败")
		return
	}
	defer handle.CloseDB()
	if !handle.Delete("DELETE FROM ms_service_img WHERE 	URL = ? ", url) {
		err = errors.New("删除失败")
	}
	return err
}

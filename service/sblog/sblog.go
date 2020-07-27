package sblog

import (
	"github.com/leoleil/magic_space/module/blog"
	"github.com/leoleil/magic_space/module/user"
)

func UploadBlog(title, body, key string) (err error) {
	u, err := user.QueryUserByKey(key)
	if err != nil {
		return err
	}
	entity := blog.BlogEntity{
		Title:    title,
		Body:     body,
		Nice:     0,
		PV:       0,
		UserName: u.Username,
		UserId:   u.Id,
	}
	err = blog.InsertBlog(entity)
	return err
}

func GetBlogListByPage(page, limit int64) (b []blog.BlogEntity, pageNum, endPageNum int64, err error) {
	n, err := blog.CountBlog()
	if err != nil {
		return
	}
	if n%limit != 0 {
		endPageNum = n/limit + 1
	} else {
		endPageNum = n / limit
	}
	b, err = blog.SelectBlogListByPage(page, limit)
	if err != nil {
		return
	}
	pageNum = page
	return
}

func UpdateBlog(id int64, title, body, key string) (err error) {
	u, err := user.QueryUserByKey(key)
	if err != nil {
		return err
	}
	entity := blog.BlogEntity{
		Id:       id,
		Title:    title,
		Body:     body,
		Nice:     0,
		PV:       0,
		UserName: u.Username,
		UserId:   u.Id,
	}
	err = blog.UpdateBlogById(entity)
	return err
}

func DeleteBlog(id int64, key string) (err error) {
	u, err := user.QueryUserByKey(key)
	if err != nil {
		return err
	}
	err = blog.DeleteBlogByIdAndUserId(id, u.Id)
	return err
}

func OpenBlog(id int64, key string) (b blog.BlogEntity, edit bool, err error) {
	b, err = blog.SelectBlogById(id)
	if key == "" {
		edit = false
		return
	}
	u, err := user.QueryUserByKey(key)
	if err != nil {
		return
	}

	if b.UserId == u.Id {
		edit = true
	}
	return
}

package simg

import (
	"io"
	"magic_space/module/img"
	"magic_space/module/user"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

func SaveImg(file multipart.File, fileName, key string) (url string, err error) {
	u, err := user.QueryUserByKey(key)
	if err != nil {
		return
	}
	uerId := strconv.FormatInt(u.Id, 10)
	t := time.Now().Unix()
	path := "view/image/" + uerId + "/" + strconv.FormatInt(t, 10) + "/"
	//创建目录
	err = os.MkdirAll(path, 0777)
	if err != nil {
		return
	}
	path = path + fileName
	out, err := os.Create(path)
	if err != nil {
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		return
	}
	url = "assets/image/" + uerId + "/" + strconv.FormatInt(t, 10) + "/" + fileName
	err = img.InsertImg(-1, path, url)
	if err != nil {
		os.RemoveAll("view/image/" + uerId + "/" + strconv.FormatInt(t, 10))
		return
	}
	return
}
func DeleteImg(url string) (err error) {
	i, err := img.SelectImgByByUrl(url)
	if err != nil {
		return
	}
	paths := strings.Split(i.Path, "/")
	path := strings.Join(paths[:len(paths)-1], "/")
	os.RemoveAll(path)
	err = img.DeleteImgByUrl(url)
	return
}

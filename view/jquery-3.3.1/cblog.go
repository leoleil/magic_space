package jquery_3_3_1

import (
	"github.com/gin-gonic/gin"
	"github.com/leoleil/magic_space/common/utilities"
	"github.com/leoleil/magic_space/service/sblog"
	"github.com/leoleil/magic_space/service/simg"
	"net/http"
	"strconv"
)

func UploadBlog(context *gin.Context) {
	title := context.DefaultPostForm("title", "")
	body := context.DefaultPostForm("body", "")
	key := utilities.GetKey(context)
	err := sblog.UploadBlog(title, body, key)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "上传失败",
		})
	} else {
		context.HTML(http.StatusOK, "blog.html", gin.H{
			"title": "MC Space",
			"page":  1,
		})
	}
	return
}

func GetBlogListByPage(context *gin.Context) {
	page := context.DefaultQuery("page", "1")
	limit := context.DefaultQuery("limit", "5")
	pageInt, _ := strconv.ParseInt(page, 10, 64)
	limitInt, _ := strconv.ParseInt(limit, 10, 64)
	data, pageNum, endPageNum, err := sblog.GetBlogListByPage(pageInt, limitInt)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "列表获取失败",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"msg":     "列表获取成功",
			"list":    data,
			"page":    pageNum,
			"endPage": endPageNum,
			"limit":   limitInt,
		})
	}
	return
}

func UpdateBlog(context *gin.Context) {
	id := context.DefaultPostForm("id", "")
	title := context.DefaultPostForm("title", "")
	body := context.DefaultPostForm("body", "")
	key := utilities.GetKey(context)
	idInt, _ := strconv.ParseInt(id, 10, 64)
	err := sblog.UpdateBlog(idInt, title, body, key)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "更新失败",
		})
	} else {
		context.HTML(http.StatusOK, "blog.html", gin.H{
			"title": "MC Space",
			"page":  1,
		})
	}
	return
}

func DeleteBlog(context *gin.Context) {
	id := context.DefaultPostForm("id", "")
	key := utilities.GetKey(context)
	idInt, _ := strconv.ParseInt(id, 10, 64)
	err := sblog.DeleteBlog(idInt, key)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "删除失败",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"msg": "删除成功",
		})
	}
	return
}

func OpenBlog(context *gin.Context) {
	id := context.DefaultQuery("id", "")
	key := utilities.GetKey(context)
	idInt, _ := strconv.ParseInt(id, 10, 64)
	entity, edit, err := sblog.OpenBlog(idInt, key)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "打开失败",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"msg":    "打开成功",
			"edit":   edit,
			"entity": entity,
		})
	}
	return
}

func LoadImg(context *gin.Context) {
	key := utilities.GetKey(context)
	file, header, err := context.Request.FormFile("file")
	//id := context.DefaultPostForm("id", "")
	//idInt, _ := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "上传失败",
		})
		return
	}
	filename := header.Filename
	url, err := simg.SaveImg(file, filename, key)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "上传失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":      "上传成功",
		"location": url,
	})
	return
}
func DeleteImg(context *gin.Context) {
	url := context.DefaultPostForm("url", "")
	err := simg.DeleteImg(url)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "删除失败",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
		"url": url,
	})
	return
}

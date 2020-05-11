package index

import (
	"github.com/gin-gonic/gin"
	"magic_space/service/sblog"
	"strconv"
)

func LoadIndex(context *gin.Context)  {
	context.HTML(200,"index.html",gin.H{
		"title": "MC Space",
	})
	return
}
func LoadBlog(context *gin.Context)  {
	page := context.DefaultQuery("page","1")
	limit := context.DefaultQuery("limit", "5")
	pageInt, _ := strconv.ParseInt(page, 10, 64)
	limitInt, _ := strconv.ParseInt(limit, 10, 64)
	data, pageNum, endPageNum, _ := sblog.GetBlogListByPage(pageInt,limitInt)
	context.SetCookie("blog_list_page",page,0,"/","localhost",false,false)
	context.HTML(200,"blog.html",gin.H{
		"title": "MC Space",
		"list": data,
		"page": pageNum,
		"endPage": endPageNum,
		"limit": limitInt,
	})
	return
}
func LoadBlogView(context *gin.Context)  {
	id := context.DefaultQuery("id", "")
	context.HTML(200,"blog-view.html",gin.H{
		"title": "MC Space",
		"id": id,
	})
	return
}
func LoadBlogDoc(context *gin.Context)  {
	context.HTML(200,"blog-doc.html",gin.H{
		"title": "MC Space",
	})
	return
}
func LoadBlogEdit(context *gin.Context)  {
	id := context.DefaultQuery("id", "")
	context.HTML(200,"blog-edit.html",gin.H{
		"title": "MC Space",
		"id": id,
	})
	return
}
func LoadSignIn(context *gin.Context)  {
	context.HTML(200,"sign-in.html",gin.H{
		"title": "MC Space",
	})
	return
}
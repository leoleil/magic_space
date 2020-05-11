package main

import (
	"github.com/gin-gonic/gin"
	"magic_space/common/config"
	"magic_space/controller/asd"
	"magic_space/controller/cblog"
	"magic_space/controller/index"
)

func init() {
	config.AppHandle.GetConf()
}
func main()  {
	router := gin.Default()
	router.Static("/assets", "./view")
	router.StaticFile("/favicon.ico", "./view/image/favicon.ico")
	router.LoadHTMLGlob("templates/*")
	router.GET("/index",index.LoadIndex)
	router.GET("/join",index.LoadSignIn)
	router.GET("/blog",index.LoadBlog)
	router.GET("/blog/view",index.LoadBlogView)
	router.GET("/blog/doc",index.LoadBlogDoc)
	router.GET("/blog/edit",index.LoadBlogEdit)
	user := router.Group("/asd")
	{
		user.POST("/login", asd.Login)
		user.POST("/sign", asd.SignIn)
		user.POST("/check", asd.Check)
	}
	blog := router.Group("/blog")
	{
		blog.GET("/list", cblog.GetBlogListByPage)
		blog.POST("/upload", cblog.UploadBlog)
		blog.POST("/update", cblog.UpdateBlog)
		blog.POST("/delete", cblog.DeleteBlog)
		blog.GET("/open", cblog.OpenBlog)
	}
	router.Run(":4010")
}

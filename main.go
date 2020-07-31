package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leoleil/magic_space/common/config"
	"github.com/leoleil/magic_space/controller/asd"
	"github.com/leoleil/magic_space/controller/cblog"
	"github.com/leoleil/magic_space/controller/cvideo"
	"github.com/leoleil/magic_space/controller/index"
)

func init() {
	config.AppHandle.GetConf()
}
func main() {
	router := gin.Default()
	router.Static("/assets", "./view")
	router.StaticFile("/favicon.ico", "./view/image/favicon.ico")
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", index.LoadIndex)
	router.GET("/join", index.LoadSignIn)
	router.GET("/blog", index.LoadBlog)
	router.GET("/blog/view", index.LoadBlogView)
	router.GET("/blog/doc", index.LoadBlogDoc)
	router.GET("/blog/edit", index.LoadBlogEdit)
	router.GET("/video", index.LoadVideo)
	router.GET("/video/open", index.LoadVideoOpen)
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
		blog.POST("/img/upload", cblog.LoadImg)
		blog.POST("/img/delete", cblog.DeleteImg)
	}
	video := router.Group("/video")
	{
		video.GET("/list", cvideo.GetVideoList)
	}
	router.Run(":4010")
}

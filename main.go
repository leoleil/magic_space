package main

import (
	"github.com/gin-gonic/gin"
	"github.com/leoleil/magic_space/common/config"
	"github.com/leoleil/magic_space/controller/asd"
	"github.com/leoleil/magic_space/controller/cvideo"
	"github.com/leoleil/magic_space/controller/index"
	"github.com/leoleil/magic_space/view/jquery-3.3.1"
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
		blog.GET("/list", jquery_3_3_1.GetBlogListByPage)
		blog.POST("/upload", jquery_3_3_1.UploadBlog)
		blog.POST("/update", jquery_3_3_1.UpdateBlog)
		blog.POST("/delete", jquery_3_3_1.DeleteBlog)
		blog.GET("/open", jquery_3_3_1.OpenBlog)
		blog.POST("/img/upload", jquery_3_3_1.LoadImg)
		blog.POST("/img/delete", jquery_3_3_1.DeleteImg)
	}
	video := router.Group("/video")
	{
		video.GET("/list", cvideo.GetVideoList)
	}
	router.Run(":4010")
}

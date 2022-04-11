package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/magic_space/common/config"
	"github.com/magic_space/controller/casd"
	"github.com/magic_space/controller/cblog"
	"github.com/magic_space/controller/cvideo"
	"github.com/magic_space/controller/index"
)

var path = flag.String("config", "app.yml", "-config app.yml")

func main() {
	flag.Parse()
	config.AppHandle.GetConf("config/" + *path)
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
		user.POST("/login", casd.Login)
		user.POST("/sign", casd.SignIn)
		user.POST("/check", casd.Check)
		user.GET("/sign/confirm", casd.ConfirmEmail)
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
	err := router.Run(":" + config.AppHandle.Host.Port)
	if err != nil {
		return
	}
}

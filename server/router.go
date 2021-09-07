package server

import (
	"github.com/bh-qt/alist/conf"
	"github.com/bh-qt/alist/server/controllers"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// InitRouter init router
func InitRouter(engine *gin.Engine) {
	log.Infof("初始化路由...")
	engine.Use(CorsHandler())
	engine.Use(static.Serve("/", static.LocalFile(conf.Conf.Server.Static, false)))
	engine.NoRoute(func(c *gin.Context) {
		c.File(conf.Conf.Server.Static + "/index.html")
	})
	InitApiRouter(engine)
}

// InitApiRouter init api router
func InitApiRouter(engine *gin.Engine) {
	apiV2 := engine.Group("/api")
	{
		apiV2.GET("/info", controllers.Info)
		apiV2.POST("/get", controllers.Get)
		apiV2.POST("/path", controllers.Path)
		apiV2.POST("/local_search", controllers.LocalSearch)
		apiV2.POST("/global_search", controllers.GlobalSearch)
		apiV2.POST("/rebuild", controllers.RebuildTree)

		apiV2.POST("/office_preview/:drive", controllers.OfficePreview)
		apiV2.POST("/video_preview/:drive", controllers.VideoPreview)
		apiV2.POST("/video_preview_play_info/:drive", controllers.VideoPreviewPlayInfo)
		engine.GET("/d/*path", controllers.Down)
		engine.GET("/p/*path",controllers.Proxy)
	}
}

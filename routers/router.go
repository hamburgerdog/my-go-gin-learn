package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"xjosiah.com/go-gin/pkg/export"
	"xjosiah.com/go-gin/pkg/setting"
	"xjosiah.com/go-gin/pkg/upload"
	"xjosiah.com/go-gin/routers/api"
	v1 "xjosiah.com/go-gin/routers/api/v1"

	_ "xjosiah.com/go-gin/docs"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)

	apiV1 := r.Group("/api/v1")
	//	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)

		apiV1.PUT("/tags/:id", v1.EditTag)

		apiV1.POST("/tags", v1.AddTag)

		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		apiV1.GET("/articles", v1.GetArticles)

		apiV1.GET("/articles/:id", v1.GetArticle)

		apiV1.POST("/articles", v1.AddArticle)

		apiV1.PUT("/articles/:id", v1.EditArticle)

		apiV1.DELETE("/articles/:id", v1.DeleteArticle)

		r.POST("/tags/export", v1.ExportTag)
		r.POST("/tags/import", v1.ImportTag)
	}
	return r
}

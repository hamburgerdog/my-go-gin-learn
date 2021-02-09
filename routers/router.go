package routers

import (
	"github.com/gin-gonic/gin"

	"xjosiah.com/go-gin/middleware/jwt"
	"xjosiah.com/go-gin/pkg/setting"
	"xjosiah.com/go-gin/routers/api"
	v1 "xjosiah.com/go-gin/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
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
	}
	return r
}

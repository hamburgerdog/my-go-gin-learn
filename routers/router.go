package routers

import (
	"github.com/gin-gonic/gin"

	"xjosiah.com/go-gin/pkg/setting"
	v1 "xjosiah.com/go-gin/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags", v1.GetTags)

		apiV1.PUT("/tags/:id", v1.EditTag)

		apiV1.POST("/tags", v1.AddTag)

		apiV1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return r
}

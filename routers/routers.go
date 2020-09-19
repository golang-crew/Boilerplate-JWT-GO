package routers

import (
	"github.com/gin-gonic/gin"
	router "github.com/golang-crew/Bolierplate-JWT-Auth/routers/v1"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/v1")
	{
		{
			router.ApplyRoutes(v1)
		}
	}

}

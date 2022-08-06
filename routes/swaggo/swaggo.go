package swaggo

import (
	"github.com/gin-gonic/gin"

	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/go_rest_api_playground/docs"
)

func Routes(route *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/api"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

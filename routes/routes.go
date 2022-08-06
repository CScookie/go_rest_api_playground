package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go_rest_api_playground/routes/book"
	"github.com/go_rest_api_playground/routes/cache"
	"github.com/go_rest_api_playground/routes/swaggo"
)

func Routes(route *gin.Engine) {
	api := route.Group("api")
	swaggo.Routes(route)
	book.Routes(api)
	cache.Routes(api)
	fmt.Println(api)

}

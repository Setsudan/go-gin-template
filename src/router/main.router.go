package router

import (
	"gallgo/app/src/router/routes"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// serve all files in public folder
	routes.ServePublicAssets(router)

	// serve all files in src/pages folder
	routes.ServeHtml(router)

	// ping route
	routes.PingRoute(router)

	// auto redirect "/" to "/page"
	router.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/page")
	})

	return router
}

func Run() {
	router := NewRouter()

	router.Run()
}

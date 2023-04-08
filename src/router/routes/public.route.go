package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServePublicAssets(router *gin.Engine) {
	// serve all files in public folder
	router.StaticFS("/public", http.Dir("public"))
}

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeHtml(router *gin.Engine) {
	router.StaticFS("/page", http.Dir("src/pages"))
}

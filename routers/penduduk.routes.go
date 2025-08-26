package routers

import (
	"github.com/EkyGalih/dukcapil-web/controllers"
	"github.com/gin-gonic/gin"
)

func PendudukRoutes(r *gin.Engine) {
	r.GET("/penduduk", controllers.PendudukIndex)
}
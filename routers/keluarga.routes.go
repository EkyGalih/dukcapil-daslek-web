package routers

import (
	"github.com/EkyGalih/dukcapil-web/controllers"
	"github.com/gin-gonic/gin"
)

func KeluargaRoutes(r *gin.Engine) {
	r.GET("/keluarga", controllers.KeluargaIndex)
}
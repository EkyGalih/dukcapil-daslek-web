package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/EkyGalih/dukcapil-web/controllers"
)

func RouterList(r *gin.Engine) {
    // Homepage
    r.GET("/dashboard", controllers.Dashboard)

    // Penduduk
    PendudukRoutes(r)

    // Keluarga
    KeluargaRoutes(r)
}

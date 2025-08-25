package routers

import (
	"github.com/gin-gonic/gin"
	"github.gom/EkyGalih/dukcapil-web/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

	// Templates & Static
	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/static", "./static")

    // Homepage
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "layouts/app", gin.H{"title": "Dashboard Pendataan"})
    })

    // Routes 7 tabel
    r.GET("/keluarga", controllers.KeluargaIndex)

    return r
}

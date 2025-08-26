package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/EkyGalih/dukcapil-web/helpers"
)

func Dashboard(c *gin.Context) {
	data := map[string]any{
		"title": "Dashboard",
	}
	helpers.RenderTemplate(c, "pages/dashboard/dashboard.html", data)
}

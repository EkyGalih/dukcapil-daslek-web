package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/EkyGalih/dukcapil-web/routers"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	routers.RouterList(r)
	for _, route := range r.Routes() {
		log.Printf("%-6s %s", route.Method, route.Path)
	}
	log.Println("server started at http://127.0.0.1:8080")
	r.Run(":8080")
}

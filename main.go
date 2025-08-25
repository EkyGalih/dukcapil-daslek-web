package main

import "github.gom/EkyGalih/dukcapil-web/routers"

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}

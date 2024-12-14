package main

import (
	"SimpleGolang/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Main Runner
	rute := gin.Default()
	router.MainRouter(rute)
	rute.Run()
}

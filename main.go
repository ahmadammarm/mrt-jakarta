package main

import (
	"github.com/ahmadammarm/mrt-jakarta/station"
	"github.com/gin-gonic/gin"
)

func InitiateRouter() {
	var (
		router = gin.Default()
		api    = router.Group("/api")
	)

	station.Initiate(api)

    router.Run(":8080")

}

func main() {
	InitiateRouter()
}

package main

import (
	"github.com/ahmadammarm/mrt-jakarta/internal/routers"
	"github.com/gin-gonic/gin"
)

func InitiateRouter() {
    router := gin.Default()
    api := router.Group("/v1/api")


    routers.Routers(api)
    router.Run(":8080")

}

func main() {

}

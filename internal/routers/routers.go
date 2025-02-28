package routers

import "github.com/gin-gonic/gin"

func Routers(router *gin.RouterGroup) {
    station := router.Group("/stations")
    station.GET("/", func (*gin.Context) {
        
    })
}
package station

import "github.com/gin-gonic/gin"

func Initiate(router *gin.RouterGroup) {

    stationService := NewService()

    station := router.Group("/station")
    station.GET("/", func(context *gin.Context) {
        GetAllStations(context, stationService)
    })
}

func GetAllStations(context *gin.Context, service StationService) {
    datas, err := service.GetAllStations()
    if err != nil {
        context.JSON(500, gin.H{"message": err.Error()})
        return
    }

    context.JSON(200, datas)
}

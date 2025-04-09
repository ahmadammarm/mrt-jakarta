package station

import (
	"net/http"

	"github.com/ahmadammarm/mrt-jakarta/common/response"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {

	stationService := NewService()

	station := router.Group("/stations")
	station.GET("/", func(context *gin.Context) {
		GetAllStations(context, stationService)
	})
}

func GetAllStations(context *gin.Context, service StationService) {
	datas, err := service.GetAllStations()
	if err != nil {
		context.JSON(http.StatusBadRequest, response.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
			Success: false,
		})
		return
	}

	context.JSON(http.StatusOK, response.APIResponse{
        Code:    http.StatusOK,
        Message: "Success get all stations",
        Data:    datas,
        Success: true,
    })
}

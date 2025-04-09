package station

import (
	"net/http"
	"time"
)

type StationService interface {
    GetAllStations() (response []StationResponse, err error)
}

type stationService struct {
    client *http.Client
}


func (service *stationService) GetAllStations() (response []StationResponse, err error) {

    url := "https://www.jakartamrt.co.id/id/val/stasiuns"

    // hit url

    // keluarkan response ke struct StationResponse

    return
}

func NewService() StationService {
    return &stationService{
        client: &http.Client{
            Timeout: 5 * time.Second,
        },
    }
}
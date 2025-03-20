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

func NewService() StationService {
    return &service{
        client: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

func (service *stationService) GetAllStations() (response []StationResponse, err error) {
    
}
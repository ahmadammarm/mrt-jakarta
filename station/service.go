package station

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ahmadammarm/mrt-jakarta/common/client"
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
    byteResponse, err := client.Request(service.client, url)

    if err != nil {
        return
    }

    var stations []Station

    err = json.Unmarshal(byteResponse, &stations)

    for _, station := range stations {
        response = append(response, StationResponse{
            Id:   station.ID,
            Name: station.Name,
        })
    }
    
    return
}

func NewService() StationService {
    return &stationService{
        client: &http.Client{
            Timeout: 5 * time.Second,
        },
    }
}
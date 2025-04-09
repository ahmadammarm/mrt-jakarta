package station

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/ahmadammarm/mrt-jakarta/common/client"
)

type StationService interface {
	GetAllStations() (response []StationResponse, err error)
	ScheduleCheckByStationId(stationId string) (response []ScheduleResponse, err error)
}

type stationService struct {
	client *http.Client
}

func (service *stationService) GetAllStations() (response []StationResponse, err error) {

	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	// hit url
	byteResponse, err := client.ClientRequest(service.client, url)

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

func (service *stationService) ScheduleCheckByStationId(stationId string) (response []ScheduleResponse, err error) {
	url := "https://www.jakartamrt.co.id/id/val/stasiuns/"

	byteResponse, err := client.ClientRequest(service.client, url)

	if err != nil {
		return
	}

	var schedules []Schedule
	err = json.Unmarshal(byteResponse, &schedules)
	if err != nil {
		return
	}

	var scheduleSelected Schedule

	for _, schedule := range schedules {
		if schedule.StationID == stationId {
			scheduleSelected = schedule
			break
		}
	}

	if scheduleSelected.StationID == "" {
		err = errors.New("station not found")
		return
	}

	response, err = ConvertDataToResponses(scheduleSelected)

	if err != nil {
		return
	}

	return

}

func ConvertDataToResponses(schedule Schedule) (response []ScheduleResponse, err error) {
	var (
		LebakBulusTripName = "Lebak Bulus Station"
		BundaranHITripName = "Bundaran HI Station"
	)

	scheduleLebakBulus := schedule.ScheduleLebakBulus
	scheduleBundaranHI := schedule.ScheduleBundaranHI

    lebakBulusParsed, err := ConvertScheduleToTime(scheduleLebakBulus)
    if err != nil {
        return
    }

    bundaranHIParsed, err := ConvertScheduleToTime(scheduleBundaranHI)

    if err != nil {
        return
    }

    for _, item := range lebakBulusParsed {
        if item.Format("15.04") > time.Now().Format("15.04") {
            response = append(response, ScheduleResponse{
                StationName: LebakBulusTripName,
                Time:        item.Format("15:04"),
            })
        }
    }

	for _, item := range bundaranHIParsed {
        if item.Format("15.04") > time.Now().Format("15.04") {
            response = append(response, ScheduleResponse{
                StationName: BundaranHITripName,
                Time:        item.Format("15:04"),
            })
        }
    }




    return

}

func ConvertScheduleToTime(schedule string) (response []time.Time, err error) {
	var (
		parsedTime time.Time
		schedules  = strings.Split(schedule, ",")
	)

	for _, schedule := range schedules {
		trimedTime := strings.TrimSpace(schedule)
		if trimedTime == "" {
			continue
		}

		parsedTime, err = time.Parse("15:04", trimedTime)

		if err != nil {
			return
		}

		response = append(response, parsedTime)
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

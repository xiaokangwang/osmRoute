package attributes

import (
	"encoding/json"
	"github.com/xiaokangwang/osmRoute/interfacew"
	"time"
)

type BikeStation struct {
	InputData
}

func (b *BikeStation) ListAllStations() []interfacew.MapLocation {
	var ret []interfacew.MapLocation
	for _, v := range b.InputData {
		if len(v.Historic) > 0 {
			if v.Historic[0].AvailableBikes != 0 && v.Historic[0].AvailableBikeStands != 0 {
				ret = append(ret, interfacew.MapLocation{
					Lat: v.Latitude,
					Lon: v.Longitude,
				})
			}
		}
	}
	return ret
}

func NewBikeStationData(input string) (BikeStation, error) {
	var inputd InputData
	err := json.Unmarshal([]byte(input), &inputd)
	if err != nil {
		return BikeStation{}, newError("Cannot unwrap data").Base(err)
	}
	return BikeStation{inputd}, nil
}

type InputData []struct {
	Address   string     `json:"address"`
	Banking   string     `json:"banking"`
	Historic  []Historic `json:"historic"`
	ID        int        `json:"id"`
	Latitude  float64    `json:"latitude"`
	Longitude float64    `json:"longitude"`
	Name      string     `json:"name"`
}
type Historic struct {
	AvailableBikeStands int       `json:"available_bike_stands"`
	AvailableBikes      int       `json:"available_bikes"`
	BikeStands          int       `json:"bike_stands"`
	Status              string    `json:"status"`
	Time                time.Time `json:"time"`
}

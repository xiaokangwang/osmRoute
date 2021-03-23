package attributes

import (
	"bytes"
	"encoding/json"
)

type BusInfoItem struct {
	RemainingTime string
}

type BusInfo struct {
	Times map[string]BusInfo
}

func NewBusInfo(input string) (BusInfo, error) {
	var bus BusInfo
	decoder := json.NewDecoder(bytes.NewReader([]byte(input)))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&bus)

	if err != nil {
		return BusInfo{}, newError("Cannot unwrap input").Base(err)
	}
	return bus, nil
}

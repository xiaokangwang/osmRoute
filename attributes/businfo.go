package attributes

import (
	"bytes"
	"encoding/json"
)

type BusInfoItem struct {
	RemainingTime int64
}

type BusInfo struct {
	Info map[string]BusInfoItem
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

func (b BusInfo) WaitTime(NodeFID string) int64 {
	if data, ok := b.Info[NodeFID]; ok {
		return data.RemainingTime
	}
	return 600
}

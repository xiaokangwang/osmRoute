package attributes

import (
	"bytes"
	"encoding/json"
)

type AreaToAvoidItem struct {
	L, R, T, B float64
}
type AreaToAvoid struct {
	Items []AreaToAvoidItem
}

func NewAreaToAvoid(input string) (AreaToAvoid, error) {
	var areaToAvoid AreaToAvoid
	decoder := json.NewDecoder(bytes.NewReader([]byte(input)))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&areaToAvoid)
	if err != nil {
		return AreaToAvoid{}, newError("Cannot unwrap input").Base(err)
	}
	return areaToAvoid, nil
}

func (a *AreaToAvoid) CheckPointInclusion(lat, lon float64) bool {
	var included bool
	for _, v := range a.Items {
		if v.CheckPointInclusion(lat, lon) {
			included = true
		}
	}
	return included
}
func (a *AreaToAvoidItem) CheckPointInclusion(lat, lon float64) bool {

	if lat > a.L && lat < a.R && lon < a.T && lon > a.B {
		return true
	}

	return false
}

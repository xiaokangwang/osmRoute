package attributes

import "github.com/xiaokangwang/osmRoute/interfacew"

type specDef struct {
	timeFactor, costFactor, sustainableFactor      float64
	canWalk, canDrive, canBike, canPublicTransport bool
	areaToAvoid                                    AreaToAvoid
	bikeStations                                   BikeStation
	busInfo                                        BusInfo
}

func (s specDef) TimeFactor() float64 {
	return s.timeFactor
}

func (s specDef) CostFactor() float64 {
	return s.costFactor
}

func (s specDef) SustainableFactor() float64 {
	return s.sustainableFactor
}

func (s specDef) CanWalk() bool {
	return s.canWalk
}

func (s specDef) CanDrive() bool {
	return s.canDrive
}

func (s specDef) CanBike() bool {
	return s.canBike
}

func (s specDef) CanPublicTransport() bool {
	return s.canPublicTransport
}

func (s specDef) ListAllStations() []interfacew.MapLocation {
	return s.bikeStations.ListAllStations()
}

func (s specDef) WaitTime(NodeFID string) int64 {
	return s.busInfo.WaitTime(NodeFID)
}

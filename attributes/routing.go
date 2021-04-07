package attributes

import "strconv"

func ParseRoutingInputAttribute(AdditionalInfo map[string]string) (*specDef, error) {
	ret := &specDef{ //default values
		timeFactor:         0.33,
		costFactor:         0.33,
		sustainableFactor:  0.33,
		canWalk:            true,
		canDrive:           true,
		canBike:            true,
		canPublicTransport: true,
		areaToAvoid:        AreaToAvoid{},
		bikeStations:       BikeStation{},
		busInfo:            BusInfo{},
	}

	for key, value := range AdditionalInfo {
		switch key {
		case "objective_time":
			fallthrough
		case "objective_cost":
			fallthrough
		case "objective_sustainable":
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, newError("validation failed for condition ",
					key, " as it have a value of ", value,
					" however, it is supposed to be a float").Base(err)
			}
			if !(floatValue < 1 && floatValue > 0) {
				return nil, newError("validation failed for condition ",
					key, " as it have a value of ", value,
					" however, it is supposed to be [0,1]")
			}
			switch key {
			case "objective_time":
				ret.timeFactor = floatValue
			case "objective_cost":
				ret.costFactor = floatValue
			case "objective_sustainable":
				ret.sustainableFactor = floatValue
			}
		case "can_publictrans":
			fallthrough
		case "can_walkLong":
			fallthrough
		case "can_drive":
			fallthrough
		case "can_bike":
			boolValue, err := strconv.ParseBool(value)
			if err != nil {
				return nil, newError("validation failed for condition ",
					key, " as it have a value of ", value,
					" however, it is supposed to be boolean")
			}
			switch key {
			case "can_publictrans":
				ret.canPublicTransport = boolValue
			case "can_walkLong":
				ret.canWalk = boolValue
			case "can_drive":
				ret.canDrive = boolValue
			case "can_bike":
				ret.canBike = boolValue
			}
		case "area_to_avoid":
			areaValue, err := NewAreaToAvoid(value)
			if err != nil {
				return nil, newError("Cannot accept value for area to avoid: ").Base(err)
			}
			ret.areaToAvoid = areaValue
		case "bike":
			bikeStations, err := NewBikeStationData(value)
			if err != nil {
				return nil, newError("Cannot accept value for area to avoid: ").Base(err)
			}
			ret.bikeStations = bikeStations
		case "bus":
			busInfo, err := NewBusInfo(value)
			if err != nil {
				return nil, newError("Cannot accept value for area to avoid: ").Base(err)
			}
			ret.busInfo = busInfo
		default:
			return nil, newError("validation failed for condition ",
				key, " as it have a value of ", value,
				" however, but this is not understood")
		}
	}
	return ret, nil
}

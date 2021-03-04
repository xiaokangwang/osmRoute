package attributes

import "strconv"

func CheckRoutingInputAttribute(AdditionalInfo map[string]string) error {
	for key, value := range AdditionalInfo {
		switch key {
		case "objective_time":
			fallthrough
		case "objective_cost":
			fallthrough
		case "objective_sustainable":
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return newError("validation failed for condition ",
					key, " as it have a value of ", value,
					" however, it is supposed to be a float").Base(err)
			}
			if floatValue > 1 || floatValue < 0 {
				return newError("validation failed for condition ",
					key, " as it have a value of ", value,
					" however, it is supposed to be [0,1]")
			}
		case "can_walkLong":
			fallthrough
		case "can_drive":
			fallthrough
		case "can_bike":
			if !(value == "true" || value == "false") {
				return newError("validation failed for condition ",
					key, " as it have a value of ", value,
					" however, it is supposed to be boolean")
			}
		default:
			return newError("validation failed for condition ",
				key, " as it have a value of ", value,
				" however, but this is not understood")
		}
	}
	return nil
}

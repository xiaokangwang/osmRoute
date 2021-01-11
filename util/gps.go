package util

import "math"

func GPStoMeter(lat1, lon1, lat2, lon2 float64) float64 {
	//From https://stackoverflow.com/questions/639695/how-to-convert-latitude-or-longitude-to-meters
	var R = 6378.137 // Radius of earth in KM
	var dLat = lat2*math.Pi/180 - lat1*math.Pi/180
	var dLon = lon2*math.Pi/180 - lon1*math.Pi/180
	var a = math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var d = R * c
	return d * 1000 // meters
}

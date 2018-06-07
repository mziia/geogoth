package geogoth

import (
	"math"
)

const rad = math.Pi / 180 // Convert degrees to radians

// EarthRadius radius of the Earth
const EarthRadius = 6371000.0

// DistancePointPointDeg count distance between two points (degree parameters)
func DistancePointPointDeg(y1, x1, y2, x2 float64) float64 {
	// x - latitude; y - longitude

	radLat1 := x1 * rad
	radLat2 := x2 * rad
	deltaX := (x2 - x1) * rad
	deltaY := (y2 - y1) * rad

	a := math.Sin(deltaX/2)*math.Sin(deltaX/2) +
		math.Cos(radLat1)*math.Cos(radLat2)*
			math.Sin(deltaY/2)*math.Sin(deltaY/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := EarthRadius * c // distance in meters
	return distance

}

// DistancePointPointRad count distance between two points (radians parameters)
func DistancePointPointRad(latA, lonA, latB, lonB float64) float64 {

	distance := math.Acos(math.Sin(latA)*math.Sin(latB)+math.Cos(latA)*math.Cos(latB)*math.Cos(lonB-lonA)) * EarthRadius
	return distance
}

// DistancePointMultipoint count distance between Point and MultiPoint
func DistancePointMultipoint(feature1, feature2 *Feature) float64 {

	var distance float64
	y1, x1 := GetPointCoordinates(feature1)             // Coordinates of Point
	coords := (feature2.Geom.Coordinates).([][]float64) // Convert interface to [][]float64
	distarr := make([]float64, 0)                       // Creating slice for distances

	for i := range coords {
		y2, x2 := GetTwoDimArrayCoordinates(feature2, i)
		distarr = append(distarr, DistancePointPointDeg(y1, x1, y2, x2)) // Adding distances to dastarr

	}

	// Searching for the smallest distance
	distance = distarr[0]
	for i := range distarr {
		if distarr[i] < distance {
			distance = distarr[i]

		}
	}

	return distance
}

// LineString distance

// Bearing Finds the bearing from one lat/lon point to another.
func Bearing(latA, lonA, latB, lonB float64) float64 {

	brng := math.Atan2(math.Sin(lonB-lonA)*math.Cos(latB), math.Cos(latA)*math.Sin(latB)-math.Sin(latA)*math.Cos(latB)*math.Cos(lonB-lonA))

	return brng
}

// DistancePointLine Calculates the shortest distance between Point and Line in meters
func DistancePointLine(plon, plat, lon1, lat1, lon2, lat2 float64) float64 {

	lat1 = lat1 * rad // Line
	lat2 = lat2 * rad // Line
	plat = plat * rad // Point
	lon1 = lon1 * rad // Line
	lon2 = lon2 * rad // Line
	plon = plon * rad // Point

	bear12 := Bearing(lat1, lon1, lat2, lon2)
	bear13 := Bearing(lat1, lon1, plat, plon)
	dis13 := DistancePointPointRad(lat1, lon1, plat, plon)

	var dxa float64
	var dxt float64
	var dis12 float64
	var dis14 float64

	if math.Abs(bear13-bear12) > (math.Pi / 2) {
		dxa = dis13

	} else {
		//  Finding the cross-track distance.
		dxt = math.Asin(math.Sin(dis13/EarthRadius)*math.Sin(bear13-bear12)) * EarthRadius

		dis12 = DistancePointPointRad(lat1, lon1, lat2, lon2)
		dis14 = math.Acos(math.Cos(dis13/EarthRadius)/math.Cos(dxt/EarthRadius)) * EarthRadius
		if dis14 > dis12 {
			dxa = DistancePointPointRad(lat2, lon2, plat, plon)
		} else {
			dxa = math.Abs(dxt)

		}
	}

	return dxa
}

// DistancePointLinstring finds the smallest distance between Point and LineString/ MultiLineString
func DistancePointLinstring(feature1, feature2 *Feature) float64 {
	var distance float64
	pointY, pointX := GetPointCoordinates(feature1) // Coordinates of Point
	distarr := make([]float64, 0)                   // Creating slice for distances between Point and edges of LineString

	switch feature2.Geom.Type {
	case "LineString": // LineString
		coords := (feature2.Geom.Coordinates).([][]float64) // Convert interface to [][]float64
		for i := range coords {
			if i > 0 {
				lineY0, lineX0 := GetTwoDimArrayCoordinates(feature2, i)   // Coordinates of LineString
				lineY1, lineX1 := GetTwoDimArrayCoordinates(feature2, i-1) // Coordinates of LineString

				distarr = append(distarr, DistancePointLine(pointY, pointX, lineY0, lineX0, lineY1, lineX1))
			}
		}

	case "MultiLineString": // MultiLineString
		coords := (feature2.Geom.Coordinates).([][][]float64) // Convert interface to [][][]float64

		for i := range coords {
			for j := range coords[i] {
				if j > 0 {
					lineY0, lineX0 := GetThreeDimArrayCoordinates(feature2, i, j)   // Coordinates of LineString
					lineY1, lineX1 := GetThreeDimArrayCoordinates(feature2, i, j-1) // Coordinates of LineString

					distarr = append(distarr, DistancePointLine(pointY, pointX, lineY0, lineX0, lineY1, lineX1))
				}
			}
		}
	}

	if len(distarr) == 1 { // if distarr array has only 1 smallest distance (line cocnsists of 2 points)
		distance = distarr[0] // the only distance is the distance between point and line
	} else { // if distarr has more than 2 points

		// Searching for the smallest distance
		distance = distarr[0]
		for i := range distarr {
			if distarr[i] < distance {
				distance = distarr[i]

			}
		}
	}

	return distance

}

// DistancePointPolygon finds the smallest distance between Point and Polygon
func DistancePointPolygon(feature1, feature2 *Feature) float64 {

	var distance float64
	pointY, pointX := GetPointCoordinates(feature1) // Coordinates of Point
	distarr := make([]float64, 0)                   // Creating slice for distances between Point and edges of LineString

	switch feature2.Geom.Type {

	case "Polygon":
		coords := (feature2.Geom.Coordinates).([][][]float64) // Convert interface to [][][]float64

		for i := range coords {
			for j := range coords[i] {
				if j > 0 {
					lineY0, lineX0 := GetThreeDimArrayCoordinates(feature2, i, j)   // Coordinates of LineString
					lineY1, lineX1 := GetThreeDimArrayCoordinates(feature2, i, j-1) // Coordinates of LineString

					distarr = append(distarr, DistancePointLine(pointY, pointX, lineY0, lineX0, lineY1, lineX1))
				}
			}
		}

	case "MultiPolygon":

		coords := (feature2.Geom.Coordinates).([][][][]float64) // Convert interface to [][][]float64

		for p := range coords {
			for i := range coords[p] {
				for j := range coords[p][i] {
					if j > 0 {
						lineY0, lineX0 := GetFourDimArrayCoordinates(feature2, p, i, j)   // Coordinates of LineString
						lineY1, lineX1 := GetFourDimArrayCoordinates(feature2, p, i, j-1) // Coordinates of LineString

						distarr = append(distarr, DistancePointLine(pointY, pointX, lineY0, lineX0, lineY1, lineX1))
					}
				}
			}
		}

	}

	// Searching for the smallest distance
	distance = distarr[0]
	for i := range distarr {
		if distarr[i] < distance {
			distance = distarr[i]

		}
	}

	return distance
}

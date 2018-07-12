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
	distarr := make([]float64, 0)                       // Creates slice for distances

	for i := range coords {
		y2, x2 := GetTwoDimArrayCoordinates(feature2, i)
		distarr = append(distarr, DistancePointPointDeg(y1, x1, y2, x2)) // Adds distances to dastarr

	}

	distance = MinDistance(distarr)

	return distance
}

// LineString distance

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
		//  Finds the cross-track distance.
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

// DistancePointLinstring finds the smallest distance between Point and LineString || MultiLineString
func DistancePointLinstring(feature1, feature2 *Feature) float64 {

	var distance float64
	pointY, pointX := GetPointCoordinates(feature1) // Coordinates of Point
	distarr := make([]float64, 0)                   // Creates slice for distances between Point and edges of LineString

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

		distance = MinDistance(distarr)
	}

	return distance

}

// DistancePointPolygon finds the smallest distance between Point and Polygon
func DistancePointPolygon(feature1, feature2 *Feature) float64 {

	var distance float64
	pointY, pointX := GetPointCoordinates(feature1) // Coordinates of Point
	distarr := make([]float64, 0)                   // Creates slice for distances between Point and edges of LineString

	switch feature2.Geom.Type {

	case "Polygon":
		coords := (feature2.Geom.Coordinates).([][][]float64) // Convert interface to [][][]float64
		if PIPJordanCurveTheorem(pointY, pointX, feature2.Geom.Coordinates) == true {
			distance = 0
		} else {
			for i := range coords {
				for j := range coords[i] {
					if j > 0 {
						lineY0, lineX0 := GetThreeDimArrayCoordinates(feature2, i, j)   // Coordinates of LineString
						lineY1, lineX1 := GetThreeDimArrayCoordinates(feature2, i, j-1) // Coordinates of LineString

						distarr = append(distarr, DistancePointLine(pointY, pointX, lineY0, lineX0, lineY1, lineX1))
					}
				}
			}
			distance = MinDistance(distarr)

		}

	case "MultiPolygon":
		coords := (feature2.Geom.Coordinates).([][][][]float64) // Convert interface to [][][]float64

		for p := range coords {
			if PIPJordanCurveTheorem(pointY, pointX, coords[p]) == true {
				distance = 0
				break
			} else {

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
			distance = MinDistance(distarr)
		}

	}

	return distance
}

// DistanceLineStringLineString counts the smallest distance between two LineStrings
func DistanceLineStringLineString(feature1, feature2 *Feature) float64 {

	var distance float64

	distarr := make([]float64, 0)
	coords1 := (feature1.Geom.Coordinates).([][]float64) // Convert interface to [][]float64
	coords2 := (feature2.Geom.Coordinates).([][]float64) // Convert interface to [][]float64

	coordsArr1 := make([][]float64, 0)
	coordsArr2 := make([][]float64, 0)

	for i := range coords1 { // Finds coords of the first LineString
		line1Y, line1X := GetTwoDimArrayCoordinates(feature1, i)
		coordsArr1 = append(coordsArr1, []float64{line1Y, line1X})

	}

	for i := range coords2 { // Finds coords of the second LineString
		line2Y, line2X := GetTwoDimArrayCoordinates(feature2, i)
		coordsArr2 = append(coordsArr2, []float64{line2Y, line2X})
	}

	for i := 0; i < len(coordsArr1)-1; i++ {

		line1Y1 := coordsArr1[i][0]
		line1X1 := coordsArr1[i][1]
		line1Y2 := coordsArr1[i+1][0]
		line1X2 := coordsArr1[i+1][1]

		for j := 0; j < len(coordsArr2)-1; j++ {

			line2Y1 := coordsArr2[j][0]
			line2X1 := coordsArr2[j][1]
			line2Y2 := coordsArr2[j+1][0]
			line2X2 := coordsArr2[j+1][1]

			distarr = append(distarr, DistanceLineLine(line1Y1, line1X1, line1Y2, line1X2, line2Y1, line2X1, line2Y2, line2X2))
		}

	}

	distance = MinDistance(distarr)

	return distance
}

// DistanceMultipointMultipoint counts distance between MultiPoint and MultiPoint
func DistanceMultipointMultipoint(feature1, feature2 *Feature) float64 {

	var distance float64

	mult1 := (feature1.Geom.Coordinates).([][]float64) // Convert interface to [][]float64
	mult2 := (feature2.Geom.Coordinates).([][]float64) // Convert interface to [][]float64

	distarr := make([]float64, 0) // Creates slice for distances

	for i := range mult1 {
		y1, x1 := GetTwoDimArrayCoordinates(feature1, i)

		for j := range mult2 {
			y2, x2 := GetTwoDimArrayCoordinates(feature2, j)
			distarr = append(distarr, DistancePointPointDeg(y1, x1, y2, x2)) // Adds distances to dastarr
		}
	}

	distance = MinDistance(distarr)

	return distance
}

// DistanceMultipointLinestring counts distance between Multipoint & Linestring
func DistanceMultipointLinestring(feature1, feature2 *Feature) float64 {
	var distance float64

	multpoint := (feature1.Geom.Coordinates).([][]float64) // Convert interface to [][]float64
	linestr := (feature2.Geom.Coordinates).([][]float64)   // Convert interface to [][]float64

	distarr := make([]float64, 0)      // Creates slice for distances
	lineCoords := make([][]float64, 0) // Creates slice for coords of the LineString

	for i := range linestr { // Finds coords of the LineString
		lineY, lineX := GetTwoDimArrayCoordinates(feature2, i)
		lineCoords = append(lineCoords, []float64{lineY, lineX})
	}

	for i := range multpoint {

		y, x := GetTwoDimArrayCoordinates(feature1, i) // Coordinates of Multipoint[i] point
		var lineY1, lineX1, lineY2, lineX2 float64     // Vars for Linestring's points

		for j := 0; j < len(lineCoords)-1; j++ {

			// for j := range lineCoords {

			lineY1 = lineCoords[j][0]
			lineX1 = lineCoords[j][1]
			lineY2 = lineCoords[j+1][0]
			lineX2 = lineCoords[j+1][1]

			distarr = append(distarr, DistancePointLine(y, x, lineY1, lineX1, lineY2, lineX2))

		}

	}

	distance = MinDistance(distarr)

	return distance
}

// DistanceMultiPointMultiLinestring counts distance between MultiPoint and MultiLineString
func DistanceMultiPointMultiLinestring(feature1, feature2 *Feature) float64 {

	var distance float64

	multpoint := (feature1.Geom.Coordinates).([][]float64)     // Convert interface to [][]float64
	multlinestr := (feature2.Geom.Coordinates).([][][]float64) // Convert interface to [][][]float64

	distarr := make([]float64, 0)         // Creates slice for distances
	lineCoords := make([][]float64, 0)    // Creates slice for coords of the line
	mlineCoords := make([][][]float64, 0) // Creates slice for coords of the MultiLineString

	for i := range multlinestr { // Finds coords of the MultiLineString
		for j := range multlinestr[i] {
			lineY, lineX := GetThreeDimArrayCoordinates(feature2, i, j)
			lineCoords = append(lineCoords, []float64{lineY, lineX})
		}
		mlineCoords = append(mlineCoords, lineCoords)

	}

	for i := range multpoint {
		y, x := GetTwoDimArrayCoordinates(feature1, i) // Coordinates of Multipoint[i] point

		for m := range mlineCoords {
			for j := 0; j < len(mlineCoords[m])-1; j++ {

				lineY1 := mlineCoords[m][j][0]
				lineX1 := mlineCoords[m][j][1]
				lineY2 := mlineCoords[m][j+1][0]
				lineX2 := mlineCoords[m][j+1][1]

				distarr = append(distarr, DistancePointLine(y, x, lineY1, lineX1, lineY2, lineX2))

			}

		}
	}
	distance = MinDistance(distarr)

	return distance
}

// DistanceMultiPointPolygon counts distance between MultiPoint and Polygon
func DistanceMultiPointPolygon(feature1, feature2 *Feature) float64 {

	var distance float64
	multipoint := (feature1.Geom.Coordinates).([][]float64) // Convert interface to [][]float64
	polygon := (feature2.Geom.Coordinates).([][][]float64)  // Convert interface to [][][]float64

	distarr := make([]float64, 0) // Creates slice for distances

	for i := range multipoint {
		yPoint, xPoint := GetTwoDimArrayCoordinates(feature1, i) // Coordinates of Multipoint[i] point

		if PIPJordanCurveTheorem(yPoint, xPoint, feature2.Geom.Coordinates) == true {
			distance = 0
			break

		} else {

			for j := range polygon {

				for p := 0; p < len(polygon[j])-1; p++ {

					yPol1 := polygon[j][p][0]
					xPol1 := polygon[j][p][1]
					yPol2 := polygon[j][p+1][0]
					xPol2 := polygon[j][p+1][1]

					distarr = append(distarr, DistancePointLine(yPoint, xPoint, yPol1, xPol1, yPol2, xPol2))
				}
			}
			distance = MinDistance(distarr)

		}
	}

	return distance
}

// DistanceMultiPointMultiPolygon counts distance between MultiPoint and MultiPolygon
func DistanceMultiPointMultiPolygon(feature1, feature2 *Feature) float64 {
	var distance float64
	multipoint := (feature1.Geom.Coordinates).([][]float64)      // Convert interface to [][]float64
	multpolygon := (feature2.Geom.Coordinates).([][][][]float64) // Convert interface to [][][][]float64

	distarr := make([]float64, 0) // Creates slice for distances

	for i := range multipoint {
		yPoint, xPoint := GetTwoDimArrayCoordinates(feature1, i) // Coordinates of Multipoint[i] point

		for m := range multpolygon {

			if PIPJordanCurveTheorem(yPoint, xPoint, multpolygon[m]) == true {
				distance = 0
				break

			} else {

				for j := range multpolygon[m] {

					for p := 0; p < len(multpolygon[m][j])-1; p++ {

						yPol1 := multpolygon[m][j][p][0]
						xPol1 := multpolygon[m][j][p][1]
						yPol2 := multpolygon[m][j][p+1][0]
						xPol2 := multpolygon[m][j][p+1][1]

						distarr = append(distarr, DistancePointLine(yPoint, xPoint, yPol1, xPol1, yPol2, xPol2))
					}
				}
				distance = MinDistance(distarr)

			}
		}
	}

	return distance

}

// DistanceLineStringMultiLineString counts distance between LineString and MultiLineString
func DistanceLineStringMultiLineString(feature1, feature2 *Feature) float64 {
	var distance float64

	linestr := (feature1.Geom.Coordinates).([][]float64)
	mlinestr := (feature2.Geom.Coordinates).([][][]float64)

	distarr := make([]float64, 0)         // Creates slice for distances
	lineStrCoords := make([][]float64, 0) // Creates slice for coords of the LineString

	for i := range linestr { // Finds coords of LineString
		lineY, lineX := GetTwoDimArrayCoordinates(feature1, i)
		lineStrCoords = append(lineStrCoords, []float64{lineY, lineX})
	}

	for i := 0; i < len(lineStrCoords)-1; i++ {
		yL1, xL1 := lineStrCoords[i][0], lineStrCoords[i][1]
		yL2, xL2 := lineStrCoords[i+1][0], lineStrCoords[i+1][1]

		for m := range mlinestr {

			for j := 0; j < len(mlinestr[m])-1; j++ {

				yM1, xM1 := mlinestr[m][j][0], mlinestr[m][j][1]
				yM2, xM2 := mlinestr[m][j+1][0], mlinestr[m][j+1][1]

				distarr = append(distarr, DistanceLineLine(yL1, xL1, yL2, xL2, yM1, xM1, yM2, xM2))

			}
		}
	}
	distance = MinDistance(distarr)

	return distance
}

// DistanceLineStringPolygon counts distance between LineString and  Polygon
func DistanceLineStringPolygon(feature1, feature2 *Feature) float64 {
	var distance float64

	linestr := (feature1.Geom.Coordinates).([][]float64)
	polyg := (feature2.Geom.Coordinates).([][][]float64)

	distarr := make([]float64, 0)         // Creates slice for distances
	lineStrCoords := make([][]float64, 0) // Creates slice for coords of the LineString

	for i := range linestr { // Finds coords of LineString
		lineY, lineX := GetTwoDimArrayCoordinates(feature1, i)
		lineStrCoords = append(lineStrCoords, []float64{lineY, lineX})
	}

	var pip bool
	// Test if any point of LineString is inside of the Polygon
	for i := range lineStrCoords {
		y, x := lineStrCoords[i][0], lineStrCoords[i][1]
		pip = PIPJordanCurveTheorem(y, x, polyg)
	}

	if pip == true {
		distance = 0
	} else {

		for i := 0; i < len(lineStrCoords)-1; i++ {
			yL1, xL1 := lineStrCoords[i][0], lineStrCoords[i][1]
			yL2, xL2 := lineStrCoords[i+1][0], lineStrCoords[i+1][1]

			for m := range polyg {

				for j := 0; j < len(polyg[m])-1; j++ {

					yP1, xP1 := polyg[m][j][0], polyg[m][j][1]
					yP2, xP2 := polyg[m][j+1][0], polyg[m][j+1][1]

					distarr = append(distarr, DistanceLineLine(yL1, xL1, yL2, xL2, yP1, xP1, yP2, xP2))

				}
			}
		}
		distance = MinDistance(distarr)
	}

	return distance
}

// DistanceLineStringMultiPolygon counts distance between LineString and  Polygon
func DistanceLineStringMultiPolygon(feature1, feature2 *Feature) float64 {
	var distance float64

	linestr := (feature1.Geom.Coordinates).([][]float64)
	mpolyg := (feature2.Geom.Coordinates).([][][][]float64)

	distarr := make([]float64, 0)         // Creates slice for distances
	lineStrCoords := make([][]float64, 0) // Creates slice for coords of the LineString

	for i := range linestr { // Finds coords of LineString
		lineY, lineX := GetTwoDimArrayCoordinates(feature1, i)
		lineStrCoords = append(lineStrCoords, []float64{lineY, lineX})
	}

	var pip bool
	// Test if any point of LineString is inside of the MultiPolygon
	for i := range lineStrCoords {
		y, x := lineStrCoords[i][0], lineStrCoords[i][1]

		for j := range mpolyg {
			pip = PIPJordanCurveTheorem(y, x, mpolyg[j])
		}
	}

	if pip == true {
		distance = 0
	} else {

		for i := 0; i < len(lineStrCoords)-1; i++ {
			yL1, xL1 := lineStrCoords[i][0], lineStrCoords[i][1]
			yL2, xL2 := lineStrCoords[i+1][0], lineStrCoords[i+1][1]

			for m := range mpolyg {

				for p := range mpolyg[m] {

					for j := 0; j < len(mpolyg[m][p])-1; j++ {

						yP1, xP1 := mpolyg[m][p][j][0], mpolyg[m][p][j][1]
						yP2, xP2 := mpolyg[m][p][j+1][0], mpolyg[m][p][j+1][1]

						distarr = append(distarr, DistanceLineLine(yL1, xL1, yL2, xL2, yP1, xP1, yP2, xP2))

					}
				}
			}
		}
		distance = MinDistance(distarr)
	}

	return distance
}

// DistanceMultiLineStringMultiLineString counts distance between MultiLineString and MultiLineString
func DistanceMultiLineStringMultiLineString(feature1, feature2 *Feature) float64 {

	var distance float64

	multlinestr := (feature1.Geom.Coordinates).([][][]float64)
	mlinestr := (feature2.Geom.Coordinates).([][][]float64)

	distarr := make([]float64, 0) // Creates slice for distances

	for ml := range multlinestr {
		for i := 0; i < len(multlinestr[ml])-1; i++ {

			yL1, xL1 := multlinestr[ml][i][0], multlinestr[ml][i][1]
			yL2, xL2 := multlinestr[ml][i+1][0], multlinestr[ml][i+1][1]

			for m := range mlinestr {
				for j := 0; j < len(mlinestr[m])-1; j++ {

					yM1, xM1 := mlinestr[m][j][0], mlinestr[m][j][1]
					yM2, xM2 := mlinestr[m][j+1][0], mlinestr[m][j+1][1]

					distarr = append(distarr, DistanceLineLine(yL1, xL1, yL2, xL2, yM1, xM1, yM2, xM2))

				}
			}
		}
	}
	distance = MinDistance(distarr)

	return distance
}

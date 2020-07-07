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
func DistancePointMultipoint(point *Point, mPoint *MultiPoint) float64 {

	var distance float64
	pointY, pointX := point.GetCoordinates()
	mpCoords := ((*mPoint).Coordinates()).([][]float64)

	distarr := make([]float64, 0) // Creates slice for distances

	for i := range mpCoords {
		mY, mX := mPoint.GetCoordinates(i)

		// Adds distances to dastarr
		distarr = append(distarr, DistancePointPointDeg(pointY, pointX, mY, mX))

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
func DistancePointLinstring(point *Point, lineStr *LineString) float64 {

	var distance float64
	pointY, pointX := point.GetCoordinates() // Coordinates of Point

	distarr := make([]float64, 0) // Creates slice for distances between Point and edges of LineString

	lineStrCoords := (lineStr.Coordinates()).([][]float64) // Convert interface to [][]float64

	for i := range lineStrCoords {
		if i > 0 {
			lineY0, lineX0 := lineStr.GetCoordinates(i)
			lineY1, lineX1 := lineStr.GetCoordinates(i - 1)

			distarr = append(distarr, DistancePointLine(pointY, pointX, lineY0, lineX0, lineY1, lineX1))

		}
	}

	if len(distarr) == 1 { // if distarr array has only 1 smallest distance (line cocnsists of 2 points)
		distance = distarr[0] // the only distance is the distance between point and line
	} else { // if distarr has more than 2 points
		distance = MinDistance(distarr)
	}

	return distance

}

// DistancePointMultiLineString finds the smallest distance between Point and MultiLineString
func DistancePointMultiLineString(point *Point, mLineString *MultiLineString) float64 {

	var distance float64
	pointY, pointX := point.GetCoordinates() // Coordinates of Point
	distarr := make([]float64, 0)            // Creates slice for distances between Point and edges of MultiLineString

	coords := (mLineString.Coordinates()).([][][]float64) // Convert interface to [][][]float64

	for i := range coords {
		for j := range coords[i] {
			if j > 0 {
				lineY0, lineX0 := mLineString.GetCoordinates(i, j)
				lineY1, lineX1 := mLineString.GetCoordinates(i, j-1)

				distarr = append(distarr, DistancePointLine(pointY, pointX, lineY0, lineX0, lineY1, lineX1))
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
func DistancePointPolygon(point *Point, polygon *Polygon) float64 {

	var distance float64
	pointY, pointX := point.GetCoordinates() // Coordinates of Point
	distarr := make([]float64, 0)            // Creates slice for distances between Point and edges of LineString

	coords := (polygon.Coordinates()).([][][]float64) // Convert interface to [][][]float64

	if PIPJordanCurveTheorem(pointY, pointX, polygon.Coordinates) == true {
		distance = 0
	} else {
		for i := range coords {
			for j := range coords[i] {
				if j > 0 {
					lineY0, lineX0 := polygon.GetCoordinates(i, j)
					lineY1, lineX1 := polygon.GetCoordinates(i, j-1)

					distarr = append(distarr, DistancePointLine(pointY, pointX, lineY0, lineX0, lineY1, lineX1))
				}
			}
		}
		distance = MinDistance(distarr)

	}

	return distance
}

// DistancePointMultiPolygon finds the smallest distance between Point and Polygon
func DistancePointMultiPolygon(point *Point, mPolygon *MultiPolygon) float64 {

	var distance float64
	pointY, pointX := point.GetCoordinates() // Coordinates of Point
	distarr := make([]float64, 0)

	coords := (mPolygon.Coordinates()).([][][][]float64) // Convert interface to [][][]float64

	for p := range coords {
		if PIPJordanCurveTheorem(pointY, pointX, coords[p]) == true {
			distance = 0
			break
		} else {

			for i := range coords[p] {
				for j := range coords[p][i] {
					if j > 0 {
						lineY0, lineX0 := mPolygon.GetCoordinates(p, i, j)
						lineY1, lineX1 := mPolygon.GetCoordinates(p, i, j-1)

						distarr = append(distarr, DistancePointLine(pointY, pointX, lineY0, lineX0, lineY1, lineX1))
					}
				}
			}
		}
		distance = MinDistance(distarr)
	}

	return distance

}

// DistanceLineStringLineString counts the smallest distance between two LineStrings
func DistanceLineStringLineString(lineString *LineString, lineStr *LineString) float64 {

	var distance float64
	distarr := make([]float64, 0)
	coords1 := (lineString.Coordinates()).([][]float64) // Convert interface to [][]float64
	coords2 := (lineStr.Coordinates()).([][]float64)    // Convert interface to [][]float64

	coordsArr1 := make([][]float64, 0)
	coordsArr2 := make([][]float64, 0)

	for i := range coords1 { // Finds coords of the first LineString
		line1Y, line1X := lineString.GetCoordinates(i)
		coordsArr1 = append(coordsArr1, []float64{line1Y, line1X})

	}

	for i := range coords2 { // Finds coords of the second LineString
		line2Y, line2X := lineStr.GetCoordinates(i)
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
func DistanceMultipointMultipoint(multiPoint *MultiPoint, mPoint *MultiPoint) float64 {

	var distance float64

	multiPointCoords := (multiPoint.Coordinates()).([][]float64) // Convert interface to [][]float64
	mPointCoords := (mPoint.Coordinates()).([][]float64)         // Convert interface to [][]float64

	distarr := make([]float64, 0) // Creates slice for distances

	for i := range multiPointCoords {
		y1, x1 := multiPoint.GetCoordinates(i)

		for j := range mPointCoords {
			y2, x2 := mPoint.GetCoordinates(j)
			distarr = append(distarr, DistancePointPointDeg(y1, x1, y2, x2)) // Adds distances to dastarr
		}
	}

	distance = MinDistance(distarr)

	return distance
}

// DistanceMultipointLinestring counts distance between Multipoint & Linestring
func DistanceMultipointLinestring(multiPoint *MultiPoint, lineStr *LineString) float64 {

	var distance float64

	multiPointCoords := (multiPoint.Coordinates()).([][]float64) // Convert interface to [][]float64
	lineStrCoords := (lineStr.Coordinates()).([][]float64)       // Convert interface to [][]float64

	distarr := make([]float64, 0)      // Creates slice for distances
	lineCoords := make([][]float64, 0) // Creates slice for coords of the LineString

	for i := range lineStrCoords { // Finds coords of the LineString
		lineY, lineX := lineStr.GetCoordinates(i)
		lineCoords = append(lineCoords, []float64{lineY, lineX})
	}

	for i := range multiPointCoords {
		y, x := multiPoint.GetCoordinates(i) // Coordinates of Multipoint[i] point

		for j := 0; j < len(lineCoords)-1; j++ {

			distarr = append(distarr, DistancePointLine(y, x, lineCoords[j][0], lineCoords[j][1], lineCoords[j+1][0], lineCoords[j+1][1]))

		}

	}

	distance = MinDistance(distarr)

	return distance
}

// DistanceMultiPointMultiLinestring counts distance between MultiPoint and MultiLineString
func DistanceMultiPointMultiLinestring(multiPoint *MultiPoint, multiLineStr *MultiLineString) float64 {

	var distance float64

	multiPointCoords := (multiPoint.Coordinates()).([][]float64)       // Convert interface to [][]float64
	multiLineStrCoords := (multiLineStr.Coordinates()).([][][]float64) // Convert interface to [][][]float64

	distarr := make([]float64, 0)         // Creates slice for distances
	lineCoords := make([][]float64, 0)    // Creates slice for coords of one line
	mlineCoords := make([][][]float64, 0) // Creates slice for coords of the MultiLineString

	for i := range multiLineStrCoords { // Finds coords of the MultiLineString
		for j := range multiLineStrCoords[i] {
			lineY, lineX := multiLineStr.GetCoordinates(i, j)
			lineCoords = append(lineCoords, []float64{lineY, lineX})
		}
		mlineCoords = append(mlineCoords, lineCoords)
		lineCoords = nil // empty slice

	}

	for i := range multiPointCoords {
		y, x := multiPoint.GetCoordinates(i) // Coordinates of Multipoint[i] point

		for m := range mlineCoords {
			for j := 0; j < len(mlineCoords[m])-1; j++ {

				distarr = append(distarr, DistancePointLine(y, x, mlineCoords[m][j][0], mlineCoords[m][j][1], mlineCoords[m][j+1][0], mlineCoords[m][j+1][1]))

			}

		}
	}
	distance = MinDistance(distarr)

	return distance
}

// DistanceMultiPointPolygon counts distance between MultiPoint and Polygon
func DistanceMultiPointPolygon(multiPoint *MultiPoint, polygon *Polygon) float64 {

	var distance float64

	multiPointCoords := (multiPoint.Coordinates()).([][]float64) // Convert interface to [][]float64
	polygonCoords := (polygon.Coordinates()).([][][]float64)     // Convert interface to [][][]float64

	distarr := make([]float64, 0) // Creates slice for distances

	for i := range multiPointCoords {
		yPoint, xPoint := multiPoint.GetCoordinates(i) // Coordinates of Multipoint[i] point

		if PIPJordanCurveTheorem(yPoint, xPoint, polygon.Coordinates) == true {
			distance = 0
			break

		} else {

			for j := range polygonCoords {

				for p := 0; p < len(polygonCoords[j])-1; p++ {

					distarr = append(distarr, DistancePointLine(yPoint, xPoint, polygonCoords[j][p][0], polygonCoords[j][p][1], polygonCoords[j][p+1][0], polygonCoords[j][p+1][1]))
				}
			}
			distance = MinDistance(distarr)

		}
	}

	return distance
}

// DistanceMultiPointMultiPolygon counts distance between MultiPoint and MultiPolygon
func DistanceMultiPointMultiPolygon(multiPoint *MultiPoint, multiPolygon *MultiPolygon) float64 {

	var distance float64

	multiPointCoords := (multiPoint.Coordinates()).([][]float64)         // Convert interface to [][]float64
	multiPolygonCoords := (multiPolygon.Coordinates()).([][][][]float64) // Convert interface to [][][][]float64

	distarr := make([]float64, 0) // Creates slice for distances

	for i := range multiPointCoords {
		yPoint, xPoint := multiPoint.GetCoordinates(i) // Coordinates of Multipoint[i] point

		for m := range multiPolygonCoords {

			if PIPJordanCurveTheorem(yPoint, xPoint, multiPolygonCoords[m]) == true {
				distance = 0
				break

			} else {

				for j := range multiPolygonCoords[m] {
					for p := 0; p < len(multiPolygonCoords[m][j])-1; p++ {

						distarr = append(distarr, DistancePointLine(yPoint, xPoint, multiPolygonCoords[m][j][p][0], multiPolygonCoords[m][j][p][1], multiPolygonCoords[m][j][p+1][0], multiPolygonCoords[m][j][p+1][1]))
					}
				}
				distance = MinDistance(distarr)

			}
		}
	}

	return distance

}

// DistanceLineStringMultiLineString counts distance between LineString and MultiLineString
func DistanceLineStringMultiLineString(lineString *LineString, multiLineStr *MultiLineString) float64 {

	var distance float64

	linestr := (lineString.Coordinates()).([][]float64)
	mlinestr := (multiLineStr.Coordinates()).([][][]float64)

	distarr := make([]float64, 0)         // Creates slice for distances
	lineStrCoords := make([][]float64, 0) // Creates slice for coords of the LineString

	for i := range linestr { // Finds coords of LineString
		lineY, lineX := lineString.GetCoordinates(i)
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
func DistanceLineStringPolygon(lineString *LineString, polygon *Polygon) float64 {
	var distance float64

	linestr := (lineString.Coordinates()).([][]float64)
	polyg := (polygon.Coordinates()).([][][]float64)

	distarr := make([]float64, 0)         // Creates slice for distances
	lineStrCoords := make([][]float64, 0) // Creates slice for coords of the LineString

	for i := range linestr { // Finds coords of LineString
		lineY, lineX := lineString.GetCoordinates(i)
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
func DistanceLineStringMultiPolygon(lineString *LineString, mPolygon *MultiPolygon) float64 {
	var distance float64

	linestr := (lineString.Coordinates()).([][]float64)
	mpolyg := (mPolygon.Coordinates()).([][][][]float64)

	distarr := make([]float64, 0)         // Creates slice for distances
	lineStrCoords := make([][]float64, 0) // Creates slice for coords of the LineString

	for i := range linestr { // Finds coords of LineString
		lineY, lineX := lineString.GetCoordinates(i)
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
func DistanceMultiLineStringMultiLineString(multiLineString *MultiLineString, multiLineStr *MultiLineString) float64 {

	var distance float64

	multlinestr := (multiLineString.Coordinates()).([][][]float64)
	mlinestr := (multiLineStr.Coordinates()).([][][]float64)

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

// DistanceMultiLineStringPolygon counts distance between MultiLineString and  Polygon
func DistanceMultiLineStringPolygon(multiLineString *MultiLineString, polygon *Polygon) float64 {
	var distance float64

	mlinestr := (multiLineString.Coordinates()).([][][]float64)
	polyg := (polygon.Coordinates()).([][][]float64)

	distarr := make([]float64, 0) // Creates slice for distances

	mlineCoords := make([][][]float64, 0) // Creates slice for coords of the MultiLineString
	lineCoords := make([][]float64, 0)    // Creates slice for coords of one line

	for i := range mlinestr { // Finds coords of the MultiLineString
		for j := range mlinestr[i] {
			lineY, lineX := multiLineString.GetCoordinates(i, j)
			lineCoords = append(lineCoords, []float64{lineY, lineX})
		}
		mlineCoords = append(mlineCoords, lineCoords)
		lineCoords = nil // empty slice

	}

	var pip bool
	// Test if any point of MultiLineString is inside of the Polygon
	for i := range mlineCoords {
		for j := range mlineCoords[i] {
			y, x := mlineCoords[i][j][0], mlineCoords[i][j][1]
			pip = PIPJordanCurveTheorem(y, x, polyg)
		}
	}

	if pip == true {
		distance = 0
	} else {

		for m := range mlineCoords {
			for i := 0; i < len(mlineCoords[m])-1; i++ {

				yL1, xL1 := mlineCoords[m][i][0], mlineCoords[m][i][1]
				yL2, xL2 := mlineCoords[m][i+1][0], mlineCoords[m][i+1][1]

				for m := range polyg {

					for j := 0; j < len(polyg[m])-1; j++ {

						yP1, xP1 := polyg[m][j][0], polyg[m][j][1]
						yP2, xP2 := polyg[m][j+1][0], polyg[m][j+1][1]

						distarr = append(distarr, DistanceLineLine(yL1, xL1, yL2, xL2, yP1, xP1, yP2, xP2))

					}
				}
			}
		}
		distance = MinDistance(distarr)
	}

	return distance
}

// DistanceMultiLineStringMultiPolygon counts distance between MultiLineString and  MultiPolygon
func DistanceMultiLineStringMultiPolygon(multiLineString *MultiLineString, multiPolygon *MultiPolygon) float64 {
	var distance float64

	mlinestr := (multiLineString.Coordinates()).([][][]float64)
	mpolyg := (multiPolygon.Coordinates()).([][][][]float64)

	distarr := make([]float64, 0)         // Creates slice for distances
	mlineCoords := make([][][]float64, 0) // Creates slice for coords of the MultiLineString
	lineCoords := make([][]float64, 0)    // Creates slice for coords of one line

	for i := range mlinestr { // Finds coords of the MultiLineString
		for j := range mlinestr[i] {
			lineY, lineX := multiLineString.GetCoordinates(i, j)
			lineCoords = append(lineCoords, []float64{lineY, lineX})
		}
		mlineCoords = append(mlineCoords, lineCoords)
		lineCoords = nil // empty slice

	}

	var pip bool
	// Test if any point of MultiLineString is inside of the MultiPolygon
	for i := range mlineCoords {
		for j := range mlineCoords[i] {
			y, x := mlineCoords[i][j][0], mlineCoords[i][j][1]

			for m := range mpolyg {
				pip = PIPJordanCurveTheorem(y, x, mpolyg[m])
			}
		}
	}

	if pip == true {
		distance = 0
	} else {

		for m := range mlineCoords {
			for i := 0; i < len(mlineCoords[m])-1; i++ {

				yL1, xL1 := mlineCoords[m][i][0], mlineCoords[m][i][1]
				yL2, xL2 := mlineCoords[m][i+1][0], mlineCoords[m][i+1][1]

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
		}
		distance = MinDistance(distarr)
	}
	return distance
}

// DistancePolygonPolygon counts distance between Polygon and Polygon
func DistancePolygonPolygon(polygon *Polygon, pol *Polygon) float64 {
	var distance float64

	polyg1 := (polygon.Coordinates()).([][][]float64)
	polyg2 := (pol.Coordinates()).([][][]float64)

	distarr := make([]float64, 0) // Creates slice for distances

	polyg1Coords := make([][]float64, 0) // Creates slice for coords of the first Polygon
	polyg2Coords := make([][]float64, 0) // Creates slice for coords of the second Polygon

	for i := range polyg1 { // Finds coords of the first Polygon
		for j := range polyg1[i] {
			y, x := polygon.GetCoordinates(i, j)
			polyg1Coords = append(polyg1Coords, []float64{y, x})
		}
	}

	for i := range polyg2 { // Finds coords of the second Polygon
		for j := range polyg2[i] {
			y, x := pol.GetCoordinates(i, j)
			polyg2Coords = append(polyg2Coords, []float64{y, x})
		}
	}

	var pip bool
	// Test if any point of Polygon 1 is inside of the Polygon 2
	for i := range polyg1Coords {
		pip = PIPJordanCurveTheorem(polyg1Coords[i][0], polyg1Coords[i][1], polyg2)
		if pip == true {
			break
		}
	}
	if pip == true {
		distance = 0
	} else {

		// Test if any point of Polygon 2 is inside of the Polygon 1
		for i := range polyg2Coords {
			pip = PIPJordanCurveTheorem(polyg2Coords[i][0], polyg2Coords[i][1], polyg1)

			if pip == true {
				break
			}
		}

		if pip == true {
			distance = 0

		} else {

			for p1 := range polyg1 {

				for i := 0; i < len(polyg1[p1])-1; i++ {
					y1P1, x1P1 := polyg1[p1][i][0], polyg1[p1][i][1]
					y2P1, x2P1 := polyg1[p1][i+1][0], polyg1[p1][i+1][1]

					for p2 := range polyg2 {

						for j := 0; j < len(polyg2[p2])-1; j++ {

							y1P2, x1P2 := polyg2[p2][j][0], polyg2[p2][j][1]
							y2P2, x2P2 := polyg2[p2][j+1][0], polyg2[p2][j+1][1]

							distarr = append(distarr, DistanceLineLine(y1P1, x1P1, y2P1, x2P1, y1P2, x1P2, y2P2, x2P2))
						}
					}
				}
			}
			distance = MinDistance(distarr)
		}
	}
	return distance
}

// DistancePolygonMultiPolygon  counts distance between Polygon and MultiPolygon
func DistancePolygonMultiPolygon(polygon *Polygon, multiPolygon *MultiPolygon) float64 {
	var distance float64

	polyg := (polygon.Coordinates()).([][][]float64)
	mpolyg := (multiPolygon.Coordinates()).([][][][]float64)

	distarr := make([]float64, 0) // Creates slice for distances

	polygCoords := make([][][]float64, 0)  // Creates slice for coords of the Polygon
	mpolygCoords := make([][][]float64, 0) // Creates slice for coords of the MultiPolygon
	plineCoords := make([][]float64, 0)    // Creates slice for coords of one line
	mlineCoords := make([][]float64, 0)    // Creates slice for coords of one line

	for i := range polyg { // Finds coords of the Polygon
		for j := range polyg[i] {
			lineY, lineX := polygon.GetCoordinates(i, j)
			plineCoords = append(plineCoords, []float64{lineY, lineX})
		}
		polygCoords = append(polygCoords, plineCoords)
		plineCoords = nil // empty slice

	}

	for m := range mpolyg {
		for p := range mpolyg[m] {
			for i := range mpolyg[m][p] {

				lineY, lineX := multiPolygon.GetCoordinates(m, p, i)
				mlineCoords = append(mlineCoords, []float64{lineY, lineX})
			}
			mpolygCoords = append(mpolygCoords, mlineCoords)
			mlineCoords = nil // empty slice

		}
	}

	var pip bool
	// Test if any point of Polygon is inside of the MultiPolygon
	for i := range polygCoords {
		for j := range polygCoords[i] {
			y, x := polygCoords[i][j][0], polygCoords[i][j][1]

			for m := range mpolyg {
				pip = PIPJordanCurveTheorem(y, x, mpolyg[m])
				if pip == true {
					break
				}
			}
		}
	}

	if pip == true {
		distance = 0
	} else {
		// Test if any point of MultiPolygon is inside of the Polygon
		for i := range mpolygCoords {
			for j := range mpolygCoords[i] {
				y, x := mpolygCoords[i][j][0], mpolygCoords[i][j][1]

				pip = PIPJordanCurveTheorem(y, x, polyg)
				if pip == true {
					break
				}
			}
		}
		if pip == true {
			distance = 0
		} else {
			for m := range polygCoords {
				for i := 0; i < len(polygCoords[m])-1; i++ {

					yL1, xL1 := polygCoords[m][i][0], polygCoords[m][i][1]
					yL2, xL2 := polygCoords[m][i+1][0], polygCoords[m][i+1][1]

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
			}
			distance = MinDistance(distarr)
		}
	}
	return distance
}

// DistanceMultiPolygonMultiPolygon counts distance between MultiPolygon and MultiPolygon
func DistanceMultiPolygonMultiPolygon(multiPolygon *MultiPolygon, multiPol *MultiPolygon) float64 {
	var distance float64

	mpolyg1 := (multiPolygon.Coordinates()).([][][][]float64)
	mpolyg2 := (multiPol.Coordinates()).([][][][]float64)

	distarr := make([]float64, 0)

	mpolyg1Coords := make([][][]float64, 0) // / Creates slice for coords of the MultiPolygon
	mline1Coords := make([][]float64, 0)    // Creates slice for coords of one line
	mpolyg2Coords := make([][][]float64, 0) // / Creates slice for coords of the MultiPolygon
	mline2Coords := make([][]float64, 0)    // Creates slice for coords of one line

	for m := range mpolyg1 {
		for p := range mpolyg1[m] {
			for i := range mpolyg1[m][p] {
				lineY, lineX := multiPolygon.GetCoordinates(m, p, i)
				mline1Coords = append(mline1Coords, []float64{lineY, lineX})
			}
			mpolyg1Coords = append(mpolyg1Coords, mline1Coords)
			mline1Coords = nil // empty slice
		}
	}

	for m := range mpolyg2 {
		for p := range mpolyg2[m] {
			for i := range mpolyg2[m][p] {
				lineY, lineX := multiPol.GetCoordinates(m, p, i)
				mline2Coords = append(mline2Coords, []float64{lineY, lineX})
			}
			mpolyg2Coords = append(mpolyg2Coords, mline2Coords)
			mline2Coords = nil // empty slice
		}
	}

	var pip bool
	// Test if any point of the firs MultiPolygon is inside of the second MultiPolygon

	for i := range mpolyg1Coords {
		for j := range mpolyg1Coords[i] {
			y, x := mpolyg1Coords[i][j][0], mpolyg1Coords[i][j][1]

			for m := range mpolyg2 {
				pip = PIPJordanCurveTheorem(y, x, mpolyg2[m])
				if pip == true {
					break
				}
			}
		}
	}
	if pip == true {
		distance = 0
	} else {

		for i := range mpolyg2Coords {
			for j := range mpolyg2Coords[i] {
				y, x := mpolyg2Coords[i][j][0], mpolyg2Coords[i][j][1]

				for m := range mpolyg1 {
					pip = PIPJordanCurveTheorem(y, x, mpolyg1[m])
					if pip == true {
						break
					}
				}
			}
		}
		if pip == true {
			distance = 0
		} else {
			for m1 := range mpolyg1 {
				for p1 := range mpolyg1[m1] {
					for i := 0; i < len(mpolyg1[m1][p1])-1; i++ {
						y1P1, x1P1 := mpolyg1[m1][p1][i][0], mpolyg1[m1][p1][i][1]
						y2P1, x2P1 := mpolyg1[m1][p1][i+1][0], mpolyg1[m1][p1][i+1][1]

						for m2 := range mpolyg2 {
							for p2 := range mpolyg2[m2] {
								for j := 0; j < len(mpolyg2[m2][p2])-1; j++ {

									y1P2, x1P2 := mpolyg2[m2][p2][j][0], mpolyg2[m2][p2][j][1]
									y2P2, x2P2 := mpolyg2[m2][p2][j+1][0], mpolyg2[m2][p2][j+1][1]

									distarr = append(distarr, DistanceLineLine(y1P1, x1P1, y2P1, x2P1, y1P2, x1P2, y2P2, x2P2))

								}
							}
						}
					}
				}
			}
			distance = MinDistance(distarr)
		}
	}
	return distance
}

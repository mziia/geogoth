package geogoth

import (
	"math"
)

// MinDistance searches for the smallest distance
func MinDistance(distarr []float64) float64 {
	distance := distarr[0]
	for i := range distarr {
		if distarr[i] < distance {
			distance = distarr[i]

		}
	}
	return distance
}

// NegToPosRad converts brng to positive if it's negative
func NegToPosRad(rad float64) float64 {
	if rad > 0 {
		return rad
	} else {
		return 2*math.Pi + rad
	}
}

// Bearing Finds the bearing from one lat/lon point to another.
func Bearing(latA, lonA, latB, lonB float64) float64 {

	brng := math.Atan2(math.Sin(lonB-lonA)*math.Cos(latB), math.Cos(latA)*math.Sin(latB)-math.Sin(latA)*math.Cos(latB)*math.Cos(lonB-lonA))

	return NegToPosRad(brng)

}

// PIPJordanCurveTheorem returns true if point(feature1) is inside of the polygon(feature2)
// returns false if point(feature1) is outside of the polygon(feature2)
// C/C++ algorithm implementation:  https://sidvind.com/wiki/Point-in-polygon:_Jordan_Curve_Theorem
func PIPJordanCurveTheorem(py, px float64, pol interface{}) bool {

	// pol - feature2.Geom.Coordinates
	polygon := (pol).([][][]float64) // Convert interface to [][][]float64

	var x1, x2, k float64
	crossing := 0

	for p := range polygon {
		for i := range polygon[p] {

			/* This is done to ensure that we get the same result when
			   the line goes from left to right and right to left */
			if polygon[p][i][1] < polygon[p][(i+1)%len(polygon[p])][1] {
				x1 = polygon[p][i][1]
				x2 = polygon[p][(i+1)%len(polygon[p])][1]
			} else {
				x1 = polygon[p][(i+1)%len(polygon[p])][1]
				x2 = polygon[p][i][1]

			}

			/* First check if the ray is possible to cross the line */
			if px > x1 && px <= x2 && (py < polygon[p][i][0] || py <= polygon[p][(i+1)%len(polygon[p])][0]) {
				eps := 0.000001

				/* Calculate the equation of the line */
				dx := polygon[p][(i+1)%len(polygon[p])][1] - polygon[p][i][1]
				dy := polygon[p][(i+1)%len(polygon[p])][0] - polygon[p][i][0]

				if math.Abs(dx) < eps {
					k = 999999999999999999
				} else {
					k = dy / dx
				}

				m := polygon[p][i][0] - k*polygon[p][i][1]

				/* Find if the ray crosses the line */
				y2 := k*px + m
				if py <= y2 {
					crossing++
				}
			}
		}

	}

	// fmt.Println(fmt.Sprintf("The point is crossing %v lines", crossing))

	if crossing%2 == 1 {
		return true // The Point is instide of the Polygon
	} else {
		return false // The Point is outside of the Polygon
	}

}

// PointInPolygon ...
func PointInPolygon(feature1, feature2 *Feature) bool {

	py, px := GetPointCoordinates(feature1) // Coords of Point
	pol := feature2.Geom.Coordinates

	return PIPJordanCurveTheorem(py, px, pol)

}

// LineLineIntersection returns true if lines intersectd; false if lines do not intersect
// Algorithm from https://ideone.com/PnPJgb
func LineLineIntersection(yA, xA, yB, xB, yC, xC, yD, xD float64) bool {

	yCmP, xCmP := yC-yA, xC-xA
	yr, xr := yB-yA, xB-xA
	ys, xs := yD-yC, xD-xC

	cmpXr := xCmP*yr - yCmP*xr
	cmpXs := xCmP*ys - yCmP*xs
	rXs := xr*ys - yr*xs

	if cmpXr == 0 {
		// Lines are collinear, and so intersect if they have any overlap
		return ((xC-xA < 0) != (xC-xB < 0)) || ((yC-yA < 0) != (yC-yB < 0))
	}

	if rXs == 0 {
		return false // Lines are parallel
	}

	rXsr := 1 / rXs
	t := cmpXs * rXsr
	u := cmpXr * rXsr

	return (t >= 0) && (t <= 1) && (u >= 0) && (u <= 1)
}

// DistanceLineLine finds distance between two lines
// func DistanceLineLine(feature1, feature2 *Feature) float64 {
func DistanceLineLine(line1Y1, line1X1, line1Y2, line1X2, line2Y1, line2X1, line2Y2, line2X2 float64) float64 {

	var distance float64

	if LineLineIntersection(line1Y1, line1X1, line1Y2, line1X2, line2Y1, line2X1, line2Y2, line2X2) == true {
		distance = 0
	} else {

		distarr := make([]float64, 0) // Creates slice for distances between Point and edges of LineString

		distarr = append(distarr, DistancePointLine(line1Y1, line1X1, line2Y1, line2X1, line2Y2, line2X2))
		distarr = append(distarr, DistancePointLine(line1Y2, line1X2, line2Y1, line2X1, line2Y2, line2X2))

		distarr = append(distarr, DistancePointLine(line2Y1, line2X1, line1Y1, line1X1, line1Y2, line1X2))
		distarr = append(distarr, DistancePointLine(line2Y2, line2X2, line1Y1, line1X1, line1Y2, line1X2))

		distance = MinDistance(distarr)
	}

	return distance

}

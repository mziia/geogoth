package geogoth

// PointPointIntersection ...
func PointPointIntersection(p1 Point, p2 Point) bool {
	// var inters bool

	coord1 := (p1.GetCoordinates()).([]float64)
	coord2 := (p2.GetCoordinates()).([]float64)

	return coord1[0] == coord2[0] && coord1[1] == coord2[1]
}

// PointMultiPointIntersection ...
func PointMultiPointIntersection(p Point, mp MultiPoint) bool {
	var inters bool

	coordp := (p.GetCoordinates()).([]float64)

	for i := range mp.Coords {
		if coordp[0] == mp.Coords[i][0] && coordp[1] == mp.Coords[i][1] {
			inters = true
			break
		} else {
			inters = false
		}
	}
	return inters
}

package geogoth

// PointPointIntersection ...
func PointPointIntersection(p1 Point, p2 Point) bool {
	var inters bool

	coord1 := (p1.GetCoordinates()).([]float64)
	coord2 := (p2.GetCoordinates()).([]float64)

	if coord1[0] == coord2[0] && coord1[1] == coord2[1] {
		inters = true
	} else {
		inters = false
	}

	return inters
}

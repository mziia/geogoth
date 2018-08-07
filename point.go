package geogoth

// Point ...
type Point struct {
	Y float64
	X float64
}

// NewPoint ...
func NewPoint(y, x float64) Point {
	return Point{
		Y: y,
		X: x,
	}
}

// GetCoordinates returns y,x of the Point
func (p Point) GetCoordinates() interface{} {

	return []float64{p.Y, p.X}
}

// GetType returns type of the Point (Point)
func (p Point) GetType() string {
	return "Point"
}

// GetLength returns length of the Point (0)
func (p Point) GetLength() float64 {
	return 0
}

// DistanceTo returns distance between two geo objects
func (p Point) DistanceTo(f Feature) float64 {

	var distance float64

	switch f.GetType() {
	case "Point":
		coord1 := (p.GetCoordinates()).([]float64)
		y1, x1 := coord1[0], coord1[1]

		coord2 := (f.GetCoordinates()).([]float64)
		y2, x2 := coord2[0], coord2[1]

		distance = DistancePointPointDeg(y1, x1, y2, x2)

	case "MultiPoint":
		mpoint := f.(*MultiPoint)
		distance = DistancePointMultipoint(p, *mpoint)

	case "LineString":
		lstr := f.(*LineString)
		distance = DistancePointLinstring(p, *lstr)

	case "MultiLineString":
		mlinestr := f.(*MultiLineString)
		distance = DistancePointMultiLineString(p, *mlinestr)

	}

	return distance
}

// // AsFeature ...
// func (p Point) AsFeature() Feature {
// 	var f Feature
// 	f = &p
// 	return f
// }

// // NewPointFeature ...
// func NewPointFeature(y, x float64) Feature {
// 	point := Point{
// 		Y: y,
// 		X: x,
// 	}
// 	return FeatureStruct{Feature: &point}
// }

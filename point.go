package geogoth

// Point Point structure
type Point struct {
	Y float64
	X float64
}

// NewPoint Creates new Point object
func NewPoint(y, x float64) Point {
	return Point{
		Y: y,
		X: x,
	}
}

// Coordinates returns y,x of the Point
// @ ToDo: Add pointers: func (p *Point) Coordinates() interface{} {
func (p Point) Coordinates() interface{} {
	return []float64{p.Y, p.X}
}

// GetCoordinates returns longitude, latitude of Point geom
func (p Point) GetCoordinates() (float64, float64) {
	// convert coordinates from interface to []float64
	coord := (p.Coordinates()).([]float64)
	lon := coord[0]
	lat := coord[1]

	return lon, lat
}

// Type returns type of the Point (Point)
func (p Point) Type() string {
	return "Point"
}

// Length returns length of the Point (0)
func (p Point) Length() float64 {
	return 0
}

// DistanceTo returns distance between two geo objects
func (p Point) DistanceTo(f Feature) float64 {

	var distance float64

	switch f.Type() {
	case "Point":
		coord1 := (p.Coordinates()).([]float64)
		y1, x1 := coord1[0], coord1[1]

		coord2 := (f.Coordinates()).([]float64)
		y2, x2 := coord2[0], coord2[1]

		distance = DistancePointPointDeg(y1, x1, y2, x2)

	case "MultiPoint":
		mpoint := f.(MultiPoint)
		distance = DistancePointMultipoint(&p, &mpoint)

	case "LineString":
		lstr := f.(LineString)
		distance = DistancePointLinstring(&p, &lstr)

	case "MultiLineString":
		mlinestr := f.(MultiLineString)
		distance = DistancePointMultiLineString(&p, &mlinestr)

	case "Polygon":
		polygon := f.(Polygon)
		distance = DistancePointPolygon(&p, &polygon)

	case "MultiPolygon":
		mpolyg := f.(MultiPolygon)
		distance = DistancePointMultiPolygon(&p, &mpolyg)
	}

	return distance
}

// IntersectsWith returns true if geoObject intersects with Feature(another geoObject)
func (p Point) IntersectsWith(f Feature) bool {
	var intersection bool

	switch f.Type() {
	case "Point":
		point := f.(*Point)
		intersection = PointPointIntersection(p, *point)

	case "MultiPoint":
		mpoint := f.(*MultiPoint)
		intersection = PointMultiPointIntersection(p, *mpoint)

	case "LineString":
		lstr := f.(*LineString)
		if p.DistanceTo(lstr) == 0 {
			intersection = true
		} else {
			intersection = false
		}

	case "MultiLineString":
		mlstr := f.(*MultiLineString)
		if p.DistanceTo(mlstr) == 0 {
			intersection = true
		} else {
			intersection = false
		}

	case "Polygon":
		polygon := f.(*Polygon)
		if p.DistanceTo(polygon) == 0 {
			intersection = true
		} else {
			intersection = false
		}

	case "MultiPolygon":
		mpolyg := f.(*MultiPolygon)
		if p.DistanceTo(mpolyg) == 0 {
			intersection = true
		} else {
			intersection = false
		}
	}

	return intersection

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

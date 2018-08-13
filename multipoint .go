package geogoth

// MultiPoint ...
type MultiPoint struct {
	Coords [][]float64
}

// NewMultiPoint creates MultiPoint
func NewMultiPoint(coords [][]float64) MultiPoint {
	return MultiPoint{
		Coords: coords,
	}
}

// GetCoordinates returns array of longitude, latitude of the MultiPoint
func (m MultiPoint) GetCoordinates() interface{} {
	return m.Coords // longitude (Y), latitude (X)
}

// GetType returns type of the MultiPoint (MultiPoint)
func (m MultiPoint) GetType() string {
	return "MultiPoint"
}

// GetLength returns length of the MultiPoint
func (m MultiPoint) GetLength() float64 {
	return 0
}

// DistanceTo returns distance between two geo objects
func (m MultiPoint) DistanceTo(f Feature) float64 {

	var distance float64

	switch f.GetType() {
	case "Point":
		point := f.(*Point)
		distance = DistancePointMultipoint(*point, m)

	case "MultiPoint":
		mpoint := f.(*MultiPoint)
		distance = DistanceMultipointMultipoint(m, *mpoint)

	case "LineString":
		linestr := f.(*LineString)
		distance = DistanceMultipointLinestring(m, *linestr)

	case "MultiLineString":
		mlinestr := f.(*MultiLineString)
		distance = DistanceMultiPointMultiLinestring(m, *mlinestr)

	case "Polygon":
		polyg := f.(*Polygon)
		distance = DistanceMultiPointPolygon(m, *polyg)

	case "MultiPolygon":
		mpol := f.(*MultiPolygon)
		distance = DistanceMultiPointMultiPolygon(m, *mpol)
	}

	return distance
}

// IntersectsWith returns true if geoObject intersects with Feature
func (m MultiPoint) IntersectsWith(f Feature) bool {
	var intersection bool

	switch f.GetType() {
	case "Point":
		// point := f.(*Point)

	case "MultiPoint":
		// mpoint := f.(*MultiPoint)

	case "LineString":
		// lstr := f.(*LineString)

	case "MultiLineString":
		// mlinestr := f.(*MultiLineString)

	case "Polygon":
		// polygon := f.(*Polygon)

	case "MultiPolygon":
		// mpolyg := f.(*MultiPolygon)
	}

	return intersection

}

package geogoth

// LineString ...
type LineString struct {
	Coords [][]float64
}

// NewLineString creates LineString
func NewLineString(coords [][]float64) LineString {
	return LineString{
		Coords: coords,
	}
}

// Coordinates returns array of longitude, latitude of the LineString
func (l LineString) Coordinates() interface{} {
	return l.Coords // longitude (Y), latitude (X)
}

// GetCoordinates returns array of longitude, latitude of LineString
// coordnum - index of coordinate arr
func (l LineString) GetCoordinates(coordnum int) (float64, float64) {
	coords := (l.Coordinates()).([][]float64)

	lon := coords[coordnum][0]
	lat := coords[coordnum][1]

	return lon, lat // longitude (Y), latitude (X)
}

// Type returns type of the LineString (LineString)
func (l LineString) Type() string {
	return "LineString"
}

// Length returns length of the LineString
func (l LineString) Length() float64 {
	return LineStringLength(&l)
}

// DistanceTo returns distance between two geo objects
func (l LineString) DistanceTo(f Feature) float64 {

	var distance float64

	switch f.Type() {
	case "Point":
		point := f.(Point)
		distance = DistancePointLinstring(&point, &l)

	case "MultiPoint":
		mpoint := f.(MultiPoint)
		distance = DistanceMultipointLinestring(&mpoint, &l)

	case "LineString":
		lstr := f.(LineString)
		distance = DistanceLineStringLineString(&l, &lstr)

	case "MultiLineString":
		mlinestr := f.(MultiLineString)
		distance = DistanceLineStringMultiLineString(&l, &mlinestr)

	case "Polygon":
		pol := f.(Polygon)
		distance = DistanceLineStringPolygon(&l, &pol)

	case "MultiPolygon":
		mpol := f.(*MultiPolygon)
		distance = DistanceLineStringMultiPolygon(l, *mpol)
	}

	return distance
}

// IntersectsWith returns true if geoObject intersects with Feature
func (l LineString) IntersectsWith(f Feature) bool {
	var intersection bool

	switch f.Type() {
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

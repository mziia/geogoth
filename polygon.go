package geogoth

// Polygon ...
type Polygon struct {
	Coords [][][]float64
}

// NewPolygon creates Polygon
func NewPolygon(coords [][][]float64) Polygon {
	return Polygon{
		Coords: coords,
	}
}

// Coordinates returns array of longitude, latitude of the Polygon
func (p Polygon) Coordinates() interface{} {
	return p.Coords // longitude (Y), latitude (X)
}

// GetCoordinates returns array of longitude, latitude of  Polygon
// coordnum - index of coordinate arr
func (p Polygon) GetCoordinates(setnum, coordnum int) (float64, float64) {
	coords := (p.Coordinates()).([][][]float64)

	lon := coords[setnum][coordnum][0]
	lat := coords[setnum][coordnum][1]

	return lon, lat // longitude (Y), latitude (X)
}

// Type returns type of the Polygon (Polygon)
func (p Polygon) Type() string {
	return "Polygon"
}

// Length returns length of the Polygon
func (p Polygon) Length() float64 {
	return PolygonLength(p)
}

// DistanceTo returns distance between two geo objects
func (p Polygon) DistanceTo(f Feature) float64 {

	var distance float64

	switch f.Type() {
	case "Point":
		point := f.(*Point)
		distance = DistancePointPolygon(*point, p)

	case "MultiPoint":
		mp := f.(*MultiPoint)
		distance = DistanceMultiPointPolygon(*mp, p)

	case "LineString":
		lstr := f.(*LineString)
		distance = DistanceLineStringPolygon(*lstr, p)

	case "MultiLineString":
		mlstr := f.(*MultiLineString)
		distance = DistanceMultiLineStringPolygon(*mlstr, p)

	case "Polygon":
		pol := f.(*Polygon)
		distance = DistancePolygonPolygon(p, *pol)

	case "MultiPolygon":
		mplgn := f.(*MultiPolygon)
		distance = DistancePolygonMultiPolygon(p, *mplgn)
	}

	return distance
}

// IntersectsWith returns true if geoObject intersects with Feature
func (p Polygon) IntersectsWith(f Feature) bool {
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

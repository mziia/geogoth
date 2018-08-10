package geogoth

// MultiLineString ...
type MultiLineString struct {
	Coords [][][]float64
}

// NewMultiLineString creates MultiLineString
func NewMultiLineString(coords [][][]float64) MultiLineString {
	return MultiLineString{
		Coords: coords,
	}
}

// GetCoordinates returns array of longitude, latitude of the MultiLineString
func (m MultiLineString) GetCoordinates() interface{} {
	return m.Coords // longitude (Y), latitude (X)
}

// GetType returns type of the MultiLineString (MultiLineString)
func (m MultiLineString) GetType() string {
	return "MultiLineString"
}

// GetLength returns length of the MultiLineString
func (m MultiLineString) GetLength() float64 {
	return MultiLineStringLength(m)
}

// DistanceTo returns distance between two geo objects
func (m MultiLineString) DistanceTo(f Feature) float64 {

	var distance float64

	switch f.GetType() {
	case "Point":
		point := f.(*Point)
		distance = DistancePointMultiLineString(*point, m)

	case "MultiPoint":
		mpoint := f.(*MultiPoint)
		distance = DistanceMultiPointMultiLinestring(*mpoint, m)

	case "LineString":
		lstr := f.(*LineString)
		distance = DistanceLineStringMultiLineString(*lstr, m)

	case "MultiLineString":
		mlstr := f.(*MultiLineString)
		distance = DistanceMultiLineStringMultiLineString(m, *mlstr)

	case "Polygon":
		pol := f.(*Polygon)
		distance = DistanceMultiLineStringPolygon(m, *pol)

	case "MultiPolygon":
		mpol := f.(*MultiPolygon)
		distance = DistanceMultiLineStringMultiPolygon(m, *mpol)

	}

	return distance
}

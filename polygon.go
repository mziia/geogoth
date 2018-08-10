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

// GetCoordinates returns array of longitude, latitude of the Polygon
func (p Polygon) GetCoordinates() interface{} {
	return p.Coords // longitude (Y), latitude (X)
}

// GetType returns type of the Polygon (Polygon)
func (p Polygon) GetType() string {
	return "Polygon"
}

// GetLength returns length of the Polygon
func (p Polygon) GetLength() float64 {
	return PolygonLength(p)
}

// DistanceTo returns distance between two geo objects
func (p Polygon) DistanceTo(f Feature) float64 {

	var distance float64

	switch f.GetType() {
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

	case "MultiPolygon":
	}

	return distance
}

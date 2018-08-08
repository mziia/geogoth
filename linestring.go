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

// GetCoordinates returns array of longitude, latitude of the LineString
func (l LineString) GetCoordinates() interface{} {
	return l.Coords // longitude (Y), latitude (X)
}

// GetType returns type of the LineString (LineString)
func (l LineString) GetType() string {
	return "LineString"
}

// GetLength returns length of the LineString
func (l LineString) GetLength() float64 {
	return LineStringLength(l)
}

// DistanceTo returns distance between two geo objects
func (l LineString) DistanceTo(f Feature) float64 {

	var distance float64

	switch f.GetType() {
	case "Point":
		point := f.(*Point)
		distance = DistancePointLinstring(*point, l)

	case "MultiPoint":
		mpoint := f.(*MultiPoint)
		distance = DistanceMultipointLinestring(*mpoint, l)

	case "LineString":

	case "MultiLineString":

	case "Polygon":

	case "MultiPolygon":
	}

	return distance
}

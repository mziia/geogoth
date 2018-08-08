package geogoth

// MultiPolygon ...
type MultiPolygon struct {
	Coords [][][][]float64
}

// NewMultiPolygon creates MultiPolygon
func NewMultiPolygon(coords [][][][]float64) MultiPolygon {
	return MultiPolygon{
		Coords: coords,
	}
}

// GetCoordinates returns array of longitude, latitude of the MultiPolygon
func (m MultiPolygon) GetCoordinates() interface{} {
	return m.Coords
}

// GetType returns type of the MultiPolygon (MultiPolygon)
func (m MultiPolygon) GetType() string {
	return "MultiPolygon"
}

// GetLength returns length of the MultiPolygon
func (m MultiPolygon) GetLength() float64 {
	return MultipolygonLength(m)
}

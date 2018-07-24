package geogoth

// NewPoint create Point with given coordinates
func NewPoint(coordinate []float64) *Geometry {
	return &Geometry{
		Type:        Point,
		Coordinates: coordinate,
	}
}

// GetPointCoordinates returns longitude, latitude of Point geom
func GetPointCoordinates(feature *Feature) (float64, float64) {
	// long = longitude  (Y)
	// lat = latitude  (X)

	// convert coordinates from interface to []float64
	coord := (feature.Geom.Coordinates).([]float64)
	lon := coord[0]
	lat := coord[1]

	return lon, lat
}

// NewMultiPoint create NewMultiPoint with given coordinates
func NewMultiPoint(coordinates [][]float64) *Geometry {
	return &Geometry{
		Type:        MultiPoint,
		Coordinates: coordinates,
	}
}

// NewLineString create NewLineString with given coordinates
func NewLineString(coordinates [][]float64) *Geometry {
	return &Geometry{
		Type:        LineString,
		Coordinates: coordinates,
	}
}

// GetTwoDimArrayCoordinates returns array of longitude, latitude of Two-dimensional arrays (MultiPoint, Linestring)
// coordnum - index of coordinate arr
func GetTwoDimArrayCoordinates(feature *Feature, coordnum int) (float64, float64) {
	coords := (feature.Geom.Coordinates).([][]float64)

	lon := coords[coordnum][0]
	lat := coords[coordnum][1]

	return lon, lat // longitude (Y), latitude (X)
}

// NewMultiLineString create NewMultiLineString with given coordinates
func NewMultiLineString(lines [][][]float64) *Geometry {
	return &Geometry{
		Type:        MultiLineString,
		Coordinates: lines,
	}
}

// NewPolygon create NewPolygon with given coordinates
func NewPolygon(polygon [][][]float64) *Geometry {
	return &Geometry{
		Type:        Polygon,
		Coordinates: polygon,
	}
}

// GetThreeDimArrayCoordinates returns array of longitude, latitude of Three-dimensional arrays (MultiLineString, Polygon)
// coordnum - index of coordinate arr
func GetThreeDimArrayCoordinates(feature *Feature, setnum, coordnum int) (float64, float64) {
	coords := (feature.Geom.Coordinates).([][][]float64)

	lon := coords[setnum][coordnum][0]
	lat := coords[setnum][coordnum][1]

	return lon, lat // longitude (Y), latitude (X)
}

// NewMultiPolygon create NewMultiPolygon with given coordinates
func NewMultiPolygon(polygons [][][][]float64) *Geometry {
	return &Geometry{
		Type:        MultiPolygon,
		Coordinates: polygons,
	}
}

// GetFourDimArrayCoordinates returns array of longitude, latitude of Three-dimensional arrays (MultiLineString, Polygon)
// coordnum - index of coordinate arr
func GetFourDimArrayCoordinates(feature *Feature, setsnum, setnum, coordnum int) (float64, float64) {
	coords := (feature.Geom.Coordinates).([][][][]float64)

	lon := coords[setsnum][setnum][coordnum][0]
	lat := coords[setsnum][setnum][coordnum][1]

	return lon, lat // longitude (Y), latitude (X)
}

package geojson

// NewPoint create Point with given coordinates
func NewPoint(coordinate []float64) *Geometry {
	return &Geometry{
		Type:        Point,
		Coordinates: coordinate,
	}
}

// NewMultiPoint create NewMultiPoint with given coordinates
func NewMultiPoint(coordinates ...[]float64) *Geometry {
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

// NewMultiLineString create NewMultiLineString with given coordinates
func NewMultiLineString(lines ...[][]float64) *Geometry {
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

// NewMultiPolygon create NewMultiPolygon with given coordinates
func NewMultiPolygon(polygons ...[][][]float64) *Geometry {
	return &Geometry{
		Type:        MultiPolygon,
		Coordinates: polygons,
	}
}

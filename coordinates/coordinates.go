package coordinates

import (
	"geojson/structs"
)

// NewPoint create Point with given coordinates
func NewPoint(coordinate []float64) *structs.Geometry {
	return &structs.Geometry{
		Type:        structs.Point,
		Coordinates: coordinate,
	}
}

// NewMultiPoint create NewMultiPoint with given coordinates
func NewMultiPoint(coordinates ...[]float64) *structs.Geometry {
	return &structs.Geometry{
		Type:        structs.MultiPoint,
		Coordinates: coordinates,
	}
}

// NewLineString create NewLineString with given coordinates
func NewLineString(coordinates [][]float64) *structs.Geometry {
	return &structs.Geometry{
		Type:        structs.LineString,
		Coordinates: coordinates,
	}
}

// NewMultiLineString create NewMultiLineString with given coordinates
func NewMultiLineString(lines ...[][]float64) *structs.Geometry {
	return &structs.Geometry{
		Type:        structs.MultiLineString,
		Coordinates: lines,
	}
}

// NewPolygon create NewPolygon with given coordinates
func NewPolygon(polygon [][][]float64) *structs.Geometry {
	return &structs.Geometry{
		Type:        structs.Polygon,
		Coordinates: polygon,
	}
}

// NewMultiPolygon create NewMultiPolygon with given coordinates
func NewMultiPolygon(polygons ...[][][]float64) *structs.Geometry {
	return &structs.Geometry{
		Type:        structs.MultiPolygon,
		Coordinates: polygons,
	}
}

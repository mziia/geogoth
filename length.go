package geogoth

// Length counts lenght of geo object
func (feature *Feature) Length() float64 {
	var length float64

	switch GetGeoType(feature) {

	case "Point": //  Point length
		length = 0
	case "MultiPoint": //  Point length
		length = 0
	case "LineString": //  LineString length
		length = LineStringLength(feature)
	case "MultiLineString": // MultiLineString length
		length = MultiLineStringLength(feature)
	case "Polygon": // Polygon length
		length = PolygonLength(feature)
	case "MultiPolygon": // MultiPolygon length
		length = MultipolygonLength(feature)
	}

	return length
}

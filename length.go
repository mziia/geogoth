package geogoth

// Length counts lenght of geo object
func (feature *Feature) Length() float64 {
	var length float64

	switch GetGeoType(feature) {

	case "LineString": // Point & LineString
		length = LineStringLength(feature)
	case "MultiLineString": // Point & MultiLineString
		length = MultiLineStringLength(feature)
	case "Polygon": // Point & Polygon
		length = PolygonLength(feature)
	case "MultiPolygon": // Point & MultiPolygon
		length = MultipolygonLength(feature)
	}

	return length
}

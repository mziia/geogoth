package geogoth

// Distance retern between teo given geo objects
func Distance(feature1, feature2 *Feature) float64 {
	// long = longitude = 37... (долгота Y)
	// lat = latitude = 55.... (широта X)

	var distance float64
	switch GetGeoType(feature1) {

	case "Point":

		switch GetGeoType(feature2) {

		case "Point": // Point & Point

			y1, x1 := GetPointCoordinates(feature1) // Coordinates of Point
			y2, x2 := GetPointCoordinates(feature2)
			distance = DistancePointPointDeg(y1, x1, y2, x2)

		case "MultiPoint": // Point & MultiPoint
			distance = DistancePointMultipoint(feature1, feature2)
		case "LineString": // Point & LineString
			distance = DistancePointLinstring(feature1, feature2)
		case "MultiLineString": // Point & MultiLineString
			distance = DistancePointLinstring(feature1, feature2)
		case "Polygon": // Point & Polygon
			distance = DistancePointPolygon(feature1, feature2)
		case "MultiPolygon": // Point & MultiPolygon
		}

		// =============================================================================

	case "MultiPoint": //

		switch GetGeoType(feature2) {

		case "Point": // MultiPoint & Point
			distance = DistancePointMultipoint(feature2, feature1)

		case "MultiPoint": // MultiPoint & MultiPoint
		case "LineString": // MultiPoint & LineString
		case "MultiLineString": // MultiLineString &
		case "Polygon": // MultiPoint & Polygon
		case "MultiPolygon": // MultiPoint & MultiPolygon
		}

		// =============================================================================

	case "LineString":
		switch GetGeoType(feature2) {
		case "Point": // LineString & Point
			distance = DistancePointLinstring(feature2, feature1)

		case "MultiPoint": // LineString &
		case "LineString": // LineString &
		case "MultiLineString": // LineString &
		case "Polygon": // LineString &
		case "MultiPolygon": // LineString &
		}

		// =============================================================================

	case "MultiLineString":
		switch GetGeoType(feature2) {
		case "Point": // MultiLineString &
		case "MultiPoint": // MultiLineString &
		case "LineString": // MultiLineString &
		case "MultiLineString": // MultiLineString &
		case "Polygon": // MultiLineString &
		case "MultiPolygon": // MultiLineString &
		}

		// =============================================================================

	case "Polygon":
		switch GetGeoType(feature2) {
		case "Point": // Polygon & Point
			distance = DistancePointPolygon(feature2, feature1)

		case "MultiPoint": // Polygon &
		case "LineString": // Polygon &
		case "MultiLineString": // Polygon &
		case "Polygon": // Polygon &
		case "MultiPolygon": // Polygon &
		}

		// =============================================================================

	case "MultiPolygon":
		switch GetGeoType(feature2) {
		case "Point": // MultiPolygon &
		case "MultiPoint": // MultiPolygon &
		case "LineString": // MultiPolygon &
		case "MultiLineString": // MultiPolygon &
		case "Polygon": // MultiPolygon &
		case "MultiPolygon": // MultiPolygon &
		}

	}
	return distance
}

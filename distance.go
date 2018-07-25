package geogoth

// Distance retern between two given geo objects
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
			distance = DistancePointMultiLineString(feature1, feature2)
		case "Polygon": // Point & Polygon
			distance = DistancePointPolygon(feature1, feature2)
		case "MultiPolygon": // Point & MultiPolygon
			distance = DistancePointMultiPolygon(feature1, feature2)

		}

		// =============================================================================

	case "MultiPoint": //
		switch GetGeoType(feature2) {

		case "Point": // MultiPoint & Point
			distance = DistancePointMultipoint(feature2, feature1)
		case "MultiPoint": // MultiPoint & MultiPoint
			distance = DistanceMultipointMultipoint(feature1, feature2)
		case "LineString": // MultiPoint & LineString
			distance = DistanceMultipointLinestring(feature1, feature2)
		case "MultiLineString": // MultiPoint & MultiLineString
			distance = DistanceMultiPointMultiLinestring(feature1, feature2)
		case "Polygon": // MultiPoint & Polygon
			distance = DistanceMultiPointPolygon(feature1, feature2)
		case "MultiPolygon": // MultiPoint & MultiPolygon
			distance = DistanceMultiPointMultiPolygon(feature1, feature2)

		}

		// =============================================================================

	case "LineString":
		switch GetGeoType(feature2) {

		case "Point": // LineString & Point
			distance = DistancePointLinstring(feature2, feature1)
		case "MultiPoint": // LineString & MultiPoint
			distance = DistanceMultipointLinestring(feature2, feature1)
		case "LineString": // LineString & LineString
			distance = DistanceLineStringLineString(feature2, feature1)
		case "MultiLineString": // LineString & MultiLineString
			distance = DistanceLineStringMultiLineString(feature1, feature2)
		case "Polygon": // LineString & Polygon
			distance = DistanceLineStringPolygon(feature1, feature2)
		case "MultiPolygon": // LineString & MultiPolygon
			distance = DistanceLineStringMultiPolygon(feature1, feature2)

		}

		// =============================================================================

	case "MultiLineString":
		switch GetGeoType(feature2) {

		case "Point": // MultiLineString & Point
			distance = DistancePointMultiLineString(feature2, feature1)
		case "MultiPoint": // MultiLineString & MultiPoint
			distance = DistanceMultiPointMultiLinestring(feature2, feature1)
		case "LineString": // MultiLineString & LineString
			distance = DistanceLineStringMultiLineString(feature2, feature1)
		case "MultiLineString": // MultiLineString & MultiLineString
			distance = DistanceMultiLineStringMultiLineString(feature1, feature2)
		case "Polygon": // MultiLineString & Polygon
			distance = DistanceMultiLineStringPolygon(feature1, feature2)
		case "MultiPolygon": // MultiLineString & MultiPolygon
			distance = DistanceMultiLineStringMultiPolygon(feature1, feature2)

		}

		// =============================================================================

	case "Polygon":
		switch GetGeoType(feature2) {

		case "Point": // Polygon & Point
			distance = DistancePointPolygon(feature2, feature1)
		case "MultiPoint": // Polygon & MultiPoint
			distance = DistanceMultiPointPolygon(feature2, feature1)
		case "LineString": // Polygon & LineString
			distance = DistanceLineStringPolygon(feature2, feature1)
		case "MultiLineString": // Polygon & MultiLineString
			distance = DistanceMultiLineStringPolygon(feature2, feature1)
		case "Polygon": // Polygon & Polygon
			distance = DistancePolygonPolygon(feature1, feature2)
		case "MultiPolygon": // Polygon & MultiPolygon
			distance = DistancePolygonMultiPolygon(feature1, feature2)
		}

		// =============================================================================

	case "MultiPolygon":
		switch GetGeoType(feature2) {

		case "Point": // MultiPolygon & Point
			distance = DistancePointMultiPolygon(feature2, feature1)
		case "MultiPoint": // MultiPolygon & MultiPoint
			distance = DistanceMultiPointMultiPolygon(feature2, feature1)
		case "LineString": // MultiPolygon & LineString
			distance = DistanceLineStringMultiPolygon(feature2, feature1)
		case "MultiLineString": // MultiPolygon & MultiLineString
			distance = DistanceMultiLineStringMultiPolygon(feature2, feature1)
		case "Polygon": // MultiPolygon & Polygon
			distance = DistancePolygonMultiPolygon(feature2, feature1)
		case "MultiPolygon": // MultiPolygon &
			distance = DistanceMultiPolygonMultiPolygon(feature1, feature2)

		}

	}
	return distance
}

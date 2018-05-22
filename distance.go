package geogoth

// Distance between teo given geo objects
func Distance(object1, object2 *Feature) {

	switch GetGeoType(object1) {

	case "Point":
		switch GetGeoType(object2) {
		case "Point":
		case "MultiPoint":
		case "LineString":
		case "MultiLineString":
		case "Polygon":
		case "MultiPolygon":

		}

	case "MultiPoint":
		switch GetGeoType(object2) {
		case "Point":
		case "MultiPoint":
		case "LineString":
		case "MultiLineString":
		case "Polygon":
		case "MultiPolygon":
		}

	case "LineString":
		switch GetGeoType(object2) {
		case "Point":
		case "MultiPoint":
		case "LineString":
		case "MultiLineString":
		case "Polygon":
		case "MultiPolygon":
		}

	case "MultiLineString":
		switch GetGeoType(object2) {
		case "Point":
		case "MultiPoint":
		case "LineString":
		case "MultiLineString":
		case "Polygon":
		case "MultiPolygon":
		}

	case "Polygon":
		switch GetGeoType(object2) {
		case "Point":
		case "MultiPoint":
		case "LineString":
		case "MultiLineString":
		case "Polygon":
		case "MultiPolygon":
		}

	case "MultiPolygon":
		switch GetGeoType(object2) {
		case "Point":
		case "MultiPoint":
		case "LineString":
		case "MultiLineString":
		case "Polygon":
		case "MultiPolygon":
		}

	}

}

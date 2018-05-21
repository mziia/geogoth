package geogoth

// GeometryType тип для констант - типов геометрии
type GeometryType string

const (
	// Point ...
	Point GeometryType = "Point"
	// MultiPoint ...
	MultiPoint GeometryType = "MultiPoint"
	// LineString ...
	LineString GeometryType = "LineString"
	// MultiLineString ...
	MultiLineString GeometryType = "MultiLineString"
	// Polygon ...
	Polygon GeometryType = "Polygon"
	// MultiPolygon ...
	MultiPolygon GeometryType = "MultiPolygon"
)

// Geometry structure for geometry
type Geometry struct {
	Type            GeometryType    `json:"type"`
	Coordinates     interface{}     `json:"coordinates"`
	Point           []float64       `json:"point,omitempty"`
	MultiPoint      [][]float64     `json:"multipoint,omitempty"`
	LineString      [][]float64     `json:"lineString,omitempty"`
	MultiLineString [][][]float64   `json:"multiLineString,omitempty"`
	Polygon         [][][]float64   `json:"polygon,omitempty"`
	MultiPolygon    [][][][]float64 `json:"multiPolygon,omitempty"`
}

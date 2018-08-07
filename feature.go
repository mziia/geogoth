package geogoth

// Feature ...
type Feature interface {
	GetCoordinates() interface{}
	GetType() string
	GetLength() float64
	DistanceTo(Feature) float64
}

// // FeatureStruct ...
// type FeatureStruct struct {
// 	Feature
// }

// FeatureStruct ...
type FeatureStruct struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	ID         string                 `json:"id"`
	// Geom       Geometry               `json:"geometry"`
}

// FeatureCollection ...
type FeatureCollection struct {
	Type     string `json:"type"`
	Features []Feature
}

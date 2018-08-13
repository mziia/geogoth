package geogoth

// Feature ...
type Feature interface {
	GetCoordinates() interface{} // returns coordinates of geoObject
	GetType() string             // returns type of geoObject
	GetLength() float64          // returns length of geoObject
	DistanceTo(Feature) float64  // returns distance from geoObject to Feature

	IntersectsWith(Feature) bool // returns true if geoObject intersects with Feature

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

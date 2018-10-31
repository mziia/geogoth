package geogoth

import (
	"encoding/json"
)

// Features structure for geojson (feature collection)
type Features struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
}

// Feature structure for a single feature
type Feature struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	ID         string                 `json:"id"`
	Geom       Geometry               `json:"geometry"`
}

// NewFeatureCollection  creates feature collection
func NewFeatureCollection() *Features {
	return &Features{
		Type:     "FeatureCollection",
		Features: make([]*Feature, 0), // creating empty array for features
	}
}

// NewFeature creates new feature
func NewFeature() *Feature {
	return &Feature{
		Type: "Feature",
	}
}

// AddFeature adds feature to the feature collection
func (collection *Features) AddFeature(feature *Feature) *Features {
	collection.Features = append(collection.Features, feature)

	return collection
}

// SetProperty set properties
func (feature *Feature) SetProperty(key string, value interface{}) {
	if feature.Properties == nil {
		feature.Properties = make(map[string]interface{})
	}
	feature.Properties[key] = value
}

// SetID set ID
func (feature *Feature) SetID(id string) {
	feature.ID = id
}

// SetGeometry sets coordinates to the feature
func (feature *Feature) SetGeometry(geom *Geometry) {
	feature.Geom = *geom
}

//GeoUnmarshal ...
func GeoUnmarshal(byteJSON []byte) Features {
	var features Features
	json.Unmarshal(byteJSON, &features)

	return features
}

// GetGeoType returns type of geo object
func GetGeoType(feature *Feature) GeometryType {

	return feature.Geom.Type
}

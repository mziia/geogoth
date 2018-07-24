package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestNewFeatureCollection(t *testing.T) {
	assert := assert.New(t)
	collection := NewFeatureCollection()

	assert.True(collection.Type == "FeatureCollection")

}

func TestNewNewFeature(t *testing.T) {
	assert := assert.New(t)

	point := NewPoint([]float64{37.6175, 55.752})

	feature := NewFeature()
	feature.SetProperty("location", "Kremlin")
	feature.SetID("000")
	feature.SetGeometry(point)

	f := (feature.Geom.Coordinates).([]float64)
	assert.True(f[0] == 37.6175)
	assert.True(f[1] == 55.752)

}

func TestAddFeature(t *testing.T) {
	assert := assert.New(t)

	collection := NewFeatureCollection()

	point1 := NewPoint([]float64{37.6175, 55.752})
	point2 := NewPoint([]float64{37.6048, 55.7649})

	feature1 := NewFeature()
	feature1.SetProperty("location", "Kremlin")
	feature1.SetID("0001")
	feature1.SetGeometry(point1)

	feature2 := NewFeature()
	feature2.SetProperty("location", "Tverskaya")
	feature2.SetID("0002")
	feature2.SetGeometry(point2)

	collection.AddFeature(feature1)
	collection.AddFeature(feature2)

	assert.True(collection.Type == "FeatureCollection")
	assert.True(len(collection.Features) == 2)
	assert.False(len(collection.Features) == 8)

}

func TestSetProperty(t *testing.T) {
	assert := assert.New(t)

	point := NewPoint([]float64{37.6175, 55.752})

	feature := NewFeature()
	feature.SetProperty("location", "Kremlin")
	feature.SetID("0000")
	feature.SetGeometry(point)

	assert.True(feature.Properties["location"] == "Kremlin")
	assert.False(feature.Properties["location"] == "Tverskaya")

}

func TestSetID(t *testing.T) {
	assert := assert.New(t)

	point := NewPoint([]float64{37.6175, 55.752})

	feature := NewFeature()
	feature.SetProperty("location", "Kremlin")
	feature.SetID("0000")
	feature.SetGeometry(point)

	assert.True(feature.ID == "0000")
	assert.False(feature.ID == "8888")
}

func TestSetGeometry(t *testing.T) {
	assert := assert.New(t)

	point := NewPoint([]float64{37.6175, 55.752})

	feature := NewFeature()
	feature.SetProperty("location", "Kremlin")
	feature.SetID("000")
	feature.SetGeometry(point)

	f := (feature.Geom.Coordinates).([]float64)
	assert.True(f[0] == 37.6175)
	assert.True(f[1] == 55.752)

}

func TestGetGeoType(t *testing.T) {
	assert := assert.New(t)
	point := NewPoint([]float64{37.6175, 55.752})

	feature := NewFeature()
	feature.SetProperty("location", "Kremlin")
	feature.SetID("000")
	feature.SetGeometry(point)

	assert.True(feature.Geom.Type == "Point")
	assert.False(feature.Geom.Type == "Polygon")

}

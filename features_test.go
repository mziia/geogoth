package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestNewFeatureCollection(t *testing.T) {
	assert := assert.New(t)

	collection := NewFeatureCollection()

	point1 := NewPoint([]float64{37.6175, 55.752})
	point2 := NewPoint([]float64{37.6048, 55.7649})

	feature1 := NewFeature()
	feature1.SetProperty("локация", "Кремль")
	feature1.SetID("0001")
	feature1.SetGeometry(point1)

	feature2 := NewFeature()
	feature2.SetProperty("локация", "Тверская")
	feature2.SetID("0002")
	feature2.SetGeometry(point2)

	collection.AddFeature(feature1)
	collection.AddFeature(feature2)

	assert.True(collection.Type == "FeatureCollection")
	assert.True(len(collection.Features) == 2)
	assert.False(len(collection.Features) == 8)

	f1 := (feature1.Geom.Coordinates).([]float64)
	assert.True(f1[0] == 37.6175)
	assert.True(f1[1] == 55.752)

}

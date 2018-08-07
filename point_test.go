package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestNewPoint(t *testing.T) {
	assert := assert.New(t)

	point := &Objects.Object1
	CreateObj1()

	assert.True(point.Y == 37.6175)
	assert.True(point.X == 55.752)
}

func TestGetCoordinatesPoint(t *testing.T) {
	assert := assert.New(t)

	point := &Objects.Object1
	CreateObj1()

	coord := (point.GetCoordinates()).([]float64)
	y, x := coord[0], coord[1]

	assert.True(y == 37.6175)
	assert.True(x == 55.752)
}

func TestGetTypePoint(t *testing.T) {
	assert := assert.New(t)

	point := &Objects.Object1
	CreateObj1()

	assert.True(point.GetType() == "Point")
}

func TestGetLengthPoint(t *testing.T) {
	assert := assert.New(t)

	point := &Objects.Object1
	CreateObj1()

	assert.True(point.GetLength() == 0)
	assert.False(point.GetLength() == 37.6175)
}

func TestDistanceToPoint(t *testing.T) {
	assert := assert.New(t)

	point1 := &Objects.Object1
	CreateObj1()

	// Point - Point
	point2 := &Objects.Object2
	CreateObj2()

	assert.True(int(point1.DistanceTo(point2)) == 1639)
	assert.True(int(point2.DistanceTo(point1)) == 1639)

	// Point - MultiPoint
	mpoint := &Objects.Object3
	CreateObj3()

	assert.True(int(point1.DistanceTo(mpoint)) == 1619)

}

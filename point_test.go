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

	point := &Objects.Object1
	CreateObj1()

	// Point - Point
	point2 := &Objects.Object2
	CreateObj2()
	assert.True(int(point.DistanceTo(point2)) == 1639)
	assert.True(int(point2.DistanceTo(point)) == 1639)

	// Point - MultiPoint
	mpoint := &Objects.Object3
	CreateObj3()
	assert.True(int(point.DistanceTo(mpoint)) == 1619)

	// Point - LineString
	linestr := &Objects.Object4
	CreateObj4()
	assert.True(int(point.DistanceTo(linestr)) == 459)

	// Point -MultiLineString
	mlinestr := &Objects.Object5
	CreateObj5()
	assert.True(int(point.DistanceTo(mlinestr)) == 163)

	// Point - Polygon
	polygon := &Objects.Object6
	CreateObj6()
	assert.True(int(point.DistanceTo(polygon)) == 315)

	// Point - MultiPolygon
	mpolyg := &Objects.Object7
	CreateObj7()
	assert.True(int(point.DistanceTo(mpolyg)) == 1445)
}

func TestIntersectsWith(t *testing.T) {
	assert := assert.New(t)

	point := &Objects.Object1
	CreateObj1()

	// Point - Point
	point1 := &Objects.Object1
	CreateObj1()
	point2 := &Objects.Object2
	CreateObj2()

	assert.True(point.IntersectsWith(point1))
	assert.False(point.IntersectsWith(point2))

	// Point - MultiPoint
	mpoint := &Objects.Object20
	CreateObj20()

	assert.True(point1.IntersectsWith(mpoint))
	assert.False(point2.IntersectsWith(mpoint))

	// Point - LineString

	// Point -MultiLineString

	// Point - Polygon

	// Point - MultiPolygon

}

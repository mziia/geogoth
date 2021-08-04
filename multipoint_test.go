package geogoth

import (
	"testing"

	// ToDO
	assert "github.com/blendlabs/go-assert"
)

func TestNewMultiPoint(t *testing.T) {
	assert := assert.New(t)

	mpoint := &Objects.Object3
	CreateObj3()

	assert.True(mpoint.Coords[0][0] == 37.570792444518)
	assert.True(mpoint.Coords[2][1] == 55.77134814581)
}

func TestCoordinatesMPoint(t *testing.T) {
	assert := assert.New(t)

	mpoint := &Objects.Object3
	CreateObj3()

	coords := (mpoint.Coordinates()).([][]float64)

	assert.True(coords[1][0] == 37.624941407105)
	assert.True(coords[1][1] == 55.738048529295)
}

func TestTypeMPoint(t *testing.T) {
	assert := assert.New(t)

	mpoint := &Objects.Object3
	CreateObj3()

	assert.True(mpoint.Type() == "MultiPoint")
}

func TestLengthMPoint(t *testing.T) {
	assert := assert.New(t)

	mpoint := &Objects.Object3
	CreateObj3()

	assert.True(mpoint.Length() == 0)
}

func TestDistanceToMPoint(t *testing.T) {
	assert := assert.New(t)

	mpoint := &Objects.Object3
	CreateObj3()

	// MultiPoint - Point
	point := &Objects.Object1
	CreateObj1()

	assert.True(int(mpoint.DistanceTo(point)) == 1619)

	// MultiPoint - MultiPoint
	mp := &Objects.Object8
	CreateObj8()

	assert.True(int(mpoint.DistanceTo(mp)) == 2541)
	assert.True(int(mp.DistanceTo(mpoint)) == 2541)

	// MultiPoint - LineString
	lstr := &Objects.Object9
	CreateObj9()

	assert.True(int(mpoint.DistanceTo(lstr)) == 2410)

	// MultiPoint - MultiLineString
	mlstr := &Objects.Object5
	CreateObj5()

	assert.True(int(mpoint.DistanceTo(mlstr)) == 736)

	// MultiPoint - Polygon
	pol := &Objects.Object10
	CreateObj10()
	multp := &Objects.Object11
	CreateObj11()

	assert.True(int(multp.DistanceTo(pol)) == 3783)

	// MultiPoint - MultiPolygon
	mpnt := &Objects.Object12
	CreateObj12()
	mpol := &Objects.Object13
	CreateObj13()

	assert.True(int(mpnt.DistanceTo(mpol)) == 61)
}

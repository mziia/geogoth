package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestNewMultiLineString(t *testing.T) {
	assert := assert.New(t)

	mlinestr := &Objects.Object5
	CreateObj5()

	assert.True(mlinestr.Coords[0][1][0] == 37.626585960388184)
	assert.True(mlinestr.Coords[1][2][1] == 55.74665640434626)

}

func TestGetCoordinatesMlinestr(t *testing.T) {
	assert := assert.New(t)

	mlstr := &Objects.Object5
	CreateObj5()

	coords := (mlstr.GetCoordinates()).([][][]float64)

	assert.True(coords[0][0][0] == 37.61911869049072)
	assert.True(coords[1][1][1] == 55.754989002265646)

}

func TestGetTypeMLineStr(t *testing.T) {
	assert := assert.New(t)

	mlinestr := &Objects.Object5
	CreateObj5()

	assert.True(mlinestr.GetType() == "MultiLineString")
}

func TestLengthMLineStr(t *testing.T) {
	assert := assert.New(t)

	multLine := NewMultiLineString([][][]float64{
		[][]float64{[]float64{37.61911869049072, 55.75634137200424},
			[]float64{37.626585960388184, 55.75218751578049},
			[]float64{37.62401103973388, 55.74465148474084}},
		[][]float64{[]float64{37.614097595214844, 55.7528396038768},
			[]float64{37.62409687042236, 55.754989002265646},
			[]float64{37.6255989074707, 55.74665640434626},
			[]float64{37.62989044189453, 55.74503798328292}}})

	assert.True(multLine.GetLength() == 3434.744143073227)

}

func TestDistanceTo(t *testing.T) {
	assert := assert.New(t)

	mlstr := &Objects.Object5
	CreateObj5()

	// MultiLineString - Point
	point := &Objects.Object1
	CreateObj1()
	assert.True(int(mlstr.DistanceTo(point)) == 163)

}

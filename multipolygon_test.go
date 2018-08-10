package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestNewMultiPolygon(t *testing.T) {
	assert := assert.New(t)

	mpolyg := &Objects.Object7
	CreateObj7()

	assert.True(mpolyg.Coords[0][0][2][0] == 37.596588134765625)
	assert.True(mpolyg.Coords[0][0][4][1] == 55.76189525593947)
	assert.True(mpolyg.Coords[1][0][0][0] == 38.08904439210892)
	assert.True(mpolyg.Coords[1][0][3][1] == 55.81212194465421)

}

func TestGetCoordinatesMPolyg(t *testing.T) {
	assert := assert.New(t)

	mpolyg := &Objects.Object7
	CreateObj7()

	coords := (mpolyg.GetCoordinates()).([][][][]float64)

	assert.True(coords[0][0][2][0] == 37.596588134765625)
	assert.True(coords[0][0][4][1] == 55.76189525593947)
	assert.True(coords[1][0][0][0] == 38.08904439210892)
	assert.True(coords[1][0][3][1] == 55.81212194465421)

}

func TestGetTypeMPolyg(t *testing.T) {
	assert := assert.New(t)

	mpolyg := &Objects.Object7
	CreateObj7()

	assert.True(mpolyg.GetType() == "MultiPolygon")
}

func TestGetLengthMPolygon(t *testing.T) {
	assert := assert.New(t)

	mpolyg := NewMultiPolygon([][][][]float64{

		[][][]float64{
			[][]float64{

				[]float64{37.57770538330078, 55.76691721773862},
				[]float64{37.57942199707031, 55.7659515060885},
				[]float64{37.63298034667969, 55.73406960018026},
				[]float64{37.644996643066406, 55.76884856927786},
				[]float64{37.57770538330078, 55.76691721773862},
			}},

		[][][]float64{
			[][]float64{

				[]float64{37.57427215576172, 55.74856460562653},
				[]float64{37.610321044921875, 55.73928842238313},
				[]float64{37.64568328857422, 55.78410288303946},
				[]float64{37.60414123535156, 55.78062774182665},
				[]float64{37.57427215576172, 55.74856460562653}}}})

	assert.True(mpolyg.GetLength() == 27769.448354979733)

}

func TestDistanceToMPolyg(t *testing.T) {
	assert := assert.New(t)

	mpolyg := &Objects.Object7
	CreateObj7()

	// MultiPolygon - Point
	point := &Objects.Object1
	CreateObj1()
	assert.True(int(mpolyg.DistanceTo(point)) == 1445)

	// MultiPolygon - MultiPoint
	mpoint := &Objects.Object12
	CreateObj12()
	mpol := &Objects.Object13
	CreateObj13()

	assert.True(int(mpol.DistanceTo(mpoint)) == 61)

	// MultiPolygon - LineString
	ln := &Objects.Object16
	CreateObj16()

	assert.True(int(mpol.DistanceTo(ln)) == 1000)

	// MultiPolygon - MultiLineString
	mpl := &Objects.Object7
	CreateObj7()
	mlstr := &Objects.Object18
	CreateObj18()

	assert.True(int(mpl.DistanceTo(mlstr)) == 3444)

}

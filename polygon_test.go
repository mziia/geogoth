package geogoth

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestNewPolygon(t *testing.T) {
	assert := assert.New(t)

	polyg := &Objects.Object6
	CreateObj6()

	assert.True(polyg.Coords[0][0][0] == 37.60027885437012)
	assert.True(polyg.Coords[0][3][1] == 55.74986889204201)

}

func TestGetCoordinatesPolyg(t *testing.T) {
	assert := assert.New(t)

	polyg := &Objects.Object6
	CreateObj6()

	coords := (polyg.GetCoordinates()).([][][]float64)

	assert.True(coords[0][4][0] == 37.61246681213379)
	assert.True(coords[0][4][1] == 55.75209090920951)
}

func TestGetTypePolyg(t *testing.T) {
	assert := assert.New(t)

	polyg := &Objects.Object6
	CreateObj6()

	assert.True(polyg.GetType() == "Polygon")
}

func TestLength(t *testing.T) {
	assert := assert.New(t)

	polyg := NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.470245361328125, 55.80706963076952},
			[]float64{37.63023376464844, 55.62993418221071},
			[]float64{37.74696350097656, 55.819801652442436},
			[]float64{37.470245361328125, 55.80706963076952}},
		[][]float64{
			[]float64{37.53822326660156, 55.76884856927786},
			[]float64{37.62062072753906, 55.73870858770892},
			[]float64{37.651519775390625, 55.792017325495095},
			[]float64{37.57598876953125, 55.79742138660978},
			[]float64{37.53822326660156, 55.76884856927786}}})

	assert.True(polyg.GetLength() == 82890.25104294329)
}

func TestDistanceToPolygon(t *testing.T) {
	assert := assert.New(t)

	polyg := &Objects.Object6
	CreateObj6()

	// Polygon - Point
	point := &Objects.Object1
	CreateObj1()

	assert.True(int(polyg.DistanceTo(point)) == 315)

	// Polygon - MultiPoint
	pol := &Objects.Object10
	CreateObj10()
	mpoint := &Objects.Object11
	CreateObj11()

	assert.True(int(pol.DistanceTo(mpoint)) == 3783)

	// Polygon - LineString
	lst := &Objects.Object14
	CreateObj14()

	assert.True(int(polyg.DistanceTo(lst)) == 83203)

	// Polygon - MultiLineString
	mlstr := &Objects.Object18
	CreateObj18()

	assert.True(int(mlstr.DistanceTo(polyg)) == 3638)

	// Polygon - Polygon
	p := &Objects.Object19
	CreateObj19()

	assert.True(int(p.DistanceTo(pol)) == 7401)
	assert.True(int(pol.DistanceTo(p)) == 7401)

}

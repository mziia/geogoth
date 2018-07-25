package geogoth

import (
	"testing"

	"github.com/blendlabs/go-assert"
)

func TestDistancePointPointDeg(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr1()
	CreateFtr2()

	y1, x1 := GetPointCoordinates(ftr.Feature1) // Coordinates of Point
	y2, x2 := GetPointCoordinates(ftr.Feature2)

	assert.True(int(DistancePointPointDeg(y1, x1, y2, x2)) == 1639)

}

// func TestDistancePointPointRad(t *testing.T) {
// 	assert := assert.New(t)
// }

func TestDistancePointMultipoint(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr1()
	CreateFtr3()

	assert.True(int(DistancePointMultipoint(ftr.Feature1, ftr.Feature3)) == 1619)
}

func TestDistancePointLine(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr1()
	CreateFtr4()

	py, px := GetPointCoordinates(ftr.Feature1)
	ly1, lx1 := GetTwoDimArrayCoordinates(ftr.Feature4, 0)
	ly2, lx2 := GetTwoDimArrayCoordinates(ftr.Feature4, 1)

	assert.True(int(DistancePointLine(py, px, ly1, lx1, ly2, lx2)) == 459)
}

func TestDistancePointLinstring(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr1()
	CreateFtr4()

	assert.True(int(DistancePointLinstring(ftr.Feature1, ftr.Feature4)) == 459)
}

func TestDistancePointMultiLinstring(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr1()
	CreateFtr5()

	assert.True(int(DistancePointMultiLineString(ftr.Feature1, ftr.Feature5)) == 163)
}

func TestDistancePointPolygon(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr1()
	CreateFtr6()

	assert.True(int(DistancePointPolygon(ftr.Feature1, ftr.Feature6)) == 315)
}

func TestDistancePointMultiPolygon(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr1()
	CreateFtr7()

	assert.True(int(DistancePointMultiPolygon(ftr.Feature1, ftr.Feature7)) == 1445)
}

func TestDistanceMultipointMultipoint(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr3()
	CreateFtr8()

	assert.True(int(DistanceMultipointMultipoint(ftr.Feature3, ftr.Feature8)) == 2541)
}

func TestDistanceMultipointLineString(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr3()
	CreateFtr9()

	assert.True(int(DistanceMultipointLinestring(ftr.Feature3, ftr.Feature9)) == 2410)
}

func TestDistanceMultipointMultiLinestring(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr3()
	CreateFtr5()

	assert.True(int(DistanceMultiPointMultiLinestring(ftr.Feature3, ftr.Feature5)) == 736)
}

func TestDistanceMultiPointPolygon(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr10()
	CreateFtr11()

	assert.True(int(DistanceMultiPointPolygon(ftr.Feature11, ftr.Feature10)) == 3783)
}

func TestDistanceMultiPointMultiPolygon(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr12()
	CreateFtr13()

	assert.True(int(DistanceMultiPointMultiPolygon(ftr.Feature12, ftr.Feature13)) == 61)
}

func TestDistanceLineStringLineString(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr14()
	CreateFtr15()

	assert.True(int(DistanceLineStringLineString(ftr.Feature15, ftr.Feature14)) == 144703)
}

func TestDistanceLineStringMultiLineString(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr14()
	CreateFtr5()

	assert.True(int(DistanceLineStringMultiLineString(ftr.Feature14, ftr.Feature5)) == 84091)
}

func TestDistanceLineStringPolygon(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr14()
	CreateFtr6()

	assert.True(int(DistanceLineStringPolygon(ftr.Feature14, ftr.Feature6)) == 83203)
}

func TestDistanceLineStringMultiPolygon(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr13()
	CreateFtr16()

	assert.True(int(DistanceLineStringMultiPolygon(ftr.Feature16, ftr.Feature13)) == 1000)
}

func TestDistanceMultiLineStringMultiLineString(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr17()
	CreateFtr18()

	assert.True(int(DistanceMultiLineStringMultiLineString(ftr.Feature17, ftr.Feature18)) == 3005)
}

func TestDistanceMultiLineStringPolygon(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr6()
	CreateFtr18()

	assert.True(int(DistanceMultiLineStringPolygon(ftr.Feature6, ftr.Feature18)) == 3638)
}

func TestDistanceMultiLineStringMultiPolygon(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr7()
	CreateFtr18()

	assert.True(int(DistanceMultiLineStringMultiPolygon(ftr.Feature18, ftr.Feature7)) == 3444)
}

func TestDistancePolygonPolygon(t *testing.T) {
	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr10()
	CreateFtr19()

	assert.True(int(DistancePolygonPolygon(ftr.Feature10, ftr.Feature19)) == 7401)
}

// func TestDistancePolygonMultiPolygon(t *testing.T) {
// 	assert := assert.New(t)

// ftr := &Ftrs
// CreateFtr()
// CreateFtr()

// assert.True(int((ftr.Feature, ftr.Feature)) ==  )
// }

// func TestDistanceMultiPolygonMultiPolygon(t *testing.T) {
// 	assert := assert.New(t)

// ftr := &Ftrs
// CreateFtr()
// CreateFtr()

// assert.True(int((ftr.Feature, ftr.Feature)) ==  )
// }

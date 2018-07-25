package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestDistance(t *testing.T) {

	assert := assert.New(t)
	ftr := &Ftrs

	CreateFtr1()
	CreateFtr2()
	CreateFtr3()
	CreateFtr4()
	CreateFtr5()
	CreateFtr6()
	CreateFtr7()
	CreateFtr8()
	CreateFtr9()
	CreateFtr10()
	CreateFtr11()
	CreateFtr12()
	CreateFtr13()
	CreateFtr14()
	CreateFtr15()
	CreateFtr16()
	CreateFtr17()
	CreateFtr18()
	CreateFtr19()
	CreateFtr20()
	CreateFtr21()
	// *****
	// *****
	// *****

	// Point - Point
	assert.True(Distance(ftr.Feature1, ftr.Feature2) == Distance(ftr.Feature2, ftr.Feature1))
	assert.True(int(Distance(ftr.Feature1, ftr.Feature2)) == 1639)
	// Point - MultiPoint
	assert.True(Distance(ftr.Feature1, ftr.Feature3) == Distance(ftr.Feature3, ftr.Feature1))
	assert.True(int(Distance(ftr.Feature1, ftr.Feature3)) == 1619)
	// Point - LineString
	assert.True(Distance(ftr.Feature1, ftr.Feature4) == Distance(ftr.Feature4, ftr.Feature1))
	assert.True(int(Distance(ftr.Feature1, ftr.Feature4)) == 459)
	// Point - MultiLineString
	assert.True(Distance(ftr.Feature1, ftr.Feature5) == Distance(ftr.Feature5, ftr.Feature1))
	assert.True(int(Distance(ftr.Feature1, ftr.Feature5)) == 163)
	// Point - Polygon
	assert.True(Distance(ftr.Feature1, ftr.Feature6) == Distance(ftr.Feature6, ftr.Feature1))
	assert.True(int(Distance(ftr.Feature1, ftr.Feature6)) == 315)
	// Point - MultiPolygon
	assert.True(Distance(ftr.Feature1, ftr.Feature7) == Distance(ftr.Feature7, ftr.Feature1))
	assert.True(int(Distance(ftr.Feature1, ftr.Feature7)) == 1445)

	// MultiPoint - MultiPoint
	assert.True(Distance(ftr.Feature3, ftr.Feature8) == Distance(ftr.Feature8, ftr.Feature3))
	assert.True(int(Distance(ftr.Feature3, ftr.Feature8)) == 2541)
	// MultiPoint - LineString
	assert.True(Distance(ftr.Feature3, ftr.Feature9) == Distance(ftr.Feature9, ftr.Feature3))
	assert.True(int(Distance(ftr.Feature3, ftr.Feature9)) == 2410)
	// MultiPoint - MultiLineString
	assert.True(Distance(ftr.Feature3, ftr.Feature5) == Distance(ftr.Feature5, ftr.Feature3))
	assert.True(int(Distance(ftr.Feature3, ftr.Feature5)) == 736)
	// MultiPoint - Polygon
	assert.True(Distance(ftr.Feature10, ftr.Feature11) == Distance(ftr.Feature11, ftr.Feature10))
	assert.True(int(Distance(ftr.Feature10, ftr.Feature11)) == 3783)
	// MultiPoint - MultiPolygon
	assert.True(Distance(ftr.Feature12, ftr.Feature13) == Distance(ftr.Feature13, ftr.Feature12))
	assert.True(int(Distance(ftr.Feature12, ftr.Feature13)) == 61)

	// LineString - LineString
	assert.True(Distance(ftr.Feature14, ftr.Feature15) == Distance(ftr.Feature15, ftr.Feature14))
	assert.True(int(Distance(ftr.Feature14, ftr.Feature15)) == 144703)

	// LineString - MultiLineString
	assert.True(Distance(ftr.Feature14, ftr.Feature5) == Distance(ftr.Feature5, ftr.Feature14))
	assert.True(int(Distance(ftr.Feature14, ftr.Feature5)) == 84091)
	// LineString - Polygon
	assert.True(Distance(ftr.Feature14, ftr.Feature6) == Distance(ftr.Feature6, ftr.Feature14))
	assert.True(int(Distance(ftr.Feature14, ftr.Feature6)) == 83203)
	// LineString - MultiPolygon
	assert.True(Distance(ftr.Feature13, ftr.Feature16) == Distance(ftr.Feature16, ftr.Feature13))
	assert.True(int(Distance(ftr.Feature13, ftr.Feature16)) == 1000)

	// MultiLineString - MultiLineString
	assert.True(Distance(ftr.Feature17, ftr.Feature18) == Distance(ftr.Feature18, ftr.Feature17))
	assert.True(int(Distance(ftr.Feature17, ftr.Feature18)) == 3005)
	// MultiLineString - Polygon
	assert.True(Distance(ftr.Feature6, ftr.Feature18) == Distance(ftr.Feature18, ftr.Feature6))
	assert.True(int(Distance(ftr.Feature6, ftr.Feature18)) == 3638)
	// MultiLineString - MultiPolygon
	assert.True(Distance(ftr.Feature7, ftr.Feature18) == Distance(ftr.Feature18, ftr.Feature7))
	assert.True(int(Distance(ftr.Feature7, ftr.Feature18)) == 3444)

	// Polygon - Polygon
	assert.True(Distance(ftr.Feature10, ftr.Feature19) == Distance(ftr.Feature19, ftr.Feature10))
	assert.True(int(Distance(ftr.Feature10, ftr.Feature19)) == 7401)
	// Polygon - MultiPolygon

	// MultiPolygon - MultiPolygon

}

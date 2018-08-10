package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestNewLineString(t *testing.T) {
	assert := assert.New(t)

	lstring := &Objects.Object4
	CreateObj4()

	assert.True(lstring.Coords[0][0] == 37.624941407105)
	assert.True(lstring.Coords[1][1] == 55.784581739175)

}

func TestGetCoordinateslStr(t *testing.T) {
	assert := assert.New(t)

	lstr := &Objects.Object4
	CreateObj4()

	coords := (lstr.GetCoordinates()).([][]float64)

	assert.True(coords[1][0] == 37.570792444518)
	assert.True(coords[1][1] == 55.784581739175)
}

func TestGetTypeLineStr(t *testing.T) {
	assert := assert.New(t)

	linestr := &Objects.Object4
	CreateObj4()

	assert.True(linestr.GetType() == "LineString")
}

func TestNGetLengthLineString(t *testing.T) {
	assert := assert.New(t)

	linestr := NewLineString([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.596588134765625, 55.72169627105833},
		[]float64{37.595901489257805, 55.7114466394498},
		[]float64{37.59727478027344, 55.70796502063464},
		[]float64{37.61272430419922, 55.70100085220915},
	})
	assert.True(linestr.GetLength() == 8023.4822342726075)

}

func TestDistanceToLineStr(t *testing.T) {
	assert := assert.New(t)

	linestr := &Objects.Object4
	CreateObj4()

	// LineString - Point
	point := &Objects.Object1
	CreateObj1()

	assert.True(int(linestr.DistanceTo(point)) == 459)

	// LineString - MultiPoint
	lstr := &Objects.Object9
	CreateObj9()
	mpoint := &Objects.Object3
	CreateObj3()

	assert.True(int(lstr.DistanceTo(mpoint)) == 2410)

	// LineString - LineString
	lst1 := &Objects.Object14
	CreateObj14()
	lst2 := &Objects.Object15
	CreateObj15()

	assert.True(int(lst1.DistanceTo(lst2)) == 144703)
	assert.True(int(lst2.DistanceTo(lst1)) == 144703)

	// LineString - MultiLineString
	mlinestr := &Objects.Object5
	CreateObj5()

	assert.True(int(lst1.DistanceTo(mlinestr)) == 84091)

	// LineString - Polygon
	pol := &Objects.Object6
	CreateObj6()

	assert.True(int(lst1.DistanceTo(pol)) == 83203)

	// LineString - MultiPolygon
	mpol := &Objects.Object13
	CreateObj13()
	ln := &Objects.Object16
	CreateObj16()

	assert.True(int(ln.DistanceTo(mpol)) == 1000)

}

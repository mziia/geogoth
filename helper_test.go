package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestMinDistance(t *testing.T) {

	assert := assert.New(t)

	assert.True(MinDistance([]float64{256, 64, 1024, 8}) == 8)
}

func TestNegToPosRad(t *testing.T) {

	assert := assert.New(t)

	assert.True(NegToPosRad(-0.26954871580518647) == 6.0136365913744)
}

func TestBearing(t *testing.T) {

	assert := assert.New(t)

	assert.True(Bearing(0.9733938460598921, 0.6565617535524124, 0.9745994309845489, 0.6559685344195972) == 6.0136365913744)
}

func TestPIPJordanCurveTheorem(t *testing.T) {

	assert := assert.New(t)

	plg := NewPolygon([][][]float64{
		[][]float64{

			[]float64{37.568435668945305, 55.75899767604707},
			[]float64{37.49702453613281, 55.71028613431717},
			[]float64{37.53547668457031, 55.68590756259903},
			[]float64{37.57598876953125, 55.65667193958129},
			[]float64{37.68413543701172, 55.68048803720246},
			[]float64{37.691001892089844, 55.71086639119248},
			[]float64{37.68516540527344, 55.75320187033113},
			[]float64{37.65907287597656, 55.7624747460818},
			[]float64{37.57392883300781, 55.76846230662442},
			[]float64{37.51213073730469, 55.761315757185166},
			[]float64{37.495994567871094, 55.741414409152696},
			[]float64{37.53856658935547, 55.70931902037369},
			[]float64{37.58628845214844, 55.68784292538311},
			[]float64{37.64019012451172, 55.696357383733535},
			[]float64{37.661476135253906, 55.73638916286441},
			[]float64{37.63160705566406, 55.74972397348254},
			[]float64{37.562599182128906, 55.73116995298137},
			[]float64{37.55950927734375, 55.70815845204154},
			[]float64{37.61238098144531, 55.696163893908825},
			[]float64{37.64362335205078, 55.72420992089696},
			[]float64{37.620277404785156, 55.73213652597946},
			[]float64{37.605857849121094, 55.74798490877827},
			[]float64{37.58766174316406, 55.72150290667663},
			[]float64{37.61100769042969, 55.71860232605969},
			[]float64{37.568435668945305, 55.75899767604707}}})

	feature := NewFeature()
	feature.SetProperty("Polyg", "[][][]")
	feature.SetID("000")
	feature.SetGeometry(plg)

	point000 := NewPoint([]float64{37.6728, 55.7772})
	feature000 := NewFeature()
	feature000.SetProperty("location", "Kremlin")
	feature000.SetID("000000")
	feature000.SetGeometry(point000)

	y, x := GetPointCoordinates(feature000)
	pol := feature.Geom.Coordinates
	assert.True(PIPJordanCurveTheorem(y, x, pol) == false)

	point001 := NewPoint([]float64{37.5945, 55.7509})
	feature001 := NewFeature()
	feature001.SetProperty("location", "Kremlin")
	feature001.SetID("000001")
	feature001.SetGeometry(point001)

	y, x = GetPointCoordinates(feature001)
	pol = feature.Geom.Coordinates
	assert.True(PointInPolygon(feature001, feature) == true)

}

func TestPointInPolygon(t *testing.T) {

	assert := assert.New(t)

	plg := NewPolygon([][][]float64{
		[][]float64{

			[]float64{37.568435668945305, 55.75899767604707},
			[]float64{37.49702453613281, 55.71028613431717},
			[]float64{37.53547668457031, 55.68590756259903},
			[]float64{37.57598876953125, 55.65667193958129},
			[]float64{37.68413543701172, 55.68048803720246},
			[]float64{37.691001892089844, 55.71086639119248},
			[]float64{37.68516540527344, 55.75320187033113},
			[]float64{37.65907287597656, 55.7624747460818},
			[]float64{37.57392883300781, 55.76846230662442},
			[]float64{37.51213073730469, 55.761315757185166},
			[]float64{37.495994567871094, 55.741414409152696},
			[]float64{37.53856658935547, 55.70931902037369},
			[]float64{37.58628845214844, 55.68784292538311},
			[]float64{37.64019012451172, 55.696357383733535},
			[]float64{37.661476135253906, 55.73638916286441},
			[]float64{37.63160705566406, 55.74972397348254},
			[]float64{37.562599182128906, 55.73116995298137},
			[]float64{37.55950927734375, 55.70815845204154},
			[]float64{37.61238098144531, 55.696163893908825},
			[]float64{37.64362335205078, 55.72420992089696},
			[]float64{37.620277404785156, 55.73213652597946},
			[]float64{37.605857849121094, 55.74798490877827},
			[]float64{37.58766174316406, 55.72150290667663},
			[]float64{37.61100769042969, 55.71860232605969},
			[]float64{37.568435668945305, 55.75899767604707}}})

	feature := NewFeature()
	feature.SetProperty("Polyg", "[][][]")
	feature.SetID("000")
	feature.SetGeometry(plg)

	point000 := NewPoint([]float64{37.6728, 55.7772})
	feature000 := NewFeature()
	feature000.SetProperty("location", "Kremlin")
	feature000.SetID("000000")
	feature000.SetGeometry(point000)

	assert.True(PointInPolygon(feature000, feature) == false)

	point001 := NewPoint([]float64{37.5945, 55.7509})
	feature001 := NewFeature()
	feature001.SetProperty("location", "Kremlin")
	feature001.SetID("000001")
	feature001.SetGeometry(point001)

	assert.True(PointInPolygon(feature001, feature) == true)

	point002 := NewPoint([]float64{37.6162, 55.7476})
	feature002 := NewFeature()
	feature002.SetProperty("location", "Kremlin")
	feature002.SetID("00015_00")
	feature002.SetGeometry(point002)

	assert.True(PointInPolygon(feature002, feature) == true)

	point003 := NewPoint([]float64{37.6062, 55.7449})
	feature003 := NewFeature()
	feature003.SetProperty("location", "Kremlin")
	feature003.SetID("000003")
	feature003.SetGeometry(point003)

	assert.True(PointInPolygon(feature003, feature) == false)

	point004 := NewPoint([]float64{37.5945, 55.7383})
	feature004 := NewFeature()
	feature004.SetProperty("location", "Kremlin")
	feature004.SetID("000004")
	feature004.SetGeometry(point004)

	assert.True(PointInPolygon(feature004, feature) == false)

	point005 := NewPoint([]float64{37.6083, 55.7362})
	feature005 := NewFeature()
	feature005.SetProperty("location", "Kremlin")
	feature005.SetID("000005")
	feature005.SetGeometry(point005)

	assert.True(PointInPolygon(feature005, feature) == true)

	point006 := NewPoint([]float64{37.5983, 55.7250})
	feature006 := NewFeature()
	feature006.SetProperty("location", "Kremlin")
	feature006.SetID("000006")
	feature006.SetGeometry(point006)

	assert.True(PointInPolygon(feature006, feature) == false)

	point007 := NewPoint([]float64{37.5702, 55.7203})
	feature007 := NewFeature()
	feature007.SetProperty("location", "Kremlin")
	feature007.SetID("000007")
	feature007.SetGeometry(point007)

	assert.True(PointInPolygon(feature007, feature) == true)

	point008 := NewPoint([]float64{37.5671, 55.7426})
	feature008 := NewFeature()
	feature008.SetProperty("location", "Kremlin")
	feature008.SetID("000008")
	feature008.SetGeometry(point008)

	assert.True(PointInPolygon(feature008, feature) == false)

	point009 := NewPoint([]float64{37.5393, 55.7163})
	feature009 := NewFeature()
	feature009.SetProperty("location", "Kremlin")
	feature009.SetID("000009")
	feature009.SetGeometry(point009)

	assert.True(PointInPolygon(feature009, feature) == false)

	point010 := NewPoint([]float64{37.5866, 55.6944})
	feature010 := NewFeature()
	feature010.SetProperty("location", "Kremlin")
	feature010.SetID("000010")
	feature010.SetGeometry(point010)

	assert.True(PointInPolygon(feature010, feature) == false)

	point011 := NewPoint([]float64{37.6141, 55.6936})
	feature011 := NewFeature()
	feature011.SetProperty("location", "Kremlin")
	feature011.SetID("000011")
	feature011.SetGeometry(point011)

	assert.True(PointInPolygon(feature011, feature) == false)

	point012 := NewPoint([]float64{37.6433, 55.7151})
	feature012 := NewFeature()
	feature012.SetProperty("location", "Kremlin")
	feature012.SetID("000012")
	feature012.SetGeometry(point012)

	assert.True(PointInPolygon(feature012, feature) == false)

	point013 := NewPoint([]float64{37.6440, 55.7346})
	feature013 := NewFeature()
	feature013.SetProperty("location", "Kremlin")
	feature013.SetID("000013")
	feature013.SetGeometry(point013)

	assert.True(PointInPolygon(feature013, feature) == false)

	point014 := NewPoint([]float64{37.5211, 55.7352})
	feature014 := NewFeature()
	feature014.SetProperty("location", "Kremlin")
	feature014.SetID("000014")
	feature014.SetGeometry(point014)

	assert.True(PointInPolygon(feature014, feature) == true)

	point015 := NewPoint([]float64{37.6217, 55.6789})
	feature015 := NewFeature()
	feature015.SetProperty("location", "Kremlin")
	feature015.SetID("000015")
	feature015.SetGeometry(point015)

	assert.True(PointInPolygon(feature015, feature) == true)

	point016 := NewPoint([]float64{37.5073, 55.7236})
	feature016 := NewFeature()
	feature016.SetProperty("location", "Kremlin")
	feature016.SetID("000016")
	feature016.SetGeometry(point016)

	assert.True(PointInPolygon(feature016, feature) == false)

	point017 := NewPoint([]float64{37.5132, 55.6549})
	feature017 := NewFeature()
	feature017.SetProperty("location", "Kremlin")
	feature017.SetID("000017")
	feature017.SetGeometry(point017)

	assert.True(PointInPolygon(feature017, feature) == false)

}

func TestLineLineIntersection(t *testing.T) {

	assert := assert.New(t)

	assert.True(LineLineIntersection(37.62036323547363, 55.75223581897627, 37.625813484191895, 55.751366352300806, 37.62409687042236, 55.754989002265646, 37.6255989074707, 55.74665640434626) == true)
	assert.True(LineLineIntersection(37.62036323547363, 55.75223581897627, 37.625813484191895, 55.751366352300806, 37.626585960388184, 55.75218751578049, 37.62401103973388, 55.74465148474084) == false)

}

func TestDistanceLineLine(t *testing.T) {

	assert := assert.New(t)

	assert.True(DistanceLineLine(37.58422851562499, 55.840434111266205, 37.513160705566406, 55.82250184886082, 37.677955627441406, 55.82404473410693, 37.63298034667969, 55.789121984291626) == 5810.770008570808)

}

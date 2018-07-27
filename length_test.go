package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestLength(t *testing.T) {

	assert := assert.New(t)

	ftr := &Ftrs
	CreateFtr1() // Point
	assert.True(ftr.Feature1.Length() == 0)

	CreateFtr3() // MultiPoint
	assert.True(ftr.Feature3.Length() == 0)

	linestr := NewLineString([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.596588134765625, 55.72169627105833},
		[]float64{37.595901489257805, 55.7114466394498},
		[]float64{37.59727478027344, 55.70796502063464},
		[]float64{37.61272430419922, 55.70100085220915},
	})

	feature := NewFeature()
	feature.SetProperty("LineStr", "[][]")
	feature.SetID("ftr")
	feature.SetGeometry(linestr)

	assert.True(feature.Length() == 8023.4822342726075)

	multLine := NewMultiLineString([][][]float64{
		[][]float64{[]float64{37.61911869049072, 55.75634137200424},
			[]float64{37.626585960388184, 55.75218751578049},
			[]float64{37.62401103973388, 55.74465148474084}},
		[][]float64{[]float64{37.614097595214844, 55.7528396038768},
			[]float64{37.62409687042236, 55.754989002265646},
			[]float64{37.6255989074707, 55.74665640434626},
			[]float64{37.62989044189453, 55.74503798328292}}})

	feature1 := NewFeature()
	feature1.SetProperty("MLineStr", "[][][]")
	feature1.SetID("ftr")
	feature1.SetGeometry(multLine)

	assert.True(feature1.Length() == 3434.744143073227)

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

	feature2 := NewFeature()
	feature2.SetProperty("Polyg", "[][][]")
	feature2.SetID("ftr")
	feature2.SetGeometry(polyg)

	assert.True(feature2.Length() == 82890.25104294329)

	multipol := NewMultiPolygon([][][][]float64{

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

	feature3 := NewFeature()
	feature3.SetProperty("Polyg", "[][][]")
	feature3.SetID("ftr")
	feature3.SetGeometry(multipol)

	assert.True(feature3.Length() == 27769.448354979733)

}

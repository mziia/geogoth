package geogoth_test

import (
	geogoth "geogoth/geojson"
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestLineStringLength(t *testing.T) {
	assert := assert.New(t)

	linestr := geogoth.NewLineString([][]float64{
		{37.566375732421875, 55.761702090644896},
		{37.58800506591796, 55.74856460562653},
		{37.58491516113281, 55.72981671057788},
		{37.596588134765625, 55.72169627105833},
		{37.595901489257805, 55.7114466394498},
		{37.59727478027344, 55.70796502063464},
		{37.61272430419922, 55.70100085220915},
	})

	feature := geogoth.NewFeature()
	feature.SetProperty("LineStr", "[][]")
	feature.SetID("ftr")
	feature.SetGeometry(linestr)

	lengths := geogoth.LineStringLength(feature)
	length := feature.Length()

	assert.True(lengths == length)
	assert.True(length == 8023.4822342726075)

}

func TestMultiLineStringLength(t *testing.T) {
	assert := assert.New(t)

	multLine := geogoth.NewMultiLineString([][][]float64{
		{{37.61911869049072, 55.75634137200424},
			{37.626585960388184, 55.75218751578049},
			{37.62401103973388, 55.74465148474084}},
		{{37.614097595214844, 55.7528396038768},
			{37.62409687042236, 55.754989002265646},
			{37.6255989074707, 55.74665640434626},
			{37.62989044189453, 55.74503798328292}}})

	feature := geogoth.NewFeature()
	feature.SetProperty("MLineStr", "[][][]")
	feature.SetID("ftr")
	feature.SetGeometry(multLine)

	lengths := geogoth.MultiLineStringLength(feature)
	length := feature.Length()

	assert.True(lengths == length)
	assert.True(length == 3434.744143073227)

}

func TestPolygonLength(t *testing.T) {
	assert := assert.New(t)

	polyg := geogoth.NewPolygon([][][]float64{
		{
			{37.470245361328125, 55.80706963076952},
			{37.63023376464844, 55.62993418221071},
			{37.74696350097656, 55.819801652442436},
			{37.470245361328125, 55.80706963076952}},
		{
			{37.53822326660156, 55.76884856927786},
			{37.62062072753906, 55.73870858770892},
			{37.651519775390625, 55.792017325495095},
			{37.57598876953125, 55.79742138660978},
			{37.53822326660156, 55.76884856927786}}})

	feature := geogoth.NewFeature()
	feature.SetProperty("Polyg", "[][][]")
	feature.SetID("ftr")
	feature.SetGeometry(polyg)

	lengths := geogoth.PolygonLength(feature)
	length := feature.Length()

	assert.True(lengths == length)
	assert.True(length == 82890.25104294329)

}

func TestMultipolygonLength(t *testing.T) {
	assert := assert.New(t)

	multipol := geogoth.NewMultiPolygon([][][][]float64{

		{
			{

				{37.57770538330078, 55.76691721773862},
				{37.57942199707031, 55.7659515060885},
				{37.63298034667969, 55.73406960018026},
				{37.644996643066406, 55.76884856927786},
				{37.57770538330078, 55.76691721773862},
			}},

		{
			{

				{37.57427215576172, 55.74856460562653},
				{37.610321044921875, 55.73928842238313},
				{37.64568328857422, 55.78410288303946},
				{37.60414123535156, 55.78062774182665},
				{37.57427215576172, 55.74856460562653}}}})

	feature := geogoth.NewFeature()
	feature.SetProperty("Polyg", "[][][]")
	feature.SetID("ftr")
	feature.SetGeometry(multipol)

	lengths := geogoth.MultipolygonLength(feature)
	length := feature.Length()

	assert.True(lengths == length)
	assert.True(length == 27769.448354979733)
}

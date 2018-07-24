package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestDistance(t *testing.T) {

	assert := assert.New(t)

	point1 := NewPoint([]float64{37.6175, 55.752})

	feature1 := NewFeature()
	feature1.SetProperty("1", "1")
	feature1.SetID("0001")
	feature1.SetGeometry(point1)
	// *****

	point2 := NewPoint([]float64{37.6048, 55.7649})

	feature2 := NewFeature()
	feature2.SetProperty("2", "2")
	feature2.SetID("0002")
	feature2.SetGeometry(point2)
	// *****

	mult := NewMultiPoint([][]float64{[]float64{37.570792444518, 55.784581739175},
		[]float64{37.624941407105, 55.738048529295}, []float64{37.633109750449, 55.77134814581}})

	feature3 := NewFeature()
	feature3.SetProperty("3", "3")
	feature3.SetID("0003")
	feature3.SetGeometry(mult)
	// *****

	linestr := NewLineString([][]float64{

		[]float64{37.624941407105, 55.738048529295},
		[]float64{37.570792444518, 55.784581739175}})

	feature4 := NewFeature()
	feature4.SetProperty("4", "4")
	feature4.SetID("0004")
	feature4.SetGeometry(linestr)
	// *****

	multLine := NewMultiLineString([][][]float64{
		[][]float64{[]float64{37.61911869049072, 55.75634137200424},
			[]float64{37.626585960388184, 55.75218751578049},
			[]float64{37.62401103973388, 55.74465148474084}},
		[][]float64{[]float64{37.614097595214844, 55.7528396038768},
			[]float64{37.62409687042236, 55.754989002265646},
			[]float64{37.6255989074707, 55.74665640434626},
			[]float64{37.62989044189453, 55.74503798328292}}})

	feature5 := NewFeature()
	feature5.SetProperty("5", "5")
	feature5.SetID("0005")
	feature5.SetGeometry(multLine)
	// *****

	pol := NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.60027885437012, 55.75573764130792},
			[]float64{37.60027885437012, 55.75556859503827},
			[]float64{37.59929180145264, 55.751632024729965},
			[]float64{37.60671615600586, 55.74986889204201},
			[]float64{37.61246681213379, 55.75209090920951},
			[]float64{37.60027885437012, 55.75573764130792}}})

	feature6 := NewFeature()
	feature6.SetProperty("6", "6")
	feature6.SetID("0006")
	feature6.SetGeometry(pol)
	// *****

	multpol := NewMultiPolygon([][][][]float64{
		[][][]float64{
			[][]float64{
				[]float64{37.55744934082031, 55.76189525593947},
				[]float64{37.56088256835937, 55.76150892439349},
				[]float64{37.596588134765625, 55.74373353535969},
				[]float64{37.58354187011719, 55.78622642787003},
				[]float64{37.55744934082031, 55.76189525593947},
			}},
		[][][]float64{
			[][]float64{
				[]float64{38.08904439210892, 55.81190491362447},
				[]float64{38.08977395296096, 55.81190491362447},
				[]float64{38.08977395296096, 55.81212194465421},
				[]float64{38.08904439210892, 55.81212194465421},
				[]float64{38.08904439210892, 55.81190491362447}}}})

	feature7 := NewFeature()
	feature7.SetProperty("6", "6")
	feature7.SetID("0006")
	feature7.SetGeometry(multpol)
	// *****

	multPoint000 := NewMultiPoint([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.596588134765625, 55.72169627105833},
	})

	feature8 := NewFeature()
	feature8.SetProperty("8", "8")
	feature8.SetID("008")
	feature8.SetGeometry(multPoint000)
	// *****

	linestr0 := NewLineString([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.596588134765625, 55.72169627105833}})

	feature9 := NewFeature()
	feature9.SetProperty("9", "9")
	feature9.SetID("10")
	feature9.SetGeometry(linestr0)
	// *****

	poly := NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.57770538330078, 55.76691721773862},
			[]float64{37.57942199707031, 55.7659515060885},
			[]float64{37.63298034667969, 55.73406960018026},
			[]float64{37.644996643066406, 55.76884856927786},
			[]float64{37.57770538330078, 55.76691721773862}}})

	feature10 := NewFeature()
	feature10.SetProperty("10", "10]")
	feature10.SetID("010")
	feature10.SetGeometry(poly)
	// *****

	multPoint := NewMultiPoint([][]float64{
		[]float64{37.69306182861328, 55.78950804222006},
		[]float64{37.71228790283203, 55.79028014659872},
		[]float64{37.71263122558593, 55.76981420916759}})

	feature11 := NewFeature()
	feature11.SetProperty("11", "11")
	feature11.SetID("011")
	feature11.SetGeometry(multPoint)
	// *****

	// 25
	mult0 := NewMultiPoint([][]float64{[]float64{37.570792444518, 55.784581739175},
		[]float64{37.624941407105, 55.738048529295}})

	feature12 := NewFeature()
	feature12.SetProperty("12", "12")
	feature12.SetID("012")
	feature12.SetGeometry(mult0)

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

	feature13 := NewFeature()
	feature13.SetProperty("13", "13")
	feature13.SetID("013")
	feature13.SetGeometry(multipol)
	// *****

	linestr14 := NewLineString([][]float64{
		[]float64{38.408203125, 57.61010702068388},
		[]float64{34.1455078125, 54.213861000644926},
		[]float64{40.3857421875, 52.348763181988105},
		[]float64{39.7265625, 48.922499263758255},
	})
	feature14 := NewFeature()
	feature14.SetProperty("14", "14")
	feature14.SetID("00014")
	feature14.SetGeometry(linestr14)
	// *****

	linestr15 := NewLineString([][]float64{
		[]float64{38.935546875, 55.229023057406344},
		[]float64{41.3525390625, 53.51418452077113},
	})
	feature15 := NewFeature()
	feature15.SetProperty("15", "15")
	feature15.SetID("0015")
	feature15.SetGeometry(linestr15)
	// *****

	linestr16 := NewLineString([][]float64{
		[]float64{37.60499954223633, 55.78979758315554},
		[]float64{37.634525299072266, 55.799158247007924},
		[]float64{37.64293670654297, 55.79607044162981},
		[]float64{37.659759521484375, 55.79935122671461},
	})
	feature16 := NewFeature()
	feature16.SetProperty("16", "16")
	feature16.SetID("016")
	feature16.SetGeometry(linestr16)
	// *****

	multLine17 := NewMultiLineString([][][]float64{
		[][]float64{[]float64{37.58766174316406, 55.78178615666911},
			[]float64{37.545433044433594, 55.74740520331641},
			[]float64{37.57495880126953, 55.71589492312098}},

		[][]float64{[]float64{37.63092041015624, 55.71492794801965},
			[]float64{37.61787414550781, 55.72769009202728},
			[]float64{37.64396667480469, 55.74817814201809},
			[]float64{37.69718170166015, 55.71879570480289}}})

	feature17 := NewFeature()
	feature17.SetProperty("17", "17")
	feature17.SetID("017")
	feature17.SetGeometry(multLine17)
	// *****

	multLine18 := NewMultiLineString([][][]float64{
		[][]float64{
			[]float64{37.62165069580078, 55.8238518767974},
			[]float64{37.644309997558594, 55.77946929254524},
			[]float64{37.67864227294922, 55.789701069749528}},

		[][]float64{
			[]float64{37.662506103515625, 55.834457609133935},
			[]float64{37.66490936279297, 55.810156563965364},
			[]float64{37.73700714111328, 55.73387629706783}}})

	feature18 := NewFeature()
	feature18.SetProperty("18", "18")
	feature18.SetID("018")
	feature18.SetGeometry(multLine18)
	// *****

	polp := NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.48809814453125, 55.867605966997786},
			[]float64{37.522430419921875, 55.831372603449026},
			[]float64{37.615814208984375, 55.83522882231773},
			[]float64{37.6007080078125, 55.8845546603819},
			[]float64{37.48809814453125, 55.867605966997786}}})

	feature19 := NewFeature()
	feature19.SetProperty("19", "19")
	feature19.SetID("00019")
	feature19.SetGeometry(polp)
	// *****

	mu := NewMultiPolygon([][][][]float64{

		[][][]float64{
			[][]float64{

				[]float64{37.55744934082031, 55.76189525593947},
				[]float64{37.56088256835937, 55.76150892439349},
				[]float64{37.596588134765625, 55.74373353535969},
				[]float64{37.58354187011719, 55.78622642787003},
				[]float64{37.55744934082031, 55.76189525593947},
			}},

		[][][]float64{
			[][]float64{

				[]float64{37.63298034667969, 55.789121984291626},
				[]float64{37.61821746826172, 55.77135918323877},
				[]float64{37.68241882324219, 55.773097205876766},
				[]float64{37.677955627441406, 55.82404473410693},
				[]float64{37.63298034667969, 55.789121984291626}}}})

	feature20 := NewFeature()
	feature20.SetProperty("Полигон", "[][][]")
	feature20.SetID("020")
	feature20.SetGeometry(mu)
	// *****

	pol21 := NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.503204345703125, 55.85835810656004},
			[]float64{37.59796142578125, 55.84640969483053},
			[]float64{37.62611389160156, 55.87261430554018},
			[]float64{37.574615478515625, 55.90573012454021},
			[]float64{37.503204345703125, 55.85835810656004}}})

	feature21 := NewFeature()
	feature21.SetProperty("Полигон", "[][][]")
	feature21.SetID("021")
	feature21.SetGeometry(pol21)
	// *****
	// *****
	// *****

	// Point - Point
	assert.True(Distance(feature1, feature2) == Distance(feature2, feature1))
	assert.True(Distance(feature1, feature2) == 1639.8005076177767)
	// Point - MultiPoint
	assert.True(Distance(feature1, feature3) == Distance(feature3, feature1))
	assert.True(Distance(feature1, feature3) == 1619.739745592433)
	// Point - LineString
	assert.True(Distance(feature1, feature4) == Distance(feature4, feature1))
	assert.True(Distance(feature1, feature4) == 459.579558062514)
	// Point - MultiLineString
	assert.True(Distance(feature1, feature5) == Distance(feature5, feature1))
	assert.True(Distance(feature1, feature5) == 163.19512713815737)
	// Point - Polygon
	assert.True(Distance(feature1, feature6) == Distance(feature6, feature1))
	assert.True(Distance(feature1, feature6) == 315.1278369461277)
	// Point - MultiPolygon
	assert.True(Distance(feature1, feature7) == Distance(feature7, feature1))
	assert.True(Distance(feature1, feature7) == 1445.9461020339781)

	// MultiPoint - MultiPoint
	assert.True(Distance(feature3, feature8) == Distance(feature8, feature3))
	assert.True(Distance(feature3, feature8) == 2541.227427018973)
	// MultiPoint - LineString
	assert.True(Distance(feature3, feature9) == Distance(feature9, feature3))
	assert.True(Distance(feature3, feature9) == 2410.3293967279924)
	// MultiPoint - MultiLineString
	assert.True(Distance(feature3, feature5) == Distance(feature5, feature3))
	assert.True(Distance(feature3, feature5) == 736.5211184694738)
	// MultiPoint - Polygon
	assert.True(Distance(feature10, feature11) == Distance(feature11, feature10))
	assert.True(Distance(feature10, feature11) == 3783.0687405935255)
	// MultiPoint - MultiPolygon
	assert.True(Distance(feature12, feature13) == Distance(feature13, feature12))
	assert.True(Distance(feature12, feature13) == 61.95584692761134)

	// LineString - LineString
	assert.True(Distance(feature14, feature15) == Distance(feature15, feature14))
	assert.True(Distance(feature14, feature15) == 144703.63893885718)

	// LineString - MultiLineString
	assert.True(Distance(feature14, feature5) == Distance(feature5, feature14))
	assert.True(Distance(feature14, feature5) == 84091.30664614956)
	// LineString - Polygon
	assert.True(Distance(feature14, feature6) == Distance(feature6, feature14))
	assert.True(Distance(feature14, feature6) == 83203.32454881325)
	// LineString - MultiPolygon
	assert.True(Distance(feature13, feature16) == Distance(feature16, feature13))
	assert.True(Distance(feature13, feature16) == 1000.5855306843869)

	// MultiLineString - MultiLineString
	assert.True(Distance(feature17, feature18) == Distance(feature18, feature17))
	assert.True(Distance(feature17, feature18) == 3005.1838191535385)
	// MultiLineString - Polygon
	assert.True(Distance(feature6, feature18) == Distance(feature18, feature6))
	assert.True(Distance(feature6, feature18) == 3638.124402575424)
	// MultiLineString - MultiPolygon
	assert.True(Distance(feature7, feature18) == Distance(feature18, feature7))
	assert.True(Distance(feature7, feature18) == 3444.600490067915)

	// Polygon - Polygon
	assert.True(Distance(feature10, feature19) == Distance(feature19, feature10))
	assert.True(Distance(feature10, feature19) == 7401.963882373347)
	// Polygon - MultiPolygon

	// MultiPolygon - MultiPolygon

}

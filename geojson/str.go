package geogoth

// Ftrs ...
var Ftrs struct {
	Feature1  *Feature
	Feature2  *Feature
	Feature3  *Feature
	Feature4  *Feature
	Feature5  *Feature
	Feature6  *Feature
	Feature7  *Feature
	Feature8  *Feature
	Feature9  *Feature
	Feature10 *Feature
	Feature11 *Feature
	Feature12 *Feature
	Feature13 *Feature
	Feature14 *Feature
	Feature15 *Feature
	Feature16 *Feature
	Feature17 *Feature
	Feature18 *Feature
	Feature19 *Feature
	Feature20 *Feature
	Feature21 *Feature
}

// CreateFtr1 ...
func CreateFtr1() {
	ftr := &Ftrs

	point1 := NewPoint([]float64{37.6175, 55.752})
	ftr.Feature1 = NewFeature()
	ftr.Feature1.SetProperty("1", "1")
	ftr.Feature1.SetID("0001")
	ftr.Feature1.SetGeometry(point1)

}

// CreateFtr2 ...
func CreateFtr2() {
	ftr := &Ftrs

	point2 := NewPoint([]float64{37.6048, 55.7649})
	ftr.Feature2 = NewFeature()
	ftr.Feature2.SetProperty("2", "2")
	ftr.Feature2.SetID("0002")
	ftr.Feature2.SetGeometry(point2)
}

// CreateFtr3 ...
func CreateFtr3() {
	ftr := &Ftrs

	mult := NewMultiPoint([][]float64{[]float64{37.570792444518, 55.784581739175},
		[]float64{37.624941407105, 55.738048529295}, []float64{37.633109750449, 55.77134814581}})

	ftr.Feature3 = NewFeature()
	ftr.Feature3.SetProperty("3", "3")
	ftr.Feature3.SetID("0003")
	ftr.Feature3.SetGeometry(mult)

}

// CreateFtr4 ...
func CreateFtr4() {
	ftr := &Ftrs

	linestr := NewLineString([][]float64{

		[]float64{37.624941407105, 55.738048529295},
		[]float64{37.570792444518, 55.784581739175}})

	ftr.Feature4 = NewFeature()
	ftr.Feature4.SetProperty("4", "4")
	ftr.Feature4.SetID("0004")
	ftr.Feature4.SetGeometry(linestr)
}

// CreateFtr5 ...
func CreateFtr5() {
	ftr := &Ftrs

	multLine := NewMultiLineString([][][]float64{
		[][]float64{[]float64{37.61911869049072, 55.75634137200424},
			[]float64{37.626585960388184, 55.75218751578049},
			[]float64{37.62401103973388, 55.74465148474084}},
		[][]float64{[]float64{37.614097595214844, 55.7528396038768},
			[]float64{37.62409687042236, 55.754989002265646},
			[]float64{37.6255989074707, 55.74665640434626},
			[]float64{37.62989044189453, 55.74503798328292}}})

	ftr.Feature5 = NewFeature()
	ftr.Feature5.SetProperty("5", "5")
	ftr.Feature5.SetID("0005")
	ftr.Feature5.SetGeometry(multLine)
}

// CreateFtr6 ...
func CreateFtr6() {
	ftr := &Ftrs

	pol := NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.60027885437012, 55.75573764130792},
			[]float64{37.60027885437012, 55.75556859503827},
			[]float64{37.59929180145264, 55.751632024729965},
			[]float64{37.60671615600586, 55.74986889204201},
			[]float64{37.61246681213379, 55.75209090920951},
			[]float64{37.60027885437012, 55.75573764130792}}})

	ftr.Feature6 = NewFeature()
	ftr.Feature6.SetProperty("6", "6")
	ftr.Feature6.SetID("0006")
	ftr.Feature6.SetGeometry(pol)
}

// CreateFtr7 ...
func CreateFtr7() {
	ftr := &Ftrs

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

	ftr.Feature7 = NewFeature()
	ftr.Feature7.SetProperty("6", "6")
	ftr.Feature7.SetID("0006")
	ftr.Feature7.SetGeometry(multpol)
}

// CreateFtr8 ...
func CreateFtr8() {
	ftr := &Ftrs

	multPoint000 := NewMultiPoint([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.596588134765625, 55.72169627105833},
	})

	ftr.Feature8 = NewFeature()
	ftr.Feature8.SetProperty("8", "8")
	ftr.Feature8.SetID("008")
	ftr.Feature8.SetGeometry(multPoint000)
}

// CreateFtr9 ...
func CreateFtr9() {
	ftr := &Ftrs

	linestr0 := NewLineString([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.596588134765625, 55.72169627105833}})

	ftr.Feature9 = NewFeature()
	ftr.Feature9.SetProperty("9", "9")
	ftr.Feature9.SetID("10")
	ftr.Feature9.SetGeometry(linestr0)
}

// CreateFtr10 ...
func CreateFtr10() {
	ftr := &Ftrs

	poly := NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.57770538330078, 55.76691721773862},
			[]float64{37.57942199707031, 55.7659515060885},
			[]float64{37.63298034667969, 55.73406960018026},
			[]float64{37.644996643066406, 55.76884856927786},
			[]float64{37.57770538330078, 55.76691721773862}}})

	ftr.Feature10 = NewFeature()
	ftr.Feature10.SetProperty("10", "10]")
	ftr.Feature10.SetID("010")
	ftr.Feature10.SetGeometry(poly)
}

// CreateFtr11 ...
func CreateFtr11() {
	ftr := &Ftrs

	multPoint := NewMultiPoint([][]float64{
		[]float64{37.69306182861328, 55.78950804222006},
		[]float64{37.71228790283203, 55.79028014659872},
		[]float64{37.71263122558593, 55.76981420916759}})

	ftr.Feature11 = NewFeature()
	ftr.Feature11.SetProperty("11", "11")
	ftr.Feature11.SetID("011")
	ftr.Feature11.SetGeometry(multPoint)
}

// CreateFtr12 ...
func CreateFtr12() {
	ftr := &Ftrs

	mult0 := NewMultiPoint([][]float64{[]float64{37.570792444518, 55.784581739175},
		[]float64{37.624941407105, 55.738048529295}})

	ftr.Feature12 = NewFeature()
	ftr.Feature12.SetProperty("12", "12")
	ftr.Feature12.SetID("012")
	ftr.Feature12.SetGeometry(mult0)

}

// CreateFtr13 ...
func CreateFtr13() {
	ftr := &Ftrs

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

	ftr.Feature13 = NewFeature()
	ftr.Feature13.SetProperty("13", "13")
	ftr.Feature13.SetID("013")
	ftr.Feature13.SetGeometry(multipol)
}

// CreateFtr14 ...
func CreateFtr14() {
	ftr := &Ftrs

	linestr14 := NewLineString([][]float64{
		[]float64{38.408203125, 57.61010702068388},
		[]float64{34.1455078125, 54.213861000644926},
		[]float64{40.3857421875, 52.348763181988105},
		[]float64{39.7265625, 48.922499263758255},
	})
	ftr.Feature14 = NewFeature()
	ftr.Feature14.SetProperty("14", "14")
	ftr.Feature14.SetID("00014")
	ftr.Feature14.SetGeometry(linestr14)
}

// CreateFtr15 ...
func CreateFtr15() {
	ftr := &Ftrs

	linestr15 := NewLineString([][]float64{
		[]float64{38.935546875, 55.229023057406344},
		[]float64{41.3525390625, 53.51418452077113},
	})
	ftr.Feature15 = NewFeature()
	ftr.Feature15.SetProperty("15", "15")
	ftr.Feature15.SetID("0015")
	ftr.Feature15.SetGeometry(linestr15)
}

// CreateFtr16 ...
func CreateFtr16() {
	ftr := &Ftrs

	linestr16 := NewLineString([][]float64{
		[]float64{37.60499954223633, 55.78979758315554},
		[]float64{37.634525299072266, 55.799158247007924},
		[]float64{37.64293670654297, 55.79607044162981},
		[]float64{37.659759521484375, 55.79935122671461},
	})
	ftr.Feature16 = NewFeature()
	ftr.Feature16.SetProperty("16", "16")
	ftr.Feature16.SetID("016")
	ftr.Feature16.SetGeometry(linestr16)
}

// CreateFtr17 ...
func CreateFtr17() {
	ftr := &Ftrs

	multLine17 := NewMultiLineString([][][]float64{
		[][]float64{[]float64{37.58766174316406, 55.78178615666911},
			[]float64{37.545433044433594, 55.74740520331641},
			[]float64{37.57495880126953, 55.71589492312098}},

		[][]float64{[]float64{37.63092041015624, 55.71492794801965},
			[]float64{37.61787414550781, 55.72769009202728},
			[]float64{37.64396667480469, 55.74817814201809},
			[]float64{37.69718170166015, 55.71879570480289}}})

	ftr.Feature17 = NewFeature()
	ftr.Feature17.SetProperty("17", "17")
	ftr.Feature17.SetID("017")
	ftr.Feature17.SetGeometry(multLine17)
}

// CreateFtr18 ...
func CreateFtr18() {
	ftr := &Ftrs

	multLine18 := NewMultiLineString([][][]float64{
		[][]float64{
			[]float64{37.62165069580078, 55.8238518767974},
			[]float64{37.644309997558594, 55.77946929254524},
			[]float64{37.67864227294922, 55.789701069749528}},

		[][]float64{
			[]float64{37.662506103515625, 55.834457609133935},
			[]float64{37.66490936279297, 55.810156563965364},
			[]float64{37.73700714111328, 55.73387629706783}}})

	ftr.Feature18 = NewFeature()
	ftr.Feature18.SetProperty("18", "18")
	ftr.Feature18.SetID("018")
	ftr.Feature18.SetGeometry(multLine18)
}

// CreateFtr19 ...
func CreateFtr19() {
	ftr := &Ftrs

	polp := NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.48809814453125, 55.867605966997786},
			[]float64{37.522430419921875, 55.831372603449026},
			[]float64{37.615814208984375, 55.83522882231773},
			[]float64{37.6007080078125, 55.8845546603819},
			[]float64{37.48809814453125, 55.867605966997786}}})

	ftr.Feature19 = NewFeature()
	ftr.Feature19.SetProperty("19", "19")
	ftr.Feature19.SetID("00019")
	ftr.Feature19.SetGeometry(polp)
}

// CreateFtr20 ...
func CreateFtr20() {
	ftr := &Ftrs

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

	ftr.Feature20 = NewFeature()
	ftr.Feature20.SetProperty("Полигон", "[][][]")
	ftr.Feature20.SetID("020")
	ftr.Feature20.SetGeometry(mu)
}

// CreateFtr21 ...
func CreateFtr21() {
	ftr := &Ftrs

	pol21 := NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.503204345703125, 55.85835810656004},
			[]float64{37.59796142578125, 55.84640969483053},
			[]float64{37.62611389160156, 55.87261430554018},
			[]float64{37.574615478515625, 55.90573012454021},
			[]float64{37.503204345703125, 55.85835810656004}}})

	ftr.Feature21 = NewFeature()
	ftr.Feature21.SetProperty("Полигон", "[][][]")
	ftr.Feature21.SetID("021")
	ftr.Feature21.SetGeometry(pol21)
}

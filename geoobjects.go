package geogoth

// Objects struct for geo objects
var Objects struct {
	Object1  Point
	Object2  Point
	Object3  MultiPoint
	Object4  LineString
	Object5  MultiLineString
	Object6  Polygon
	Object7  MultiPolygon
	Object8  MultiPoint
	Object9  LineString
	Object10 Polygon
	Object11 MultiPoint
	Object12 MultiPoint
	Object13 MultiPolygon
	Object14 LineString
	Object15 LineString
	Object16 LineString
	Object17 MultiLineString
	Object18 MultiLineString
	Object19 Polygon
	// ***
	Object20 MultiPoint
}

// CreateObj1 creates Point
func CreateObj1() {
	Obj := &Objects
	Obj.Object1 = NewPoint(37.6175, 55.752)
}

// CreateObj2 creates Point
func CreateObj2() {
	Obj := &Objects
	Obj.Object2 = NewPoint(37.6048, 55.7649)
}

// CreateObj3 creates MultiPoint
func CreateObj3() {
	Obj := &Objects
	Obj.Object3 = NewMultiPoint([][]float64{
		[]float64{37.570792444518, 55.784581739175},
		[]float64{37.624941407105, 55.738048529295},
		[]float64{37.633109750449, 55.77134814581}})
}

// CreateObj4 creates LineString
func CreateObj4() {
	Obj := &Objects
	Obj.Object4 = NewLineString([][]float64{
		[]float64{37.624941407105, 55.738048529295},
		[]float64{37.570792444518, 55.784581739175}})

}

// CreateObj5 creates MultiLineString
func CreateObj5() {
	Obj := &Objects
	Obj.Object5 = NewMultiLineString([][][]float64{
		[][]float64{[]float64{37.61911869049072, 55.75634137200424},
			[]float64{37.626585960388184, 55.75218751578049},
			[]float64{37.62401103973388, 55.74465148474084}},
		[][]float64{[]float64{37.614097595214844, 55.7528396038768},
			[]float64{37.62409687042236, 55.754989002265646},
			[]float64{37.6255989074707, 55.74665640434626},
			[]float64{37.62989044189453, 55.74503798328292}}})

}

// CreateObj6 creates Polygon
func CreateObj6() {
	Obj := &Objects
	Obj.Object6 = NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.60027885437012, 55.75573764130792},
			[]float64{37.60027885437012, 55.75556859503827},
			[]float64{37.59929180145264, 55.751632024729965},
			[]float64{37.60671615600586, 55.74986889204201},
			[]float64{37.61246681213379, 55.75209090920951},
			[]float64{37.60027885437012, 55.75573764130792}}})

}

// CreateObj7 creates MultiPolygon
func CreateObj7() {
	Obj := &Objects
	Obj.Object7 = NewMultiPolygon([][][][]float64{
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

}

// CreateObj8  creates MultiPoint
func CreateObj8() {
	Obj := &Objects
	Obj.Object8 = NewMultiPoint([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.596588134765625, 55.72169627105833},
	})

}

// CreateObj9  creates MultiPoint
func CreateObj9() {
	Obj := &Objects
	Obj.Object9 = NewLineString([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.596588134765625, 55.72169627105833}})

}

// CreateObj10  creates Polygon
func CreateObj10() {
	Obj := &Objects
	Obj.Object10 = NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.57770538330078, 55.76691721773862},
			[]float64{37.57942199707031, 55.7659515060885},
			[]float64{37.63298034667969, 55.73406960018026},
			[]float64{37.644996643066406, 55.76884856927786},
			[]float64{37.57770538330078, 55.76691721773862}}})

}

// CreateObj11 creates MultiPoint
func CreateObj11() {
	Obj := &Objects
	Obj.Object11 = NewMultiPoint([][]float64{
		[]float64{37.69306182861328, 55.78950804222006},
		[]float64{37.71228790283203, 55.79028014659872},
		[]float64{37.71263122558593, 55.76981420916759}})

}

// CreateObj12 creates MultiPoint
func CreateObj12() {
	Obj := &Objects
	Obj.Object12 = NewMultiPoint([][]float64{
		[]float64{37.570792444518, 55.784581739175},
		[]float64{37.624941407105, 55.738048529295}})

}

// CreateObj13 creates MultiPolygon
func CreateObj13() {
	Obj := &Objects
	Obj.Object13 = NewMultiPolygon([][][][]float64{

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

}

// CreateObj14   creates
func CreateObj14() {
	Obj := &Objects
	Obj.Object14 = NewLineString([][]float64{
		[]float64{38.408203125, 57.61010702068388},
		[]float64{34.1455078125, 54.213861000644926},
		[]float64{40.3857421875, 52.348763181988105},
		[]float64{39.7265625, 48.922499263758255}})

}

// CreateObj15   creates
func CreateObj15() {
	Obj := &Objects
	Obj.Object15 = NewLineString([][]float64{
		[]float64{38.935546875, 55.229023057406344},
		[]float64{41.3525390625, 53.51418452077113}})

}

// CreateObj16 creates
func CreateObj16() {
	Obj := &Objects
	Obj.Object16 = NewLineString([][]float64{
		[]float64{37.60499954223633, 55.78979758315554},
		[]float64{37.634525299072266, 55.799158247007924},
		[]float64{37.64293670654297, 55.79607044162981},
		[]float64{37.659759521484375, 55.79935122671461},
	})

}

// CreateObj17 creates
func CreateObj17() {
	Obj := &Objects
	Obj.Object17 = NewMultiLineString([][][]float64{
		[][]float64{[]float64{37.58766174316406, 55.78178615666911},
			[]float64{37.545433044433594, 55.74740520331641},
			[]float64{37.57495880126953, 55.71589492312098}},

		[][]float64{[]float64{37.63092041015624, 55.71492794801965},
			[]float64{37.61787414550781, 55.72769009202728},
			[]float64{37.64396667480469, 55.74817814201809},
			[]float64{37.69718170166015, 55.71879570480289}}})

}

// CreateObj18 creates
func CreateObj18() {
	Obj := &Objects
	Obj.Object18 = NewMultiLineString([][][]float64{
		[][]float64{
			[]float64{37.62165069580078, 55.8238518767974},
			[]float64{37.644309997558594, 55.77946929254524},
			[]float64{37.67864227294922, 55.789701069749528}},

		[][]float64{
			[]float64{37.662506103515625, 55.834457609133935},
			[]float64{37.66490936279297, 55.810156563965364},
			[]float64{37.73700714111328, 55.73387629706783}}})

}

// CreateObj19 creates
func CreateObj19() {
	Obj := &Objects
	Obj.Object19 = NewPolygon([][][]float64{
		[][]float64{
			[]float64{37.48809814453125, 55.867605966997786},
			[]float64{37.522430419921875, 55.831372603449026},
			[]float64{37.615814208984375, 55.83522882231773},
			[]float64{37.6007080078125, 55.8845546603819},
			[]float64{37.48809814453125, 55.867605966997786}}})

}

// CreateObj20  creates MultiPoint
func CreateObj20() {
	Obj := &Objects
	Obj.Object20 = NewMultiPoint([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.6175, 55.752},
		[]float64{37.596588134765625, 55.72169627105833},
	})

}

// // CreateObj   creates
// func CreateObj() {
// 	Obj := &Objects
// 	Obj.Object

// }

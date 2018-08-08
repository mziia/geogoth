package geogoth

// Objects struct for geo objects
var Objects struct {
	Object1 Point
	Object2 Point
	Object3 MultiPoint
	Object4 LineString
	Object5 MultiLineString
	Object6 Polygon
	Object7 MultiPolygon
	Object8 MultiPoint
	Object9 LineString
	// Object10
	// Object11
	// Object12
	// Object13
	// Object14
	// Object15
	// Object16
	// Object17
	// Object18
	// Object19
	// Object20
	// Object21
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

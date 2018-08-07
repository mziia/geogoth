package geogoth

// Objects struct for geo objects
var Objects struct {
	Object1 Point
	Object2 Point
	Object3 MultiPoint
	Object4 LineString
	Object5 MultiLineString
	Object6 Polygon
	// Object7
	// Object8
	// Object9
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

// CreateObj6 creates MultiLineString
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

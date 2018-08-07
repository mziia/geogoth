package geogoth

// Objects struct for geo objects
var Objects struct {
	Object1 Point
	Object2 Point
	Object3 MultiPoint
	Object4 LineString
	// Object5
	// Object6
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

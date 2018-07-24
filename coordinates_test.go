package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestNewPoint(t *testing.T) {
	assert := assert.New(t)

	point := NewPoint([]float64{37.6175, 55.752}) // lon, lat

	p := (point.Coordinates).([]float64)
	assert.True(p[0] == 37.6175)
	assert.True(p[1] == 55.752)

}

func TestGetPointCoordinates(t *testing.T) {
	assert := assert.New(t)

	point := NewPoint([]float64{37.6175, 55.752})

	feature := NewFeature()
	feature.SetProperty("Point", "[]")
	feature.SetID("000")
	feature.SetGeometry(point)

	y, x := GetPointCoordinates(feature)

	assert.True(y == 37.6175)
	assert.True(x == 55.752)

}

func TestNewMultiPoint(t *testing.T) {
	assert := assert.New(t)

	mult := NewMultiPoint([][]float64{[]float64{37.570792444518, 55.784581739175},
		[]float64{37.624941407105, 55.738048529295}, []float64{37.633109750449, 55.77134814581}})

	m := (mult.Coordinates).([][]float64)

	assert.True(m[0][0] == 37.570792444518)
	assert.True(m[0][1] == 55.784581739175)

	assert.True(m[2][0] == 37.633109750449)
	assert.True(m[2][1] == 55.77134814581)

}

func TestNewLineString(t *testing.T) {
	assert := assert.New(t)

	linestr := NewLineString([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.596588134765625, 55.72169627105833}})

	l := (linestr.Coordinates).([][]float64)

	assert.True(l[0][0] == 37.566375732421875)
	assert.True(l[0][1] == 55.761702090644896)

	assert.True(l[3][0] == 37.596588134765625)
	assert.True(l[3][1] == 55.72169627105833)

}

func TestGetTwoDimArrayCoordinates(t *testing.T) {
	assert := assert.New(t)

	mult := NewMultiPoint([][]float64{[]float64{37.570792444518, 55.784581739175},
		[]float64{37.624941407105, 55.738048529295}, []float64{37.633109750449, 55.77134814581}})

	featureP := NewFeature()
	featureP.SetProperty("MPoint", "[][][]")
	featureP.SetID("000")
	featureP.SetGeometry(mult)

	yP, xP := GetTwoDimArrayCoordinates(featureP, 1)

	assert.True(yP == 37.624941407105)
	assert.True(xP == 55.738048529295)

	linestr := NewLineString([][]float64{
		[]float64{37.566375732421875, 55.761702090644896},
		[]float64{37.58800506591796, 55.74856460562653},
		[]float64{37.58491516113281, 55.72981671057788},
		[]float64{37.596588134765625, 55.72169627105833}})

	featureLn := NewFeature()
	featureLn.SetProperty("LineStr", "[][][]")
	featureLn.SetID("000")
	featureLn.SetGeometry(linestr)

	yL, xL := GetTwoDimArrayCoordinates(featureLn, 1)

	assert.True(yL == 37.58800506591796)
	assert.True(xL == 55.74856460562653)

}

func TestNewMultiLineString(t *testing.T) {
	assert := assert.New(t)

	mLine := NewMultiLineString([][][]float64{
		[][]float64{[]float64{37.58766174316406, 55.78178615666911},
			[]float64{37.545433044433594, 55.74740520331641},
			[]float64{37.57495880126953, 55.71589492312098}},

		[][]float64{[]float64{37.63092041015624, 55.71492794801965},
			[]float64{37.61787414550781, 55.72769009202728},
			[]float64{37.64396667480469, 55.74817814201809},
			[]float64{37.69718170166015, 55.71879570480289}}})

	l := (mLine.Coordinates).([][][]float64)

	assert.True(l[0][1][0] == 37.545433044433594)
	assert.True(l[0][1][1] == 55.74740520331641)

	assert.True(l[1][1][0] == 37.61787414550781)
	assert.True(l[1][1][1] == 55.72769009202728)

}

func TestNewPolygon(t *testing.T) {
	assert := assert.New(t)

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

	p := (polyg.Coordinates).([][][]float64)

	assert.True(p[0][1][0] == 37.63023376464844)
	assert.True(p[0][1][1] == 55.62993418221071)

	assert.True(p[1][1][0] == 37.62062072753906)
	assert.True(p[1][1][1] == 55.73870858770892)

}

func TestGetThreeDimArrayCoordinates(t *testing.T) {
	assert := assert.New(t)

	mLine := NewMultiLineString([][][]float64{
		[][]float64{[]float64{37.58766174316406, 55.78178615666911},
			[]float64{37.545433044433594, 55.74740520331641},
			[]float64{37.57495880126953, 55.71589492312098}},

		[][]float64{[]float64{37.63092041015624, 55.71492794801965},
			[]float64{37.61787414550781, 55.72769009202728},
			[]float64{37.64396667480469, 55.74817814201809},
			[]float64{37.69718170166015, 55.71879570480289}}})

	featureML := NewFeature()
	featureML.SetProperty("MLineStr", "[][][]")
	featureML.SetID("000")
	featureML.SetGeometry(mLine)

	yML, xML := GetThreeDimArrayCoordinates(featureML, 0, 2)

	assert.True(yML == 37.57495880126953)
	assert.True(xML == 55.71589492312098)

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

	featureP := NewFeature()
	featureP.SetProperty("Polyg", "[][][]")
	featureP.SetID("000")
	featureP.SetGeometry(polyg)

	yP, xP := GetThreeDimArrayCoordinates(featureP, 1, 2)

	assert.True(yP == 37.651519775390625)
	assert.True(xP == 55.792017325495095)

}

func TestNewMultiPolygon(t *testing.T) {
	assert := assert.New(t)

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

	mp := (multpol.Coordinates).([][][][]float64)

	assert.True(mp[0][0][1][0] == 37.56088256835937)
	assert.True(mp[0][0][1][1] == 55.76150892439349)

}

func TestGetFourDimArrayCoordinates(t *testing.T) {
	assert := assert.New(t)

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

	feature := NewFeature()
	feature.SetProperty("MPolyg", "[][][][]")
	feature.SetID("000")
	feature.SetGeometry(multpol)

	y, x := GetFourDimArrayCoordinates(feature, 1, 0, 3)

	assert.True(y == 38.08904439210892)
	assert.True(x == 55.81212194465421)

}

package geogoth

import (
	"testing"
)

func TestNewPoint(t *testing.T) {
	point := NewPoint([]float64{37.6175, 55.752}) // lon, lat
	p := (point.Coordinates).([]float64)

	if p[0] != 37.6175 {
		t.Error("expexted p[0] == 37.6175, but p[0] == ", p[0])
	}

	if p[1] != 55.752 {
		t.Error("expexted p[1] == 55.752, but p[1] == ", p[0])
	}
}

func TestGetPointCoordinates(t *testing.T) {

	point := NewPoint([]float64{37.6175, 55.752})

	feature := NewFeature()
	feature.SetProperty("Point", "[]")
	feature.SetID("000")
	feature.SetGeometry(point)

	y, x := GetPointCoordinates(feature)

	if y != 37.6175 {
		t.Error("expexted x == 37.6175, but x == ", x)
	}

	if x != 55.752 {
		t.Error("expexted y == 55.752, but y == ", y)
	}
}

func TestNewMultiPoint(t *testing.T) {

	mult := NewMultiPoint([][]float64{{37.570792444518, 55.784581739175},
		{37.624941407105, 55.738048529295}, {37.633109750449, 55.77134814581}})

	m := (mult.Coordinates).([][]float64)

	if m[0][0] != 37.570792444518 {
		t.Error("expexted mp[0][0] == 37.570792444518, but m[0][0] == ", m[0][0])
	}

	if m[0][1] != 55.784581739175 {
		t.Error("expexted mp[0][1] == 55.784581739175, but m[][] == ", m[0][1])
	}
	if m[2][0] != 37.633109750449 {
		t.Error("expexted mp[2][0] == 37.633109750449, but m[][] == ", m[2][0])
	}
	if m[2][1] != 55.77134814581 {
		t.Error("expexted mp[2][1] == 55.77134814581, but m[][] == ", m[2][1])
	}

}

func TestNewLineString(t *testing.T) {

	linestr := NewLineString([][]float64{
		{37.566375732421875, 55.761702090644896},
		{37.58800506591796, 55.74856460562653},
		{37.58491516113281, 55.72981671057788},
		{37.596588134765625, 55.72169627105833}})

	l := (linestr.Coordinates).([][]float64)

	if l[0][0] != 37.566375732421875 {
		t.Error("expexted l[0][0] == , but l[0][0] == ", l[0][0])
	}
	if l[0][1] != 55.761702090644896 {
		t.Error("expexted l[0][1] == , but l[0][1] == ", l[0][1])
	}

	if l[3][0] != 37.596588134765625 {
		t.Error("expexted l[3][0] == , but l[3][0] == ", l[0][0])
	}
	if l[3][1] != 55.72169627105833 {
		t.Error("expexted l[3][1] == , but l[3][1] == ", l[0][1])
	}

}

func TestGetTwoDimArrayCoordinates(t *testing.T) {

	mult := NewMultiPoint([][]float64{{37.570792444518, 55.784581739175},
		{37.624941407105, 55.738048529295}, {37.633109750449, 55.77134814581}})

	featureP := NewFeature()
	featureP.SetProperty("MPoint", "[][][]")
	featureP.SetID("000")
	featureP.SetGeometry(mult)

	yP, xP := GetTwoDimArrayCoordinates(featureP, 1)

	if yP != 37.624941407105 {
		t.Error("expected yP = 37.624941407105, but yP = ", yP)
	}

	if xP != 55.738048529295 {
		t.Error("expected xP = 55.738048529295, but xP = ", xP)
	}

	linestr := NewLineString([][]float64{
		{37.566375732421875, 55.761702090644896},
		{37.58800506591796, 55.74856460562653},
		{37.58491516113281, 55.72981671057788},
		{37.596588134765625, 55.72169627105833}})

	featureLn := NewFeature()
	featureLn.SetProperty("LineStr", "[][][]")
	featureLn.SetID("000")
	featureLn.SetGeometry(linestr)

	yL, xL := GetTwoDimArrayCoordinates(featureLn, 1)

	if yL != 37.58800506591796 {
		t.Error("expected yL = 37.58800506591796, but yL = ", yP)
	}

	if xL != 55.74856460562653 {
		t.Error("expected xL = 55.74856460562653, but xL = ", xP)
	}
}

func TestNewMultiLineString(t *testing.T) {

	mLine := NewMultiLineString([][][]float64{
		{{37.58766174316406, 55.78178615666911},
			{37.545433044433594, 55.74740520331641},
			{37.57495880126953, 55.71589492312098}},

		{{37.63092041015624, 55.71492794801965},
			{37.61787414550781, 55.72769009202728},
			{37.64396667480469, 55.74817814201809},
			{37.69718170166015, 55.71879570480289}}})

	l := (mLine.Coordinates).([][][]float64)

	if l[0][1][0] != 37.545433044433594 {
		t.Error("expected l[0][1][0] =, but l[0][1][0] = ", l[0][1][0])
	}
	if l[0][1][1] != 55.74740520331641 {
		t.Error("expected l[0][1][1] =, but l[0][1][1] = ", l[0][1][1])
	}

	if l[1][1][0] != 37.61787414550781 {
		t.Error("expected l[1][1][0] =, but l[1][1][0] = ", l[1][1][0])
	}
	if l[1][1][1] != 55.72769009202728 {
		t.Error("expected l[1][1][1] =, but l[1][1][1] = ", l[1][1][1])
	}

}

func TestNewPolygon(t *testing.T) {

	polyg := NewPolygon([][][]float64{
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

	p := (polyg.Coordinates).([][][]float64)

	if p[0][1][0] != 37.63023376464844 {
		t.Error("Expected p[0][1]0[] = ,but p[0][1][0] = ", p[0][1][0])
	}
	if p[0][1][1] != 55.62993418221071 {
		t.Error("Expected p[0][1][1] = ,but p[0][1][1] = ", p[0][1][1])
	}

	if p[1][1][0] != 37.62062072753906 {
		t.Error("Expected p[][][0] = ,but p[1][1][0] = ", p[1][1][0])
	}
	if p[1][1][1] != 55.73870858770892 {
		t.Error("Expected p[1][1][1] = ,but p[1][1][1] = ", p[1][1][1])
	}

}

func TestGetThreeDimArrayCoordinates(t *testing.T) {

	mLine := NewMultiLineString([][][]float64{
		{{37.58766174316406, 55.78178615666911},
			{37.545433044433594, 55.74740520331641},
			{37.57495880126953, 55.71589492312098}},

		{{37.63092041015624, 55.71492794801965},
			{37.61787414550781, 55.72769009202728},
			{37.64396667480469, 55.74817814201809},
			{37.69718170166015, 55.71879570480289}}})

	featureML := NewFeature()
	featureML.SetProperty("MLineStr", "[][][]")
	featureML.SetID("000")
	featureML.SetGeometry(mLine)

	yML, xML := GetThreeDimArrayCoordinates(featureML, 0, 2)

	if yML != 37.57495880126953 {
		t.Error("Expected yML = 37.57495880126953, but yML = ", yML)
	}
	if xML != 55.71589492312098 {
		t.Error("Expected xML = 55.71589492312098, but xML = ", xML)
	}

	polyg := NewPolygon([][][]float64{
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

	featureP := NewFeature()
	featureP.SetProperty("Polyg", "[][][]")
	featureP.SetID("000")
	featureP.SetGeometry(polyg)

	yP, xP := GetThreeDimArrayCoordinates(featureP, 1, 2)

	if yP != 37.651519775390625 {
		t.Error("Expected yP = 37.651519775390625, but yP = ", yP)
	}
	if xP != 55.792017325495095 {
		t.Error("Expected xP = 55.792017325495095, but xP = ", xP)
	}

}

func TestNewMultiPolygon(t *testing.T) {

	multpol := NewMultiPolygon([][][][]float64{

		{
			{

				{37.55744934082031, 55.76189525593947},
				{37.56088256835937, 55.76150892439349},
				{37.596588134765625, 55.74373353535969},
				{37.58354187011719, 55.78622642787003},
				{37.55744934082031, 55.76189525593947},
			}},

		{
			{

				{38.08904439210892, 55.81190491362447},
				{38.08977395296096, 55.81190491362447},
				{38.08977395296096, 55.81212194465421},
				{38.08904439210892, 55.81212194465421},
				{38.08904439210892, 55.81190491362447}}}})

	mp := (multpol.Coordinates).([][][][]float64)

	if mp[0][0][1][0] != 37.56088256835937 {
		t.Error("Expected mp[0][0][1][0] = 37.56088256835937, but mp[0][0][1][0] = ", mp[0][0][1][0])
	}
	if mp[0][0][1][1] != 55.76150892439349 {
		t.Error("Expected mp[0][0][1][1] = 55.76150892439349, but mp[0][0][1][1] = ", mp[0][0][1][1])
	}

}

func TestGetFourDimArrayCoordinates(t *testing.T) {
	multpol := NewMultiPolygon([][][][]float64{

		{
			{

				{37.55744934082031, 55.76189525593947},
				{37.56088256835937, 55.76150892439349},
				{37.596588134765625, 55.74373353535969},
				{37.58354187011719, 55.78622642787003},
				{37.55744934082031, 55.76189525593947},
			}},

		{
			{

				{38.08904439210892, 55.81190491362447},
				{38.08977395296096, 55.81190491362447},
				{38.08977395296096, 55.81212194465421},
				{38.08904439210892, 55.81212194465421},
				{38.08904439210892, 55.81190491362447}}}})

	feature := NewFeature()
	feature.SetProperty("MPolyg", "[][][][]")
	feature.SetID("000")
	feature.SetGeometry(multpol)

	y, x := GetFourDimArrayCoordinates(feature, 1, 0, 3)

	if y != 38.08904439210892 {
		t.Error("Expecting y = 38.08904439210892, but y = ", y)
	}
	if x != 55.81212194465421 {
		t.Error("Expecting x = 55.81212194465421, but x = ", x)
	}
}

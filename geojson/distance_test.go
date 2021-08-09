package geogoth

import (
	"testing"
)

func TestDistance(t *testing.T) {

	ftr := &Ftrs

	CreateFtr1()
	CreateFtr2()
	CreateFtr3()
	CreateFtr4()
	CreateFtr5()
	CreateFtr6()
	CreateFtr7()
	CreateFtr8()
	CreateFtr9()
	CreateFtr10()
	CreateFtr11()
	CreateFtr12()
	CreateFtr13()
	CreateFtr14()
	CreateFtr15()
	CreateFtr16()
	CreateFtr17()
	CreateFtr18()
	CreateFtr19()
	CreateFtr20()
	CreateFtr21()
	// *****
	// *****
	// *****

	// Point - Point
	if Distance(ftr.Feature1, ftr.Feature2) != Distance(ftr.Feature2, ftr.Feature1) {
		t.Error("Expected Distance(ftr.Feature1, ftr.Feature2) == Distance(ftr.Feature2, ftr.Feature1)")
	}
	if int(Distance(ftr.Feature1, ftr.Feature2)) != 1639 {
		t.Error("Expected Distance(ftr.Feature1, ftr.Feature2)) == 1639")
	}
	// Point - MultiPoint
	if Distance(ftr.Feature1, ftr.Feature3) != Distance(ftr.Feature3, ftr.Feature1) {
	}
	if int(Distance(ftr.Feature1, ftr.Feature3)) != 1619 {
		t.Error("Expected Distance(ftr.Feature1, ftr.Feature3) == Distance(ftr.Feature3, ftr.Feature1)")
	}
	// Point - LineString
	if Distance(ftr.Feature1, ftr.Feature4) != Distance(ftr.Feature4, ftr.Feature1) {
		t.Error("Expected Distance(ftr.Feature1, ftr.Feature4) == Distance(ftr.Feature4, ftr.Feature1)")
	}
	if int(Distance(ftr.Feature1, ftr.Feature4)) != 459 {
		t.Error("Expected int(Distance(ftr.Feature1, ftr.Feature4)) == 459")
	}
	// Point - MultiLineString
	if Distance(ftr.Feature1, ftr.Feature5) != Distance(ftr.Feature5, ftr.Feature1) {
		t.Error("Expected Distance(ftr.Feature1, ftr.Feature5) == Distance(ftr.Feature5, ftr.Feature1)")
	}
	if int(Distance(ftr.Feature1, ftr.Feature5)) != 163 {
		t.Error("Expected int(Distance(ftr.Feature1, ftr.Feature5)) == 163")
	}
	// Point - Polygon
	if Distance(ftr.Feature1, ftr.Feature6) != Distance(ftr.Feature6, ftr.Feature1) {
		t.Error("Expected Distance(ftr.Feature1, ftr.Feature6) == Distance(ftr.Feature6, ftr.Feature1)")
	}
	if int(Distance(ftr.Feature1, ftr.Feature6)) != 315 {
		t.Error("int(Distance(ftr.Feature1, ftr.Feature6)) == 315")
	}
	// Point - MultiPolygon
	if Distance(ftr.Feature1, ftr.Feature7) != Distance(ftr.Feature7, ftr.Feature1) {
		t.Error("Expected Distance(ftr.Feature1, ftr.Feature7) == Distance(ftr.Feature7, ftr.Feature1)")
	}
	if int(Distance(ftr.Feature1, ftr.Feature7)) != 1445 {
		t.Error("Expected int(Distance(ftr.Feature1, ftr.Feature7)) == 1445")
	}

	// MultiPoint - MultiPoint
	if Distance(ftr.Feature3, ftr.Feature8) != Distance(ftr.Feature8, ftr.Feature3) {
		t.Error("Expected Distance(ftr.Feature3, ftr.Feature8) == Distance(ftr.Feature8, ftr.Feature3)")
	}
	if int(Distance(ftr.Feature3, ftr.Feature8)) != 2541 {
		t.Error("Expected int(Distance(ftr.Feature3, ftr.Feature8)) == 2541")
	}
	// MultiPoint - LineString
	if Distance(ftr.Feature3, ftr.Feature9) != Distance(ftr.Feature9, ftr.Feature3) {
		t.Error("Expected Distance(ftr.Feature3, ftr.Feature9) == Distance(ftr.Feature9, ftr.Feature3)")
	}
	if int(Distance(ftr.Feature3, ftr.Feature9)) != 2410 {
		t.Error("Expected int(Distance(ftr.Feature3, ftr.Feature9)) == 2410")
	}
	// MultiPoint - MultiLineString
	if Distance(ftr.Feature3, ftr.Feature5) != Distance(ftr.Feature5, ftr.Feature3) {
		t.Error("Expected Distance(ftr.Feature3, ftr.Feature5) == Distance(ftr.Feature5, ftr.Feature3)")
	}
	if int(Distance(ftr.Feature3, ftr.Feature5)) != 736 {
		t.Error("Expected int(Distance(ftr.Feature3, ftr.Feature5)) == 736")
	}
	// MultiPoint - Polygon
	if Distance(ftr.Feature10, ftr.Feature11) != Distance(ftr.Feature11, ftr.Feature10) {
		t.Error("Expected Distance(ftr.Feature10, ftr.Feature11) == Distance(ftr.Feature11, ftr.Feature10)")
	}
	if int(Distance(ftr.Feature10, ftr.Feature11)) != 3783 {
		t.Error("Expected int(Distance(ftr.Feature10, ftr.Feature11)) == 3783")
	}
	// MultiPoint - MultiPolygon
	if Distance(ftr.Feature12, ftr.Feature13) != Distance(ftr.Feature13, ftr.Feature12) {
		t.Error("Expected Distance(ftr.Feature12, ftr.Feature13) == Distance(ftr.Feature13, ftr.Feature12)")
	}
	if int(Distance(ftr.Feature12, ftr.Feature13)) != 61 {
		t.Error("Expected int(Distance(ftr.Feature12, ftr.Feature13)) == 61")
	}

	// LineString - LineString
	if Distance(ftr.Feature14, ftr.Feature15) != Distance(ftr.Feature15, ftr.Feature14) {
		t.Error("Expected Distance(ftr.Feature14, ftr.Feature15) == Distance(ftr.Feature15, ftr.Feature14)")
	}
	if int(Distance(ftr.Feature14, ftr.Feature15)) != 144703 {
		t.Error("Expected int(Distance(ftr.Feature14, ftr.Feature15)) == 144703")
	}

	// LineString - MultiLineString
	if Distance(ftr.Feature14, ftr.Feature5) != Distance(ftr.Feature5, ftr.Feature14) {
		t.Error("Expected Distance(ftr.Feature14, ftr.Feature5) == Distance(ftr.Feature5, ftr.Feature14)")
	}
	if int(Distance(ftr.Feature14, ftr.Feature5)) != 84091 {
		t.Error("Expected int(Distance(ftr.Feature14, ftr.Feature5)) == 84091")
	}
	// LineString - Polygon
	if Distance(ftr.Feature14, ftr.Feature6) != Distance(ftr.Feature6, ftr.Feature14) {
		t.Error("Expected Distance(ftr.Feature14, ftr.Feature6) == Distance(ftr.Feature6, ftr.Feature14)")
	}
	if int(Distance(ftr.Feature14, ftr.Feature6)) != 83203 {
		t.Error("Expected int(Distance(ftr.Feature14, ftr.Feature6)) == 83203")
	}
	// LineString - MultiPolygon
	if Distance(ftr.Feature13, ftr.Feature16) != Distance(ftr.Feature16, ftr.Feature13) {
		t.Error("Expected  Distance(ftr.Feature13, ftr.Feature16) == Distance(ftr.Feature16, ftr.Feature13)s")
	}
	if int(Distance(ftr.Feature13, ftr.Feature16)) != 1000 {
		t.Error("Expected int(Distance(ftr.Feature13, ftr.Feature16)) == 1000 ")
	}

	// MultiLineString - MultiLineString
	if Distance(ftr.Feature17, ftr.Feature18) != Distance(ftr.Feature18, ftr.Feature17) {
		t.Error("Expected Distance(ftr.Feature17, ftr.Feature18) == Distance(ftr.Feature18, ftr.Feature17)")
	}
	if int(Distance(ftr.Feature17, ftr.Feature18)) != 3005 {
		t.Error("Expected int(Distance(ftr.Feature17, ftr.Feature18)) == 3005")
	}
	// MultiLineString - Polygon
	if Distance(ftr.Feature6, ftr.Feature18) != Distance(ftr.Feature18, ftr.Feature6) {
		t.Error("Expected Distance(ftr.Feature6, ftr.Feature18) == Distance(ftr.Feature18, ftr.Feature6)")
	}
	if int(Distance(ftr.Feature6, ftr.Feature18)) != 3638 {
		t.Error("Expected int(Distance(ftr.Feature6, ftr.Feature18)) == 3638")
	}
	// MultiLineString - MultiPolygon
	if Distance(ftr.Feature7, ftr.Feature18) != Distance(ftr.Feature18, ftr.Feature7) {
		t.Error("Expected Distance(ftr.Feature7, ftr.Feature18) == Distance(ftr.Feature18, ftr.Feature7)")
	}
	if int(Distance(ftr.Feature7, ftr.Feature18)) != 3444 {
		t.Error("Expected int(Distance(ftr.Feature7, ftr.Feature18)) == 3444")
	}

	// Polygon - Polygon
	if Distance(ftr.Feature10, ftr.Feature19) != Distance(ftr.Feature19, ftr.Feature10) {
		t.Error("Expected Distance(ftr.Feature10, ftr.Feature19) == Distance(ftr.Feature19, ftr.Feature10)")
	}
	if int(Distance(ftr.Feature10, ftr.Feature19)) != 7401 {
		t.Error("Expected int(Distance(ftr.Feature10, ftr.Feature19)) == 7401")
	}
	// Polygon - MultiPolygon

	// // MultiPolygon - MultiPolygon

}

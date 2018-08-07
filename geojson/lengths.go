package geogoth

// LineStringLength counts lenght of LineString
func LineStringLength(feature *Feature) float64 {
	var length float64

	coords := (feature.Geom.Coordinates).([][]float64) // Convert interface to [][][]float64
	lineCoords := make([][]float64, 0)                 // Creates slice for coords of one line

	for i := range coords {
		y, x := GetTwoDimArrayCoordinates(feature, i) // Coordinates of LineString
		lineCoords = append(lineCoords, []float64{y, x})
	}

	for i := 0; i < len(lineCoords)-1; i++ {

		lengthTmp := DistancePointPointDeg(lineCoords[i][0], lineCoords[i][1], lineCoords[i+1][0], lineCoords[i+1][1])
		length = length + lengthTmp
	}

	return length
}

// MultiLineStringLength counts lenght of MultiLineString
func MultiLineStringLength(feature *Feature) float64 {
	var length float64

	lineCoords := make([][]float64, 0)                        // Creates slice for coords of one line
	mlineCoords := make([][][]float64, 0)                     // Creates slice for coords of the MultiLineString
	multlinestr := (feature.Geom.Coordinates).([][][]float64) // Convert interface to [][][]float64

	for i := range multlinestr { // Finds coords of the MultiLineString
		for j := range multlinestr[i] {
			y, x := GetThreeDimArrayCoordinates(feature, i, j)
			lineCoords = append(lineCoords, []float64{y, x})
		}
		mlineCoords = append(mlineCoords, lineCoords)
		lineCoords = nil // empty slice

	}

	for i := range mlineCoords {

		for j := 0; j < len(mlineCoords[i])-1; j++ {

			lengthTmp := DistancePointPointDeg(mlineCoords[i][j][0], mlineCoords[i][j][1], mlineCoords[i][j+1][0], mlineCoords[i][j+1][1])
			length = length + lengthTmp
		}

	}

	return length
}

// PolygonLength counts lenght of Polygon
func PolygonLength(feature *Feature) float64 {
	var length float64

	lineCoords := make([][]float64, 0)                    // Creates slice for coords of one line
	polygCoords := make([][][]float64, 0)                 // Creates slice for coords of the MultiLineString
	polygon := (feature.Geom.Coordinates).([][][]float64) // Convert interface to [][][]float64

	for i := range polygon { // Finds coords of the MultiLineString
		for j := range polygon[i] {
			y, x := GetThreeDimArrayCoordinates(feature, i, j)
			lineCoords = append(lineCoords, []float64{y, x})
		}
		polygCoords = append(polygCoords, lineCoords)
		lineCoords = nil // empty slice

	}

	for i := range polygCoords {

		for j := 0; j < len(polygCoords[i])-1; j++ {

			lengthTmp := DistancePointPointDeg(polygCoords[i][j][0], polygCoords[i][j][1], polygCoords[i][j+1][0], polygCoords[i][j+1][1])
			length = length + lengthTmp
		}

	}

	return length
}

// MultipolygonLength counts lenght of MultipolygonLength
func MultipolygonLength(feature *Feature) float64 {
	var length float64

	mpolyg := (feature.Geom.Coordinates).([][][][]float64)

	mpolygCoords := make([][][]float64, 0) // Creates slice for coords of the MultiPolygon
	mlineCoords := make([][]float64, 0)    // Creates slice for coords of one line

	for m := range mpolyg { // Finds coords of MultiPolygon
		for p := range mpolyg[m] {
			for i := range mpolyg[m][p] {

				y, x := GetFourDimArrayCoordinates(feature, m, p, i)
				mlineCoords = append(mlineCoords, []float64{y, x})

			}
			mpolygCoords = append(mpolygCoords, mlineCoords)
			mlineCoords = nil // empty slice

		}
	}

	for i := range mpolygCoords {
		for j := 0; j < len(mpolygCoords[i])-1; j++ {

			lengthTmp := DistancePointPointDeg(mpolygCoords[i][j][0], mpolygCoords[i][j][1], mpolygCoords[i][j+1][0], mpolygCoords[i][j+1][1])
			length = length + lengthTmp
		}
	}

	return length

}

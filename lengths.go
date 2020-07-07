package geogoth

// LineStringLength counts length of LineString
func LineStringLength(linestr *LineString) float64 {
	var length float64

	coords := linestr.Coords
	lineCoords := make([][]float64, 0) // Creates slice for coords of one line

	for i := range coords {
		y, x := linestr.Coords[i][0], linestr.Coords[i][1]
		lineCoords = append(lineCoords, []float64{y, x})
	}

	for i := 0; i < len(lineCoords)-1; i++ {

		lengthTmp := DistancePointPointDeg(lineCoords[i][0], lineCoords[i][1], lineCoords[i+1][0], lineCoords[i+1][1])
		length = length + lengthTmp
	}

	return length
}

// MultiLineStringLength counts length of MultiLineString
func MultiLineStringLength(mlineStr MultiLineString) float64 {
	var length float64

	lineCoords := make([][]float64, 0)    // Creates slice for coords of one line
	mlineCoords := make([][][]float64, 0) // Creates slice for coords of the MultiLineString
	multlinestr := (mlineStr.Coords)      // Convert interface to [][][]float64

	for i := range multlinestr { // Finds coords of the MultiLineString
		for j := range multlinestr[i] {
			y, x := multlinestr[i][j][0], multlinestr[i][j][1]
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

// PolygonLength counts length of Polygon
func PolygonLength(polygon Polygon) float64 {
	var length float64

	lineCoords := make([][]float64, 0)  // Creates slice for coords of one line
	polCoords := make([][][]float64, 0) // Creates slice for coords of the MultiLineString
	polygonCoords := polygon.Coords     // Convert interface to [][][]float64

	for i := range polygonCoords { // Finds coords of the MultiLineString
		for j := range polygonCoords[i] {
			y, x := polygon.Coords[i][j][0], polygon.Coords[i][j][1]
			lineCoords = append(lineCoords, []float64{y, x})
		}
		polCoords = append(polCoords, lineCoords)
		lineCoords = nil // empty slice

	}

	for i := range polCoords {

		for j := 0; j < len(polCoords[i])-1; j++ {

			lengthTmp := DistancePointPointDeg(polCoords[i][j][0], polCoords[i][j][1], polCoords[i][j+1][0], polCoords[i][j+1][1])
			length = length + lengthTmp
		}

	}

	return length
}

// MultipolygonLength counts length of MultipolygonLength
func MultipolygonLength(mpolygon MultiPolygon) float64 {
	var length float64

	mpolyg := mpolygon.Coords

	mpolygCoords := make([][][]float64, 0) // Creates slice for coords of the MultiPolygon
	mlineCoords := make([][]float64, 0)    // Creates slice for coords of one line

	for m := range mpolyg { // Finds coords of MultiPolygon
		for p := range mpolyg[m] {
			for i := range mpolyg[m][p] {

				y, x := mpolygon.Coords[m][p][i][0], mpolygon.Coords[m][p][i][1]
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

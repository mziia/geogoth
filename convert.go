package geogoth

import (
	"fmt"
	"math"
)

// def degrees2meters(lon, lat):

//     x = lon * 20037508.34 / 180
//     y = math.log(math.tan((90 + lat) * math.pi / 360)) / (math.pi / 180)
//     y = y * 20037508.34 / 180
//     return (x, y)

// формула на питоне для перевода градусов в метры (сфера на плоскость для Земли)

// DegreesToMeters converte degrees to meters (spheere to flat)
func DegreesToMeters(feature *Feature) (float64, float64) {

	geom := feature.Geom
	var lon float64
	var lat float64
	// fmt.Println("geom.Coordinates: ", geom.Coordinates)
	fmt.Println()
	switch geom.Type {

	case "Point":
		coord := (geom.Coordinates).([]float64)
		lon = coord[0]
		lat = coord[1]

	case "MultiPoint": // Доделать: куда то помещать все x и y
		// coord := (geom.Coordinates).([][]float64)
		// 	fmt.Println(coord)
		// 	for i := range coord {
		// 		lon = coord[i][0]
		// 		lat = coord[i][1]
		// 	}
	case "LineString":
		// coord := (geom.Coordinates).([][]float64)
	case "MultiLineString":
		// coord := (geom.Coordinates).([][][]float64)
	case "Polygon":
		// coord := (geom.Coordinates).([][][]float64)
	case "MultiPolygon":
		// coord := (geom.Coordinates).([][][][]float64)

	}
	//lon - y
	//lat - x
	x := lon * 20037508.34 / 180
	y := math.Log(math.Tan((90+lat)*math.Pi/360)) / (math.Pi / 180)
	y = y * 20037508.34 / 180
	return x, y

}

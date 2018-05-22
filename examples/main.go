package main

import (
	"encoding/json"
	"fmt"

	geogoth "github.com/mziia/geogoth"
)

func main() {

	collection := geogoth.NewFeatureCollection() // GeoJson

	point := geogoth.NewPoint([]float64{125.6, 10.1}) // Point
	feature1 := geogoth.NewFeature()                  // Feature

	feature1.SetProperty("id_num", 1)

	feature1.SetID("0001")
	feature1.SetGeometry(point)
	collection.AddFeature(feature1)

	newls := geogoth.NewLineString([][]float64{{52.370725881211314, 4.889259338378906}, {52.3711451105601, 4.895267486572266}})

	feature2 := geogoth.NewFeature()
	feature1.SetProperty("id_num", 1)

	feature2.SetProperty("coordin", [][]float64{{52.370725881211314, 4.889259338378906}, {52.3711451105601, 4.895267486572266}})
	feature2.SetID("0002")
	feature2.SetGeometry(newls)
	collection.AddFeature(feature2)

	m, _ := json.Marshal(collection)
	fmt.Println(string(m))

	fmt.Println(geogoth.GetGeoType(feature1))
	fmt.Println(string(geogoth.GetGeoType(collection.Features[1])))

}

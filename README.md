# geogoth

> GeoGoth is a package for creating geojson using golang language  <img src="https://user-images.githubusercontent.com/24193681/41141904-75a44f22-6afc-11e8-83c0-5ee152d600e3.png" width="250" align="right"> 


## Contents
- [Installation](#installation)
- [Quick start](#quick-start)




## Installation 

Download and install:
```sh
$ go get github.com/mziia/geogoth
```

Import: 
```go
import "github.com/mziia/geogoth"
```

## Quick start

```go
package main

import (
	"fmt"

	"github.com/mziia/geogoth"
)

func main() {
	collection := geogoth.NewFeatureCollection()

	point1 := geogoth.NewPoint([]float64{37.6175, 55.752})  // lon, lat
	point2 := geogoth.NewPoint([]float64{37.6048, 55.7649}) // y, x

	feature1 := geogoth.NewFeature()
	feature1.SetProperty("локация", "Кремль")
	feature1.SetID("0001")
	feature1.SetGeometry(point1)

	feature2 := geogoth.NewFeature()
	feature2.SetProperty("локация", "Тверская")
	feature2.SetID("0002")
	feature2.SetGeometry(point2)

	collection.AddFeature(feature1)
	collection.AddFeature(feature2)

	// The order of parameters' transfer does not matter
	fmt.Println("Distance between Points (feature1 - feature2): ", geogoth.Distance(feature1, feature2))
	fmt.Println("Distance between Points (feature2 - feature1): ", geogoth.Distance(feature2, feature1))

}
```

```sh
Distance between Points (feature1 - feature2):  1639.8005076177767
Distance between Points (feature2 - feature1):  1639.8005076177767

```

You can transfer to Distance() any kind of geometry (Point, MultiPoint, LineString, Polygon etc):


```go

package main

import (
	"fmt"

	"github.com/mziia/geogoth"
)

func main() {
	mu1 := geogoth.NewMultiPolygon([][][][]float64{

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

				[]float64{37.63298034667969, 55.789121984291626},
				[]float64{37.61821746826172, 55.77135918323877},
				[]float64{37.68241882324219, 55.773097205876766},
				[]float64{37.677955627441406, 55.82404473410693},
				[]float64{37.63298034667969, 55.789121984291626}}}})

	feature1 := geogoth.NewFeature()
	feature1.SetProperty("MPolyg", "[][][]")
	feature1.SetID("0001")
	feature1.SetGeometry(mu1)

	mu2 := geogoth.NewMultiPolygon([][][][]float64{

		[][][]float64{
			[][]float64{

				[]float64{37.62542724609375, 55.70409619139846},
				[]float64{37.680015563964844, 55.70312892423751},
				[]float64{37.68241882324219, 55.74605252374095},
				[]float64{37.62542724609375, 55.70409619139846},
			}},

		[][][]float64{
			[][]float64{

				[]float64{37.513160705566406, 55.82250184886082},
				[]float64{37.540283203125, 55.80205284218845},
				[]float64{37.60379791259765, 55.81652259056771},
				[]float64{37.58422851562499, 55.840434111266205},
				[]float64{37.513160705566406, 55.82250184886082}}}})

	feature2 := geogoth.NewFeature()
	feature2.SetProperty("MPolyg", "[][][]")
	feature2.SetID("0002")
	feature2.SetGeometry(mu2)

	fmt.Println("Distance MultiPolygon -  MultiPolygon ( feature1, feature2 ):  ", geogoth.Distance(feature1, feature2))

}

```

```sh

Distance MultiPolygon -  MultiPolygon ( feature1, feature2 ):   2647.108678983299


```





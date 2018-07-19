# geogoth

> GeoGoth is a package for creating geojson using golang language  <img src="https://user-images.githubusercontent.com/24193681/41141904-75a44f22-6afc-11e8-83c0-5ee152d600e3.png" width="250" align="right"> 



## Installation and Usage 

Download and install:
```sh
$ go get github.com/mziia/geogoth
```

Import: 
```go
import "github.com/mziia/geogoth"
```

## Quick start

Create GeoJSON:

```go
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
```


Find distance between Features: 

```go
// The order of parameters' transfer does not matter
distance := geogoth.Distance(feature1, feature2)
OR
distance := geogoth.Distance(feature2, feature1)

```


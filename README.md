# geogoth

> GeoGoth is a package for creating geojson using golang language  <img src="https://user-images.githubusercontent.com/24193681/41141904-75a44f22-6afc-11e8-83c0-5ee152d600e3.png" width="250" align="right"> 



## Installation and Usage 

Download and install:

> go get github.com/mziia/geogoth

Import: 

> import "github.com/mziia/geogoth"


## Quick start

```
collection := features.NewFeatureCollection() 

point := coordinates.NewPoint([]float64{125.6, 10.1}) 

feature := features.NewFeature()  
                    
feature.SetProperty("id_num", 1)
feature.SetID("0001")
feature.SetGeometry(point)
	
collection.AddFeature(feature)
```

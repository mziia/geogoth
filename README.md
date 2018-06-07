# geogoth

> GeoGoth is a package for creating geojson using golang language  <img src="https://user-images.githubusercontent.com/24193681/41080096-79aa12ac-6a2d-11e8-8cae-73f826b20d34.png" width="250" align="right"> 


## go get
> go get github.com/mziia/geogoth

## import
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

# geogoth

> <img src="https://user-images.githubusercontent.com/24193681/40360416-d803582e-5dce-11e8-9998-d01980f0affa.png" width="250"> GeoGoth is a package for creating geojson using golang language 

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

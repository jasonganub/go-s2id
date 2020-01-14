package main

import (
	//"encoding/json"
	"fmt"
	"github.com/paulmach/go.geojson"
)
//import "github.com/golang/geo/s2"
//import "io/ioutil"

//func LoopsFromFeatureCollection() {
//
//}

func LoadGeojson() {
	rawFeatureJSON := []byte(`{
	 "type": "FeatureCollection",
	 "features": [
		{
		  "type": "Feature",
		  "properties": {},
		  "geometry": {
			"type": "Polygon",
			"coordinates": [
			  [
				[
				  98.87701749801636,
				  3.633157634191541
				],
				[
				  98.87803137302399,
				  3.632113674421417
				],
				[
				  98.87891113758087,
				  3.6329327814200156
				],
				[
				  98.87787580490112,
				  3.6339232039220333
				],
				[
				  98.87701749801636,
				  3.633157634191541
				]
			  ]
			]
		  }
		}
	 ]
	}`)

	fc, err := geojson.UnmarshalFeatureCollection(rawFeatureJSON)

	fmt.Println("Coordinates Type: " + fc.Features[0].Geometry.Type)
	fmt.Println(fc.Features[0].Geometry.Polygon[0][0][0])
	fmt.Println(fc.Features[0].Geometry.Polygon[0][0][1])

	//
	//// Geometry
	//rawGeometryJSON := []byte(`{"type": "Point", "coordinates": [102.0, 0.5]}`)
	//g, err := geojson.UnmarshalFeatureCollection(rawGeometryJSON)
	//
	//println(g.IsPoint())
	//println(g.Point[0])




	if err != nil {
		println(err)
	}



}

func main() {
	LoadGeojson()
}
package main

import (
	"fmt"
	"github.com/golang/geo/s2"
	geojson "github.com/paulmach/go.geojson"
	//s2 "github.com/golang/geo/s2"
)
//import "io/ioutil"


func LoadPolygon() *s2.Loop {
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
	if err != nil {
		println(err)
	}

	fmt.Println("Coordinates Type: " + fc.Features[0].Geometry.Type)
	fmt.Println(fc.Features[0].Geometry.Polygon[0][0][0])
	fmt.Println(fc.Features[0].Geometry.Polygon[0][0][1])

	coordinates := fc.Features[0].Geometry.Polygon[0]

	var points []s2.Point
	for _, coordinate := range coordinates {
		s2ll := s2.LatLngFromDegrees(coordinate[0], coordinate[1])
		point := s2.PointFromLatLng(s2ll)
		points = append(points, point)
	}

	return s2.LoopFromPoints(points)
}

func GetOuterCovering(loops *s2.Loop) {
	//rc := s2.RegionCoverer{
	//	MaxCells: 100,
	//	MinLevel: 5,
	//	MaxLevel: 10,
	//}
	//
	//var coverings []s2.RegionCoverer
	for _, loop := range loops.Vertices() {
		fmt.Println(loop)
	}

}

func main() {
	polygon := LoadPolygon()

	fmt.Println(polygon)

	GetOuterCovering(polygon)
}
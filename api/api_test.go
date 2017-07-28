package api

import (
	"fmt"
	"testing"
)

func TestArtistSetlistsByMBID(t *testing.T) {
	res, _ := ArtistSetlistsByMBID("65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab", 0)
	fmt.Println(res)
}

func TestCityByGeoID(t *testing.T) {
	res, _ := CityByGeoID("5357527")
	fmt.Println(res)
}


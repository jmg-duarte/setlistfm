package api

import (
	"fmt"
	"testing"
)

func TestArtistSetlistsByMBID(t *testing.T) {
	res, _ := ArtistSetlistsByMBID("65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab", 2)
	fmt.Println(res)
}

func TestArtistByMBID(t *testing.T) {
	res, _ := ArtistByMBID("65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab")
	fmt.Println(res)
}

func TestCityByGeoID(t *testing.T) {
	res, _ := CityByGeoID("5357527")
	fmt.Println(res)
}

func TestSearchForArtists(t *testing.T) {
	res, _ := SearchForArtists("", "Metallica", "", 0)
	fmt.Println(res)
	res, _ = SearchForArtists("65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab", "", "", 0)
	fmt.Println(res)
}

func TestSearchForCities(t *testing.T) {
	res, _ := SearchForCities("Portugal", "", "", "", 0)
	fmt.Println(res)
	res, _ = SearchForCities("", "Pasadena", "", "", 0)
	fmt.Println(res)
}

func TestListAllCountries(t *testing.T) {
	res, _ := ListAllCountries()
	fmt.Println(res)
}

package setlistfm

import (
	"fmt"
	"testing"
)

func TestArtistSetlistsByMBID(t *testing.T) {
	c := NewClient("")
	res, _ := c.ArtistSetlistsByMBID("65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab", 2)
	fmt.Println(res)
}

func TestArtistByMBID(t *testing.T) {
	c := NewClient("")
	res, _ := c.ArtistByMBID("65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab")
	fmt.Println(res)
}

func TestCityByGeoID(t *testing.T) {
	c := NewClient("")
	res, _ := c.CityByGeoID("5357527")
	fmt.Println(res)
}

func TestSearchForArtists(t *testing.T) {
	c := NewClient("")
	res, _ := c.SearchForArtists(ArtistsQuery{ArtistName: "Metallica"})
	fmt.Println(res)
	//res, _ = SearchForArtists("65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab", "", "", 0)
	//fmt.Println(res)
}

func TestSearchForCities(t *testing.T) {
	c := NewClient("")
	res, _ := c.SearchForCities(CityQuery{Name: "Lisbon"})
	fmt.Println(res)
}

func TestListAllCountries(t *testing.T) {
	c := NewClient("")
	res, _ := c.ListAllCountries()
	fmt.Println(res)
}

func TestSearchForSetlists(t *testing.T) {
	c := NewClient("")
	res, _ := c.SearchForSetlists(SetlistQuery{ArtistName: "Opeth"})
	fmt.Println(res)
}

func TestSearchForVenues(t *testing.T) {
	c := NewClient("")
	res, _ := c.SearchForVenues(VenueQuery{CountryCode: "PT"})
	fmt.Println(res)
}

func TestSetlistByVersionID(t *testing.T) {
	c := NewClient("")
	res, _ := c.SetlistByVersionID("7be1aaa0")
	fmt.Println(res)
}

func TestSetlistByID(t *testing.T) {
	c := NewClient("")
	res, _ := c.SetlistByID("63de4613")
	fmt.Println(res)
}

func TestUserByID(t *testing.T) {
	c := NewClient("")
	res, _ := c.UserByID("ExecutiveChimp")
	fmt.Println(res)
}

func TestUserAttendedConcerts(t *testing.T) {
	c := NewClient("")
	res, _ := c.UserAttendedConcerts("ExecutiveChimp", 0)
	fmt.Println(res)
}

func TestUserEditedSetlists(t *testing.T) {
	c := NewClient("")
	res, _ := c.UserEditedSetlists("ExecutiveChimp", 0)
	fmt.Println(res)
}

func TestVenueByID(t *testing.T) {
	c := NewClient("")
	res, _ := c.VenueByID("73d466b5")
	fmt.Println(res)
}

func TestVenueSetlists(t *testing.T) {
	c := NewClient("")
	res, _ := c.VenueSetlists("73d466b5", 0)
	fmt.Println(res)
}

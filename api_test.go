package setlistfm

import (
	"context"
	"fmt"
	"testing"
)

func TestArtistSetlistsByMBID(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.ArtistSetlistsByMBID(context.TODO(), "65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab", 2)
	fmt.Println(res)
}

func TestArtistByMBID(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.ArtistByMBID(context.TODO(), "65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab")
	fmt.Println(res)
}

func TestCityByGeoID(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.CityByGeoID(context.TODO(), "5357527")
	fmt.Println(res)
}

func TestSearchForArtists(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.SearchForArtists(context.TODO(), ArtistsQuery{ArtistName: "Metallica"})
	fmt.Println(res)
	//res, _ = SearchForArtists("65f4f0c5-ef9e-490c-aee3-909e7ae6b2ab", "", "", 0)
	//fmt.Println(res)
}

func TestSearchForCities(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.SearchForCities(context.TODO(), CityQuery{CountryCode: "PT" ,Name: "Lisbon"})
	fmt.Println(res)
}

func TestListAllCountries(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.ListAllCountries(context.TODO())
	fmt.Println(res)
}

func TestSearchForSetlists(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.SearchForSetlists(context.TODO(), SetlistQuery{ArtistName: "Opeth"})
	fmt.Println(res)
}

func TestSearchForVenues(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.SearchForVenues(context.TODO(), VenueQuery{CountryCode: "PT"})
	fmt.Println(res)
}

func TestSetlistByVersionID(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.SetlistByVersionID(context.TODO(), "7be1aaa0")
	fmt.Println(res)
}

func TestSetlistByID(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.SetlistByID(context.TODO(), "63de4613")
	fmt.Println(res)
}

func TestUserByID(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.UserByID(context.TODO(), "ExecutiveChimp")
	fmt.Println(res)
}

func TestUserAttendedConcerts(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.UserAttendedConcerts(context.TODO(), "ExecutiveChimp", 0)
	fmt.Println(res)
}

func TestUserEditedSetlists(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.UserEditedSetlists(context.TODO(), "ExecutiveChimp", 0)
	fmt.Println(res)
}

func TestVenueByID(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.VenueByID(context.TODO(), "73d466b5")
	fmt.Println(res)
}

func TestVenueSetlists(t *testing.T) {
	c := NewClient("e1d7b2e5-c3b7-459b-9597-cf5fd9296182")
	res, _ := c.VenueSetlists(context.TODO(), "73d466b5", 0)
	fmt.Println(res)
}

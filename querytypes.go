package setlistfm

import (
	"net/http"
	"strconv"
)

// Query - Defines a Query type build to simplify the methods querying the API
type Query interface {
	// AddQuery
	AddQuery(r http.Request) *http.Request
}

// ArtistsQuery - Used by SearchForArtists
// If Page < 1 then the value will not be passed to the request and the API's default (1) will be used
// If the Artist information doesn't match you'll most likely won't receive a result
// e.g. If you search for the name "Opeth" and the "Mbid" passed is from Celine Dion
// Tip: Only use one at a time
type ArtistsQuery struct {
	ArtistMbid string
	ArtistName string
	ArtistTmid int
	Page       int
}

// AddQuery - Takes an http.Request and adds the Receivers query
func (a ArtistsQuery) AddQuery(r http.Request) *http.Request {
	q := r.URL.Query()

	q.Add("artistMbid", a.ArtistMbid)
	q.Add("artistName", a.ArtistName)
	q.Add("artistTmid", strconv.Itoa(a.ArtistTmid))
	// Consider changing page to string
	if a.Page > 0 {
		q.Add("p", strconv.Itoa(a.Page))
	}
	r.URL.RawQuery = q.Encode()
	return &r
}

// CityQuery - Used by SearchForCities
// If Page < 1 then the value will not be passed to the request and the API's default (1) will be used
// As in other queries, the fields cannot be exclusive (they should make sense)
// The CountryCode is represented like:
// "PT" - Portugal
// In case of doubt consult the API documentation - https://api.setlist.fm/docs/1.0/
// Or you can make use of this awesome wrapper and check with ListAllCountries
type CityQuery struct {
	CountryCode string
	Name        string
	State       string
	StateCode   string
	Page        int
}

// AddQuery - Takes an http.Request and adds the Receivers query
func (c CityQuery) AddQuery(r http.Request) *http.Request {
	q := r.URL.Query()

	q.Add("country", c.CountryCode)
	q.Add("name", c.Name)
	q.Add("state", c.State)
	q.Add("stateCode", c.StateCode)
	// Consider changing page to string
	if c.Page > 0 {
		q.Add("p", strconv.Itoa(c.Page))
	}

	r.URL.RawQuery = q.Encode()

	return &r
}

// SetlistQuery - Used by SearchForSetlists
// If Page < 1 then the value will not be passed to the request and the API's default (1) will be used
// Multiple uses of the LastUpdated field may change results
type SetlistQuery struct {
	ArtistMbid  string
	ArtistName  string
	ArtistTmid  int
	CityID      string
	CityName    string
	CountryCode string
	Date        Date
	LastUpdated DateTime
	State       string
	StateCode   string
	TourName    string
	VenueID     string
	VenueName   string
	Year        int
	Page        int
}

// AddQuery - Takes an http.Request and adds the Receivers query
func (s SetlistQuery) AddQuery(r http.Request) *http.Request {
	q := r.URL.Query()
	q.Add("artistMbid", s.ArtistMbid)
	q.Add("artistName", s.ArtistName)
	q.Add("cityId", s.CityID)
	q.Add("cityName", s.CityName)
	q.Add("countryCode", s.CountryCode)
	q.Add("date", s.Date.String())
	q.Add("lastUpdated", s.LastUpdated.String())
	q.Add("state", s.State)
	q.Add("stateCode", s.StateCode)
	q.Add("tourName", s.TourName)
	q.Add("venueId", s.VenueID)
	q.Add("venueName", s.VenueName)

	if s.ArtistTmid != 0 {
		q.Add("artistTmid", strconv.Itoa(s.ArtistTmid))
	}
	if s.Year != 0 {
		q.Add("year", strconv.Itoa(s.Year))
	}

	if s.Page > 0 {
		q.Add("p", strconv.Itoa(s.Page))
	}

	r.URL.RawQuery = q.Encode()

	return &r
}

// VenueQuery - Used by SearchForVenues
// If Page < 1 then the value will not be passed to the request and the API's default (1) will be used
type VenueQuery struct {
	CityID      string
	CityName    string
	CountryCode string
	VenueName   string
	State       string
	StateCode   string
	Page        int
}

// AddQuery - Takes an http.Request and adds the Receivers query
func (v VenueQuery) AddQuery(r http.Request) *http.Request {
	q := r.URL.Query()
	q.Add("cityId", v.CityID)
	q.Add("cityName", v.CityName)
	q.Add("country", v.CountryCode)
	q.Add("name", v.VenueName)
	q.Add("state", v.State)
	q.Add("stateCode", v.StateCode)

	if v.Page > 0 {
		q.Add("p", strconv.Itoa(v.Page))
	}
	r.URL.RawQuery = q.Encode()

	return &r
}

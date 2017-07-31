package setlistfm

import (
	"net/http"
	"strconv"
)

type Query interface {
	BuildQuery(r http.Request) *http.Request
}

type ArtistsQuery struct {
	ArtistMbid string
	ArtistName string
	ArtistTmid int
	Page       int
}

func (a ArtistsQuery) BuildQuery(r http.Request) *http.Request {
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

type CityQuery struct {
	Country   string
	Name      string
	State     string
	StateCode string
	Page      int
}

func (c CityQuery) BuildQuery(r http.Request) *http.Request {
	q := r.URL.Query()

	q.Add("country", c.Country)
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

func (s SetlistQuery) BuildQuery(r http.Request) *http.Request {
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

type VenueQuery struct {
	CityID      string
	CityName    string
	CountryCode string
	VenueName   string
	State       string
	StateCode   string
	Page        int
}

func (v VenueQuery) BuildQuery(r http.Request) *http.Request {
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

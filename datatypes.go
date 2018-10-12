package setlistfm

import (
	"encoding/json"

	"fmt"
	"time"
)

// Artist - This class represents an artist.
// An artist is a musician or a group of musicians.
// Each artist has a definite Musicbrainz Identifier (MBID)
// with which the artist can be uniquely identified.
type Artist struct {
	MBID           string `json:"mbid"`
	TMID           int    `json:"tmid"`
	Name           string `json:"name"`
	SortName       string `json:"sortName"`
	Disambiguation string `json:"disambiguation"`
	URL            string `json:"url"`
}

func (a Artist) String() string {
	jsonString, err := json.Marshal(a)
	if err != nil {
		panic("Artist.String()")
	}
	return string(jsonString)
}

// Artists - A Result consisting of a list of artists.
type Artists struct {
	Type         string   `json:"type,omitempty"`
	Artists      []Artist `json:"artist"`
	Total        int      `json:"total"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
}

func (a Artists) String() string {
	jsonString, err := json.Marshal(a)
	if err != nil {
		panic("Artists.String()")
	}
	return string(jsonString)
}

// Coordinates - Coordinates of a point on the globe. Mostly used for Cities.
type Coordinates struct {
	Longitude float32 `json:"long"`
	Latitude  float32 `json:"lat"`
}

func (c Coordinates) String() string {
	jsonString, err := json.Marshal(c)
	if err != nil {
		panic("Coordinates.String()")
	}
	return string(jsonString)
}

// CountryCode - This class represents a country on earth.
type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (c Country) String() string {
	jsonString, err := json.Marshal(c)
	if err != nil {
		panic("CountryCode.String()")
	}
	return string(jsonString)
}

// Countries - A Result consisting of a list of countries.
type Countries struct {
	Countries    []Country `json:"country"`
	Total        int       `json:"total"`
	Page         int       `json:"page"`
	ItemsPerPage int       `json:"itemsPerPage"`
}

func (c Countries) String() string {
	jsonString, err := json.Marshal(c)
	if err != nil {
		panic("Countries.String()")
	}
	return string(jsonString)
}

// City -  	This class represents a city where Venues are located.
// Most of the original city data was taken from Geonames.org.
type City struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	StateCode string      `json:"stateCode"`
	State     string      `json:"state"`
	Coords    Coordinates `json:"coords"`
	Country   Country     `json:"country"`
}

func (c City) String() string {
	jsonString, err := json.Marshal(c)
	if err != nil {
		panic("City.String()")
	}
	return string(jsonString)
}

// Cities - A Result consisting of a list of cities.
type Cities struct {
	Cities       []City `json:"cities"`
	Total        int    `json:"total"`
	Page         int    `json:"page"`
	ItemsPerPage int    `json:"itemsPerPage"`
}

func (c Cities) String() string {
	jsonString, err := json.Marshal(c)
	if err != nil {
		panic("Cities.String()")
	}
	return string(jsonString)
}

// Error - Returned in case of an error.
type Error struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func (e Error) String() string {
	jsonString, err := json.Marshal(e)
	if err != nil {
		panic("Error.String()")
	}
	return string(jsonString)
}

// User - This class represents a user.
type User struct {
	UserID   string `json:"userId"`
	Fullname string `json:"fullname"`
	LastFM   string `json:"lastFm"`
	MySpace  string `json:"mySpace"`
	Twitter  string `json:"twitter"`
	Flickr   string `json:"flickr"`
	Website  string `json:"website"`
	About    string `json:"about"`
	URL      string `json:"url"`
}

func (u User) String() string {
	jsonString, err := json.Marshal(u)
	if err != nil {
		panic("User.String()")
	}
	return string(jsonString)
}

// Tour - The tour a setlist was a part of.
type Tour struct {
	Name string `json:"name"`
}

func (t Tour) String() string {
	jsonString, err := json.Marshal(t)
	if err != nil {
		panic("Tour.String()")
	}
	return string(jsonString)
}

// Song - This class represents a song that is part of a Set.
type Song struct {
	Name  string `json:"name"`
	With  Artist `json:"with,omitempty"`
	Cover Artist `json:"cover,omitempty"`
	Info  string `json:"info"`
	Tape  bool   `json:"tape"`
}

func (s Song) String() string {
	jsonString, err := json.Marshal(s)
	if err != nil {
		panic("Song.String()")
	}
	return string(jsonString)
}

// Set - A setlist consists of different (at least one) sets.
// Sets can either be sets as defined in the Guidelines or encores.
type Set struct {
	Name   string `json:"name"`
	Encore int    `json:"encore"`
	Song   []Song `json:"song"`
}

func (s Set) String() string {
	jsonString, err := json.Marshal(s)
	if err != nil {
		panic("Set.String()")
	}
	return string(jsonString)
}

type Sets struct {
	Set []Set `json:"set"`
}

func (s Sets) String() string {
	jsonString, err := json.Marshal(s)
	if err != nil {
		panic("Sets.String()")
	}
	return string(jsonString)
}

// Setlist - Setlists, that's what it's all about.
// So if you're trying to use this API without knowing what a setlist is
// then you're kinda wrong on this page ;-).
// A setlist can be distinguished from other setlists by its unique id.
// But as setlist.fm works the wiki way, there can be different versions of one setlist
// (each time a user updates a setlist a new version gets created).
// These different versions have a unique id on its own.
// So setlists can have the same id although they differ as far as the content is concerned -
// thus the best way to check if two setlists are the same is to compare their versionIds.
type Setlist struct {
	Artist      Artist `json:"artist"`
	Venue       Venue  `json:"venue"`
	Tour        Tour   `json:"tour"`
	Sets        Sets   `json:"sets"`
	Info        string `json:"info"`
	URL         string `json:"url"`
	ID          string `json:"id"`
	VersionID   string `json:"versionId"`
	EventDate   string `json:"eventDate"`
	LastUpdated string `json:"lastUpdated"`
}

func (s Setlist) String() string {
	jsonString, err := json.Marshal(s)
	if err != nil {
		panic("Setlist.String()")
	}
	return string(jsonString)
}

// IsEqual - Compares two setlists according to the setlist.fm API docs
func (s Setlist) IsEqual(s1 Setlist) bool {
	return s.ID == s1.ID && s.VersionID == s1.VersionID
}

// Setlists - A Result consisting of a list of setlists.
type Setlists struct {
	Setlists     []Setlist `json:"setlist"`
	Total        int       `json:"total"`
	Page         int       `json:"page"`
	ItemsPerPage int       `json:"itemsPerPage"`
}

func (s Setlists) String() string {
	jsonString, err := json.Marshal(s)
	if err != nil {
		panic("Setlists.String()")
	}
	return string(jsonString)
}

// Venue - Venues are places where concerts take place.
// They usually consist of a venue name and a city -
// but there are also some venues that do not have a city attached yet.
// In such a case, the city simply isn't set and the city and country
// may (but do not have to) be in the name.
type Venue struct {
	City City   `json:"city"`
	URL  string `json:"url"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (v Venue) String() string {
	jsonString, err := json.Marshal(v)
	if err != nil {
		panic("Venue.String()")
	}
	return string(jsonString)
}

// Venues - A Result consisting of a list of venues.
type Venues struct {
	Venues       []Venue `json:"venue"`
	Total        int     `json:"total"`
	Page         int     `json:"page"`
	ItemsPerPage int     `json:"itemsPerPage"`
}

func (v Venues) String() string {
	jsonString, err := json.Marshal(v)
	if err != nil {
		panic("Venues.String()")
	}
	return string(jsonString)
}

// Date - Represents a date on the calendar
type Date struct {
	Day, Month, Year int
}

// IsEmpty - Checks if the given Date object is equal to the default value object of Date
func (d Date) IsEmpty() bool {
	empty := Date{}
	return d == empty
}

func (d *Date) String() string {
	if d.IsEmpty() {
		return ""
	}
	if d.Day == 0 {
		d.Day = time.Now().Day()
	}
	if d.Month == 0 {
		d.Month = int(time.Now().Month())
	}
	if d.Year == 0 {
		d.Year = time.Now().Year()
	}
	return fmt.Sprintf("%d-%d-%d", d.Day, d.Month, d.Year)
}

// DateTime - Adds hours, minutes and seconds to the Date object
type DateTime struct {
	Date   Date
	Hour   int
	Minute int
	Second int
}

// IsEmpty - Checks if the given DateTime object is equal to the default value object of DateTime
func (d DateTime) IsEmpty() bool {
	empty := DateTime{}
	return d == empty
}

func (d *DateTime) String() string {
	if d.IsEmpty() {
		return ""
	}
	if d.Date.Year == 0 {
		d.Date.Year = time.Now().Year()
	}
	if d.Date.Month == 0 {
		d.Date.Month = int(time.Now().Month())
	}
	if d.Date.Day == 0 {
		d.Date.Day = time.Now().Day()
	}
	if d.Hour == 0 {
		d.Hour = time.Now().Hour()
	}
	if d.Minute == 0 {
		d.Minute = time.Now().Minute()
	}
	if d.Second == 0 {
		d.Second = time.Now().Second()
	}
	return fmt.Sprintf("%d%d%d%d%d%d",
		d.Date.Year,
		d.Date.Month,
		d.Date.Day,
		d.Hour,
		d.Minute,
		d.Second)
}

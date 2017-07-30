package api

import (
	"encoding/json"

	"time"
	"fmt"
)

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

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (c Country) String() string {
	jsonString, err := json.Marshal(c)
	if err != nil {
		panic("Country.String()")
	}
	return string(jsonString)
}

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

type Setlist struct {
	Artist      Artist `json:"artist"`
	Venue       Venue  `json:"venue"`
	Tour        Tour   `json:"tour"`
	Set         []Set  `json:"set"`
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

type Date struct {
	Day, Month, Year int
}

func (d Date) IsEmpty() bool{
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

type DateTime struct {
	Date   Date
	Hour   int
	Minute int
	Second int
}

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
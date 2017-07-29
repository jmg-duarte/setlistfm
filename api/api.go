package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	SETLIST_FM_API_KEY = ""
)

const (
	REST_ENDPOINT           = "https://api.setlist.fm/rest/"
	ARTIST_BY_MBID          = "1.0/artist/%s"
	ARTIST_SETLISTS_BY_MBID = "1.0/artist/%s/setlists"
	ARTISTS_SEARCH          = "1.0/search/artists"
	CITY_BY_GEOID           = "1.0/city/%s"
	CITY_SEARCH = "1.0/search/cities"
	COUNTRIES_LIST = "1.0/search/countries"
	HEADER_KEY              = "x-api-key"
	HEADER_ACCEPT_KEY       = "Accept"
	HEADER_ACCEPT_VALUE     = "application/json"
)

func ArtistByMBID(MBID string) (*Artist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+ARTIST_BY_MBID, MBID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	artist := new(Artist)
	err = json.Unmarshal(body, &artist)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func ArtistSetlistsByMBID(MBID string, page int) (*Setlists, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+ARTIST_SETLISTS_BY_MBID, MBID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	q := req.URL.Query()
	// Consider changing page to string
	if page > 0 {
		q.Add("p", strconv.Itoa(page))
	}
	req.URL.RawQuery = q.Encode()

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	setlists := new(Setlists)
	err = json.Unmarshal(body, &setlists)
	if err != nil {
		return nil, err
	}

	return setlists, nil
}

func CityByGeoID(geoID string) (*City, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+CITY_BY_GEOID, geoID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	city := new(City)
	err = json.Unmarshal(body, &city)
	if err != nil {
		return nil, err
	}

	return city, nil
}

func SearchForArtists(artistMbid, artistName string, artistTmid, page int) (*Artists, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+ARTISTS_SEARCH, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	q := req.URL.Query()

	if artistMbid != "" {
		q.Add("artistMbid", artistMbid)
	}

	if artistName != "" {
		q.Add("artistName", artistName)
	}

	if artistTmid > 0 {
		q.Add("artistTmid", strconv.Itoa(artistTmid))
	}

	// Consider changing page to string
	if page > 0 {
		q.Add("p", strconv.Itoa(page))
	}
	req.URL.RawQuery = q.Encode()

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	artists := new(Artists)
	err = json.Unmarshal(body, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func SearchForCities(country, city, state, stateCode string, page int) (*Cities, error){
	req, err := http.NewRequest("GET", REST_ENDPOINT + CITY_SEARCH, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	q := req.URL.Query()

	if country != ""{
		q.Add("country", country)
	}

	if city == ""{
		return nil, fmt.Errorf("setlistfmapi: you must specify a city when searching for a city")
	}
	q.Add("name", city)

	if state != ""{
		q.Add("state", state)
	}

	if stateCode != ""{
		q.Add("stateCode", stateCode)
	}

	// Consider changing page to string
	if page > 0 {
		q.Add("p", strconv.Itoa(page))
	}
	req.URL.RawQuery = q.Encode()

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	cities := new(Cities)
	err = json.Unmarshal(body, &cities)
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func ListAllCountries() (*Countries, error){
	req, err := http.NewRequest("GET", REST_ENDPOINT + COUNTRIES_LIST, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	countries := new(Countries)
	err = json.Unmarshal(body, &countries)
	if err != nil {
		return nil, err
	}

	return countries, nil
}

func executeRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

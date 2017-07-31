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
	SETLIST_FM_API_KEY = "e1d7b2e5-c3b7-459b-9597-cf5fd9296182"
)

const (
	REST_ENDPOINT = "https://api.setlist.fm/rest/"

	ARTIST_BY_MBID          = "1.0/artist/%s"
	ARTIST_SETLISTS_BY_MBID = "1.0/artist/%s/setlists"
	ARTISTS_SEARCH          = "1.0/search/artists"
	CITY_BY_GEOID           = "1.0/city/%s"
	CITY_SEARCH             = "1.0/search/cities"
	COUNTRIES_LIST          = "1.0/search/countries"
	HEADER_KEY              = "x-api-key"
	HEADER_ACCEPT_KEY       = "Accept"
	HEADER_ACCEPT_VALUE     = "application/json"
	SETLIST_SEARCH          = "1.0/search/setlists"
	SETLIST_BY_VERSIONID    = "1.0/setlist/version/%s"
	SETLIST_BY_ID           = "1.0/setlist/%s"
	USER_BY_ID              = "1.0/user/%s"
	USER_ATTENDED_CONCERTS  = "1.0/user/%s/attended"
	USER_EDITED_PLAYLISTS   = "1.0/user/%s/edited"
	VENUE_SEARCH            = "1.0/search/venues"
	VENUE_BY_ID             = "1.0/venue/%s"
	VENUE_SETLISTS_BY_ID    = "/1.0/venue/%s/setlists"
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

func SearchForArtists(a ArtistsQuery) (*Artists, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+ARTISTS_SEARCH, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	req = a.BuildQuery(*req)

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

func SearchForCities(c CityQuery) (*Cities, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+CITY_SEARCH, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	req = c.BuildQuery(*req)

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

func ListAllCountries() (*Countries, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+COUNTRIES_LIST, nil)
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

func SearchForSetlists(s SetlistQuery) (*Setlists, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+SETLIST_SEARCH, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	req = s.BuildQuery(*req)
	fmt.Println(req.URL.RawQuery)
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

func SearchForVenues(v VenueQuery) (*Venues, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+VENUE_SEARCH, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	req = v.BuildQuery(*req)
	fmt.Println(req.URL.RawQuery)
	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	venues := new(Venues)
	err = json.Unmarshal(body, &venues)
	if err != nil {
		return nil, err
	}

	return venues, nil
}

func SetlistByVersionID(versionID string) (*Setlist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+SETLIST_BY_VERSIONID, versionID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	setlist := new(Setlist)
	err = json.Unmarshal(body, &setlist)
	if err != nil {
		return nil, err
	}

	return setlist, nil
}

func SetlistByID(ID string) (*Setlist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+SETLIST_BY_ID, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	setlist := new(Setlist)
	err = json.Unmarshal(body, &setlist)
	if err != nil {
		return nil, err
	}

	return setlist, nil
}

func UserByID(ID string) (*User, error){
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+USER_BY_ID, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(req)
	if err != nil {
		return nil, err
	}

	user := new(User)
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UserAttendedConcerts(ID string, page int) (*Setlists, error){
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+USER_ATTENDED_CONCERTS, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	q:= req.URL.Query()
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

func UserEditedPlaylists(ID string, page int) (*Setlists, error){
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+USER_EDITED_PLAYLISTS, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, SETLIST_FM_API_KEY)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	q:= req.URL.Query()
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

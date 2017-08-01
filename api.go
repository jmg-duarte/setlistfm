package setlistfm

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	REST_ENDPOINT = "https://api.setlist.fm/rest/"

	ARTIST_BY_MBID          = "1.0/artist/%s"
	ARTIST_SETLISTS_BY_MBID = "1.0/artist/%s/setlists"
	ARTISTS_SEARCH          = "1.0/search/artists"
	CITY_BY_GEOID           = "1.0/city/%s"
	CITY_SEARCH             = "1.0/search/cities"
	COUNTRIES_LIST          = "1.0/search/countries"
	SETLIST_SEARCH          = "1.0/search/setlists"
	SETLIST_BY_VERSIONID    = "1.0/setlist/version/%s"
	SETLIST_BY_ID           = "1.0/setlist/%s"
	USER_BY_ID              = "1.0/user/%s"
	USER_ATTENDED_CONCERTS  = "1.0/user/%s/attended"
	USER_EDITED_SETLISTS    = "1.0/user/%s/edited"
	VENUE_SEARCH            = "1.0/search/venues"
	VENUE_BY_ID             = "1.0/venue/%s"
	VENUE_SETLISTS_BY_ID    = "1.0/venue/%s/setlists"

	HEADER_KEY          = "x-api-key"
	HEADER_ACCEPT_KEY   = "Accept"
	HEADER_ACCEPT_VALUE = "application/json"
)

func (cl Client) ArtistByMBID(ctx context.Context, MBID string) (*Artist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+ARTIST_BY_MBID, MBID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(ctx, req)
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

func (cl Client) ArtistSetlistsByMBID(ctx context.Context, MBID string, page int) (*Setlists, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+ARTIST_SETLISTS_BY_MBID, MBID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	q := req.URL.Query()
	// Consider changing page to string
	if page > 0 {
		q.Add("p", strconv.Itoa(page))
	}
	req.URL.RawQuery = q.Encode()

	body, err := executeRequest(ctx, req)
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

func (cl Client) CityByGeoID(ctx context.Context, geoID string) (*City, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+CITY_BY_GEOID, geoID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(ctx, req)
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

func (cl Client) SearchForArtists(ctx context.Context, a ArtistsQuery) (*Artists, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+ARTISTS_SEARCH, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	req = a.BuildQuery(*req)

	body, err := executeRequest(ctx, req)
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

func (cl Client) SearchForCities(ctx context.Context, c CityQuery) (*Cities, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+CITY_SEARCH, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	req = c.BuildQuery(*req)

	body, err := executeRequest(ctx, req)
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

func (cl Client) ListAllCountries(ctx context.Context) (*Countries, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+COUNTRIES_LIST, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(ctx, req)
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

func (cl Client) SearchForSetlists(ctx context.Context, s SetlistQuery) (*Setlists, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+SETLIST_SEARCH, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	req = s.BuildQuery(*req)
	fmt.Println(req.URL.RawQuery)
	body, err := executeRequest(ctx, req)
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

func (cl Client) SearchForVenues(ctx context.Context, v VenueQuery) (*Venues, error) {
	req, err := http.NewRequest("GET", REST_ENDPOINT+VENUE_SEARCH, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	req = v.BuildQuery(*req)
	fmt.Println(req.URL.RawQuery)
	body, err := executeRequest(ctx, req)
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

func (cl Client) SetlistByVersionID(ctx context.Context, versionID string) (*Setlist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+SETLIST_BY_VERSIONID, versionID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(ctx, req)
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

func (cl Client) SetlistByID(ctx context.Context, ID string) (*Setlist, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+SETLIST_BY_ID, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(ctx, req)
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

func (cl Client) UserByID(ctx context.Context, ID string) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+USER_BY_ID, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(ctx, req)
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

func (cl Client) UserAttendedConcerts(ctx context.Context, ID string, page int) (*Setlists, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+USER_ATTENDED_CONCERTS, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	q := req.URL.Query()
	if page > 0 {
		q.Add("p", strconv.Itoa(page))
	}
	req.URL.RawQuery = q.Encode()

	body, err := executeRequest(ctx, req)
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

func (cl Client) UserEditedSetlists(ctx context.Context, ID string, page int) (*Setlists, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+USER_EDITED_SETLISTS, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	q := req.URL.Query()
	if page > 0 {
		q.Add("p", strconv.Itoa(page))
	}
	req.URL.RawQuery = q.Encode()

	body, err := executeRequest(ctx, req)
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

func (cl Client) VenueByID(ctx context.Context, ID string) (*Venue, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+VENUE_BY_ID, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	body, err := executeRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	venue := new(Venue)
	err = json.Unmarshal(body, &venue)
	if err != nil {
		return nil, err
	}

	return venue, nil
}

func (cl Client) VenueSetlists(ctx context.Context, ID string, page int) (*Setlists, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(REST_ENDPOINT+VENUE_SETLISTS_BY_ID, ID), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(HEADER_KEY, cl.APIKey)
	req.Header.Add(HEADER_ACCEPT_KEY, HEADER_ACCEPT_VALUE)

	q := req.URL.Query()
	if page > 0 {
		q.Add("p", strconv.Itoa(page))
	}
	req.URL.RawQuery = q.Encode()

	body, err := executeRequest(ctx, req)
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

func executeRequest(ctx context.Context, req *http.Request) ([]byte, error) {
	req = req.WithContext(ctx)
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

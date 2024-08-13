package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) ListLocation(areaIDorName *string) (RespShallowLocation, error) {
	url := baseURL + "/location-area/"
	if areaIDorName != nil {
		url += *areaIDorName
	} else {
		return RespShallowLocation{}, errors.New("need to provide a location area id or name")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocation{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return RespShallowLocation{}, errors.New("location not found")
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocation{}, err
	}

	locationsResp := RespShallowLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocation{}, err
	}

	return locationsResp, nil
}

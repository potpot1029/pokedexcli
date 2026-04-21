package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// creating GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	// make request, get and process response
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	// decode response
	var data LocationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return LocationArea{}, err
	}

	return data, nil
}

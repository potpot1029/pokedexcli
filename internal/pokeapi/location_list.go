package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cached, ok := c.pokeCache.Get(url)
	if ok {
		// unmarshal cached
		cachedLocation := LocationArea{}
		err := json.Unmarshal(cached, &cachedLocation)
		if err != nil {
			return LocationArea{}, err
		}
		return cachedLocation, nil
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

	// unmarshal response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.pokeCache.Add(url, data)

	return locationArea, nil
}

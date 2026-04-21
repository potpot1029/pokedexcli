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

func (c *Client) GetLocationPokemon(name string) (LocationPokemons, error) {
	url := baseURL + "/location-area/" + name

	cached, ok := c.pokeCache.Get(url)
	if ok {
		// unmarshal cached
		cachedPokemons := LocationPokemons{}
		err := json.Unmarshal(cached, &cachedPokemons)
		if err != nil {
			return LocationPokemons{}, err
		}
		return cachedPokemons, nil
	}

	// creating GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationPokemons{}, err
	}

	// make request, get and process response
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationPokemons{}, err
	}
	defer res.Body.Close()

	// unmarshal response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationPokemons{}, err
	}

	locationPokemons := LocationPokemons{}
	err = json.Unmarshal(data, &locationPokemons)
	if err != nil {
		return LocationPokemons{}, err
	}

	c.pokeCache.Add(url, data)

	return locationPokemons, nil
}

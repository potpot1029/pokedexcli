package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

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

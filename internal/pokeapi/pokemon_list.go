package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	cached, ok := c.pokeCache.Get(url)
	if ok {
		// unmarshal cached
		cachedLocation := Pokemon{}
		err := json.Unmarshal(cached, &cachedLocation)
		if err != nil {
			return Pokemon{}, err
		}
		return cachedLocation, nil
	}

	// creating GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	// make request, get and process response
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	// unmarshal response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.pokeCache.Add(url, data)

	return pokemon, nil
}

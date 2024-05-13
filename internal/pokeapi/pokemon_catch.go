package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) CatchPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	if val, ok := c.cache.Get(url); ok {
		locationsResp := Pokemon{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Pokemon{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	locationsResp := Pokemon{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
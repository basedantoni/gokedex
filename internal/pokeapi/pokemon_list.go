package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListPokemon(area string) (FullLocationArea, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		locationsResp := FullLocationArea{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return FullLocationArea{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return FullLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return FullLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return FullLocationArea{}, err
	}

	locationsResp := FullLocationArea{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return FullLocationArea{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
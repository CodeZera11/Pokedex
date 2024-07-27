package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		respPokemon := RespPokemon{}

		err := json.Unmarshal(val, &respPokemon)

		if err != nil {
			return RespPokemon{} ,err
		}

		return respPokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return RespPokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RespPokemon{}, err
	}

	respPokemon := RespPokemon{}

	err = json.Unmarshal(data, &respPokemon)

	if err != nil {
		return RespPokemon{}, err
	}

	return respPokemon, nil
}
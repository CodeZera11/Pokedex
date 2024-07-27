package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (RespLocation, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespLocation{}

		err := json.Unmarshal(val, &locationResp)

		if err != nil {
			return RespLocation{}, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return RespLocation{}, err
	}


	resp, err := c.httpClient.Do(req)

	if err != nil {
		return RespLocation{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RespLocation{}, err
	}

	locationResp := RespLocation{}
	err = json.Unmarshal(data, &locationResp)

	if err != nil {
		return RespLocation{}, err
	}

	return locationResp, nil
}
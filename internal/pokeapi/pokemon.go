package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(idOrName *string) (RespPokemon, error) {
	url := baseURL + "/pokemon/"
	if idOrName != nil {
		url += *idOrName
	} else {
		return RespPokemon{}, errors.New("need to provide a pokemon area id or name")
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

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	return pokemonResp, nil
}

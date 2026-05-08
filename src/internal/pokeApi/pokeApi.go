package pokeApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	pokeCache "go-pokedex/src/internal/pokeCache"
)

const baseUrl = "https://pokeapi.co/api/v2/"

func GetLocationAreas(url string, cache *pokeCache.Cache) (LocationAreaResponse, error) {
	if url == "" {
		url = baseUrl + "/location-area/"
	}

	cachedResponse, cacheHit := cache.Get(url)
	if cacheHit {
		fmt.Println("Cache Hit!")
		results := LocationAreaResponse{}
		err := json.Unmarshal(cachedResponse, &results)
		if err != nil {
			return LocationAreaResponse{}, errors.New("error unmarshaling cache")
		}
		return results, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, errors.New("HTTP error")
	}
	if resp.StatusCode != http.StatusOK {
		return LocationAreaResponse{}, errors.New("error reaching api")
	}
	decoder := json.NewDecoder(resp.Body)
	results := LocationAreaResponse{}
	err = decoder.Decode(&results)
	if err != nil {
		return LocationAreaResponse{}, errors.New("error decoding api response")
	}

	cachedbody, err := json.Marshal(results)
	if err != nil {
		return results, errors.New("cannot cache results")
	}
	
	cache.Add(url, cachedbody)

	return results, nil
}

func GetLocationArea(name string, cache *pokeCache.Cache) (ExploreResponse, error) {
	url := baseUrl + "location-area/" + name

	if cached, ok := cache.Get(url); ok {
		result := ExploreResponse{}
		if err := json.Unmarshal(cached, &result); err != nil {
			return ExploreResponse{}, errors.New("error unmarshaling cache")
		}
		return result, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return ExploreResponse{}, errors.New("HTTP error")
	}
	if resp.StatusCode != http.StatusOK {
		return ExploreResponse{}, fmt.Errorf("location area %q not found", name)
	}

	result := ExploreResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return ExploreResponse{}, errors.New("error decoding api response")
	}

	if body, err := json.Marshal(result); err == nil {
		cache.Add(url, body)
	}

	return result, nil
}

func GetPokemon(name string, cache *pokeCache.Cache) (Pokemon, error) {
	url := baseUrl + "pokemon/" + name

	if cached, ok := cache.Get(url); ok {
		result := Pokemon{}
		if err := json.Unmarshal(cached, &result); err != nil {
			return Pokemon{}, errors.New("error unmarshaling cache")
		}
		return result, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return Pokemon{}, errors.New("HTTP error")
	}
	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("pokemon %q not found", name)
	}

	result := Pokemon{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Pokemon{}, errors.New("error decoding api response")
	}

	if body, err := json.Marshal(result); err == nil {
		cache.Add(url, body)
	}

	return result, nil
}
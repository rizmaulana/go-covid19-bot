package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type client struct {
	httpClient     http.Client
	covidKalselURL string
	hereMapsURL    string
	hereMapsKey    string
}

type Client interface {
	FetchCovidData() ([]CovidResponse, error)
	ReverseCoordinate(lat float32, lon float32) (*HereMapsResponse, error)
}

func NewClient(covidKalselURL string, hereMapsURL string, hereMapsKey string) Client {
	httpClient := http.Client{
		Timeout: 10,
	}

	return &client{httpClient, covidKalselURL, hereMapsURL, hereMapsKey}
}

func (c *client) FetchCovidData() ([]CovidResponse, error) {
	resp, err := c.httpClient.Get(c.covidKalselURL)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var covidResponse []CovidResponse
	err = json.Unmarshal(b, &covidResponse)
	if err != nil {
		return nil, err
	}

	return covidResponse, nil
}

func (c *client) ReverseCoordinate(lat float32, lon float32) (*HereMapsResponse, error) {
	u, err := url.Parse(c.hereMapsURL)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add("prox", fmt.Sprintf("%f,%f", lat, lon))
	q.Add("apiKey", c.hereMapsKey)

	resp, err := c.httpClient.Get(q.Encode())
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var hereMapResponse HereMapsResponse
	err = json.Unmarshal(b, &hereMapResponse)
	if err != nil {
		return nil, err
	}

	return &hereMapResponse, nil
}

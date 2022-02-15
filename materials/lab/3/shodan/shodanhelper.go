package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client) HostFacets() (*[]string, error) {
	res, err := http.Get(
		fmt.Sprintf("%s/shodan/host/search/facets?key=%s", BaseURL, s.apiKey),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret []string
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (s *Client) HostFilters() (*[]string, error) {
	res, err := http.Get(
		fmt.Sprintf("%s/shodan/host/search/filters?key=%s", BaseURL, s.apiKey),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret []string
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
// Copyright 2026 mnisyif
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package goblinapi is meant to talk with the osrs wiki api and retrieve data on osrs items
package goblinapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchAndCache[T any](gobClient *Client, url string, ttl time.Duration) (T, error) {
	var result T

	if gobClient.cache != nil {
		cached, exists := gobClient.cache.Get(url)
		if exists {
			err := json.Unmarshal(cached, &result)
			return result, err
		}
	}

	req, err := http.NewRequestWithContext(gobClient.ctx, "GET", url, nil)
	if err != nil {
		return result, err
	}

	req.Header.Add("User-Agent", gobClient.userAgent)

	res, err := gobClient.httpClient.Do(req)
	if err != nil {
		return result, err
	}

	if res.StatusCode != http.StatusOK {
		return result, fmt.Errorf("API returned %d", res.StatusCode)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	if gobClient.cache != nil {
		gobClient.cache.Add(url, data, ttl)
	}

	err = json.Unmarshal(data, &result)

	return result, err
}

func (c *Client) FetchMappings() ([]Item, error) {
	url := fmt.Sprintf("%s/mapping", baseURL)

	return fetchAndCache[[]Item](c, url, 30*time.Minute)
}

func (c *Client) FetchLatest() (LatestPrices, error) {
	url := fmt.Sprintf("%s/latest", baseURL)

	return fetchAndCache[LatestPrices](c, url, 60*time.Second)
}

func (c *Client) Fetch5Min() (AveragePrices, error) {
	url := fmt.Sprintf("%s/5m", baseURL)

	return fetchAndCache[AveragePrices](c, url, 5*time.Minute)
}

func (c *Client) Fetch1Hour() (AveragePrices, error) {
	url := fmt.Sprintf("%s/1h", baseURL)

	return fetchAndCache[AveragePrices](c, url, 1*time.Hour)
}

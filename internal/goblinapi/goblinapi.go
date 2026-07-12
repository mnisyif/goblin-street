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
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func fetch[T any](gobClient *Client, url string) (T, error) {
	var result T

	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return result, err
	}

	req.Header.Add("User-Agent", gobClient.userAgent)

	res, err := gobClient.httpClient.Do(req)
	if err != nil {
		return result, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)

	return result, err
}

func (c *Client) FetchMappings() ([]Item, error) {
	url := mappings

	return fetch[[]Item](c, url)
}

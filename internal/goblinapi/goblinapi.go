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

type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Examine  string `json:"examine"`
	Members  bool   `json:"members"`
	Icon     string `json:"icon"`
	Value    int    `json:"value"`
	Limit    int    `json:"limit"`
	LowAlch  int    `json:"lowalch"`
	HighAlch int    `json:"highalch"`
}

func (c *Client) FetchMappings() ([]Item, error) {
	var result []Item

	req, err := http.NewRequestWithContext(context.Background(), "GET", mappings, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.userAgent)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

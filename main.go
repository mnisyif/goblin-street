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

package main

import (
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnisyif/goblin-street/internal/goblinapi"
	"github.com/mnisyif/goblin-street/internal/goblinengine"
	"github.com/mnisyif/goblin-street/internal/goblintui"
)

func main() {
	userAgent := "goblin-street/v0.1 (github.com/mnisyif/goblin-street; mnisyif@gmail.com)"
	client := goblinapi.New(userAgent)

	items, err := client.FetchMappings()
	if err != nil {
		os.Exit(1)
	}

	prices, err := client.Fetch5Min()
	if err != nil {
		os.Exit(1)
	}

	var rows []goblintui.ModelMarket
	for _, item := range items {
		id := strconv.Itoa(item.ID)
		entry, ok := prices.Data[id]
		if !ok {
			continue // skip untraded items
		}
		if entry.AvgBuy == 0 || entry.AvgSell == 0 {
			continue // skip incomplete data
		}

		spread := entry.AvgBuy - entry.AvgSell
		roi := goblinengine.ROI(entry.AvgSell, entry.AvgBuy)
		volume := entry.BuyVolume + entry.SellVolume

		rows = append(rows, goblintui.ModelMarket{
			Name:   item.Name,
			Buy:    entry.AvgBuy,
			Sell:   entry.AvgSell,
			Spread: spread,
			ROI:    roi,
			Volume: volume,
		})
	}

	if len(rows) > 100 {
		rows = rows[:10]
	}

	history := []string{} // empty for now

	p := tea.NewProgram(goblintui.New(rows, history))
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}

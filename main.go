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
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnisyif/goblin-street/internal/goblinapi"
	"github.com/mnisyif/goblin-street/internal/goblincache"
	"github.com/mnisyif/goblin-street/internal/goblinengine"
	"github.com/mnisyif/goblin-street/internal/goblintui"
)

func main() {
	userAgent := "goblin-street/v0.1 (github.com/mnisyif/goblin-street; mnisyif@gmail.com)"
	newCache, _ := goblincache.New(30 * time.Second)
	client := goblinapi.New(userAgent, newCache)

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

	history := []goblintui.ModelHistory{
		{Item: "Cannonball", Qty: 10000, BuyPrice: 192, SellPrice: 204, Profit: 98000, Date: "2025-07-14"},
		{Item: "Nature rune", Qty: 5000, BuyPrice: 241, SellPrice: 255, Profit: 59300, Date: "2025-07-14"},
		{Item: "Dragon bones", Qty: 500, BuyPrice: 1820, SellPrice: 1960, Profit: 59300, Date: "2025-07-13"},
		{Item: "Yew logs", Qty: 2000, BuyPrice: 305, SellPrice: 318, Profit: 19740, Date: "2025-07-13"},
		{Item: "Rune essence", Qty: 25000, BuyPrice: 4, SellPrice: 5, Profit: 14750, Date: "2025-07-12"},
		{Item: "Death rune", Qty: 8000, BuyPrice: 213, SellPrice: 224, Profit: 71200, Date: "2025-07-12"},
		{Item: "Magic shortbow", Qty: 200, BuyPrice: 910, SellPrice: 985, Profit: 12830, Date: "2025-07-11"},
		{Item: "Adamant arrow", Qty: 50000, BuyPrice: 32, SellPrice: 34, Profit: 83000, Date: "2025-07-11"},
		{Item: "Green dragonhide", Qty: 1000, BuyPrice: 1620, SellPrice: 1710, Profit: 71100, Date: "2025-07-10"},
		{Item: "Lobster", Qty: 3000, BuyPrice: 182, SellPrice: 194, Profit: 26940, Date: "2025-07-10"},
	}

	p := tea.NewProgram(goblintui.New(rows, history))
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}

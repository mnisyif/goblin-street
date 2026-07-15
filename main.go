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
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnisyif/goblin-street/internal/goblinapi"
	"github.com/mnisyif/goblin-street/internal/goblintui"
	"github.com/mnisyif/goblin-street/internal/goblinengine"
)

func main() {
	rows := []goblintui.ModelRow{}
	history := []string{}

	userAgent := "goblin-street/v0.1 (github.com/mnisyif/goblin-street; mnisyif@gmail.com)"

	goblinClient := goblinapi.New(userAgent)
	goblinTui := goblintui.New(rows, history)

	items, err := goblinClient.FetchMappings()
	if err != nil {
		fmt.Printf("Couldnt fetch mappings from osrs wiki: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Item name: %s\n", items[0].Name)
	p := tea.NewProgram(goblinTui)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)

	// avgPrices, err := goblinClient.FetchLatest()
	avgPrices, err := goblinClient.Fetch1Hour()
	if err != nil {
		fmt.Printf("Could not fetch prices of last hour: %s", err)
		os.Exit(1)
	}

	id_int := 10
	id_str := strconv.Itoa(items[id_int].ID)
	entry, ok := avgPrices.Data[id_str]
	if !ok {
		fmt.Printf("Item %s not traded recently\n", items[id_int].Name)
		return
	}

	fmt.Printf("Item: %v\n", entry)
	profit := goblinengine.ProfitGP(entry.AvgBuy, entry.AvgSell)
	roi := goblinengine.ROI(entry.AvgBuy, entry.AvgSell)
	margin := goblinengine.MarginPct(entry.AvgBuy, entry.AvgSell)

	fmt.Printf("Item: %s\n", items[id_int].Name)
	fmt.Printf("Profit per item: %d gp\n", profit)
	fmt.Printf("ROI: %.2f%%\n", roi)
	fmt.Printf("Margin: %.2f%%\n", margin)
}

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

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mnisyif/goblin-street/internal/goblinapi"
	"github.com/mnisyif/goblin-street/internal/goblintui"
)

func main() {
	rows := []goblintui.ModelRow{
		{Name: "Cannonball", Buy: 200, Sell: 210, Spread: 10, ROI: 5.0, Volume: 10000},
	}
	history := []string{
		"Bought 10k Cannonball @ 200",
		"Sold 10k Cannonball @ 210 — profit 90k",
	}

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
		os.Exit(1)
	}
}

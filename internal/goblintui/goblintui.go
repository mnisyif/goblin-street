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

// Package goblintui is implementd to represennt goblin-street in a terminal environment
package goblintui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func New(rows []ModelMarket, history []ModelHistory) *Model {
	return &Model{
		ActiveTab:    0,
		Cursor:       0,
		WindowHeight: 10,
		Rows:         rows,
		History:      history,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.WindowHeight = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab":
			if m.ActiveTab == 0 {
				m.MarketCursor = m.Cursor
				m.MarketOffset = m.ScrollOffset
			} else {
				m.HistoryCursor = m.Cursor
				m.HistoryOffset = m.ScrollOffset
			}

			m.ActiveTab = 1 - m.ActiveTab

			if m.ActiveTab == 0 {
				m.Cursor = m.MarketCursor
				m.ScrollOffset = m.MarketOffset
			} else {
				m.Cursor = m.HistoryCursor
				m.ScrollOffset = m.HistoryOffset
			}
		case "up", "k":
			// visibleRows := m.WindowHeight - 7
			if m.Cursor > 0 {
				m.Cursor--
			}
			if m.Cursor < m.ScrollOffset {
				m.ScrollOffset--
			}

		case "down", "j":
			visibleRows := m.WindowHeight - 7
			max := len(m.Rows) - 1
			if m.ActiveTab == 1 {
				max = len(m.History) - 1
			}
			if m.Cursor < max {
				m.Cursor++
			}
			if m.Cursor-m.ScrollOffset >= visibleRows {
				m.ScrollOffset++
			}
		}
	}
	return m, nil
}

func (m *Model) View() string {
	s := "Welcome to Goblin Street, your advisor to immense wealth in gelinor\n\n"
	tabBar := " "
	tableWidth := 70 // 2(cursor) + 22(name) + 8(buy) + 8(sell) + 8(spread) + 8(roi%) + 8(volume)
	padding := 0

	switch m.ActiveTab {
	// if tab = 0 -> market
	case 0:
		tabBar = "> Market < History  "
	// if tab = 1 -> history
	case 1:
		tabBar = "  Market > History <"
	}

	padding = (tableWidth - len(tabBar)) / 2
	s += strings.Repeat(" ", padding) + tabBar + "\n"
	s += strings.Repeat("-", tableWidth) + "\n"

	if m.ActiveTab == 0 {
		s += fmt.Sprintf("%2s %-22s %8s %8s %8s %8s %8s\n", "", "Name", "Buy", "Sell", "Spread", "ROI%", "Volume")
		visibleRows := m.WindowHeight - 7
		if visibleRows < 1 {
			visibleRows = 1
		}
		start := m.ScrollOffset
		end := start + visibleRows
		if end > len(m.Rows) {
			end = len(m.Rows)
		}

		for i := start; i < end; i++ {
			row := m.Rows[i]
			// for i, row := range m.Rows {
			cursor := "  "
			if m.Cursor == i {
				cursor = "> "
			}
			s += fmt.Sprintf("%2s %-22s %8d %8d %8d %7.1f%% %8d\n",
				cursor, row.Name, row.Buy, row.Sell, row.Spread, row.ROI, row.Volume)
		}
	} else {
		s += fmt.Sprintf("%2s %-16s %8s %8s %8s %10s %12s\n", "", "Item", "Qty", "Buy", "Sell", "Profit", "Date")
		for i, entry := range m.History {
			cursor := "  "
			if m.Cursor == i {
				cursor = "> "
			}
			s += fmt.Sprintf("%2s %-16s %8d %8d %8d %10d %12s\n", cursor, entry.Item, entry.Qty, entry.BuyPrice, entry.SellPrice, entry.Profit, entry.Date)
		}
	}

	s += " \nPress q to quit. Tab to switch views.\n"
	return s
}

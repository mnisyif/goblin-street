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

import tea "github.com/charmbracelet/bubbletea"

type ModelRow struct {
	Name   string
	Buy    int
	Sell   int
	Spread int
	ROI    float64
	Volume int
}

type Model struct {
	ActiveTab int
	Cursor    int
	Rows      []ModelRow
	History   []string
}

func New(rows []ModelRow, history []string) *Model {
	return &Model{
		ActiveTab: 0,
		Cursor:    0,
		Rows:      rows,
		History:   history,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

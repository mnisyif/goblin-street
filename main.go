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

	"github.com/mnisyif/goblin-street/internal/goblinapi"
)

func main() {
	userAgent := "goblin-street/v0.1 (github.com/mnisyif/goblin-street; mnisyif@gmail.com)"
	goblinClient := goblinapi.New(userAgent)

	items, err := goblinClient.FetchMappings()
	if err != nil {
		fmt.Printf("Couldnt fetch mappings from osrs wiki: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Item name: %s\n", items[0].Name)
}

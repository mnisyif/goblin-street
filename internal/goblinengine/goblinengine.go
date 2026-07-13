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

// Package goblinengine is a package designed to compute scores for items that are best for flipping
package goblinengine

func GETax(sellPrice int) int {
	return sellPrice / 100 // GE tax is 1% on selling price
}

func ProfitGP(buyPrice, sellPrice int) int {
	return sellPrice - GETax(sellPrice) - buyPrice
}

func ROI(buyPrice, sellPrice int) float64 {
	return float64(ProfitGP(buyPrice, sellPrice)) / float64(buyPrice) * 100
}

func MarginPct(buyPrice, sellPrice int) float64 {
	return float64(ProfitGP(buyPrice, sellPrice)) / float64(sellPrice) * 100
}

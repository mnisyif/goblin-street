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

// Package goblincache is for caching requests from OSRS wiki prices api
package goblincache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
	ttl       time.Duration
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.Mutex
}

func New(refreshTimer time.Duration) (*Cache, error) {
	newCache := &Cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}

	go newCache.reapLoop(refreshTimer)
	return newCache, nil
}

func (c *Cache) Add(key string, val []byte, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
		ttl:       ttl,
	}
	c.cacheMap[key] = newEntry

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	result, exists := c.cacheMap[key]

	return result.val, exists
}

func (c *Cache) reapLoop(refreshTimer time.Duration) {
	for {
		time.Sleep(refreshTimer)
		c.mu.Lock()
		for key, entry := range c.cacheMap {
			if time.Since(entry.createdAt) > entry.ttl {
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}

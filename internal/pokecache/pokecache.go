package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	Val       []byte
	CreatedAt time.Time
}

type Pokecache struct {
	CacheMap map[string]cacheEntry
	sync.RWMutex
}

func (c *Pokecache) reapLoop(interval time.Duration) {
	intervalSeconds := interval.Seconds()
	c.Lock()
	for key, val := range c.CacheMap {
		diffInSeconds := time.Since(val.CreatedAt).Seconds()
		if diffInSeconds >= intervalSeconds {
			delete(c.CacheMap, key)
		}
	}
	c.Unlock()
}

func NewPokecache(interval time.Duration) *Pokecache {
	c := Pokecache{}
	c.CacheMap = make(map[string]cacheEntry)

	go func() {
		ticker := time.NewTicker(interval)

		for range ticker.C {
			c.reapLoop(interval)
		}
		/*
			these two are equivalent
			for {
				select {
				case <-ticker.C:
				}
			}
		*/
	}()

	return &c
}

func (c *Pokecache) Add(entry string, val []byte) {
	newCacheEntry := cacheEntry{}
	c.Lock()
	newCacheEntry.Val = val
	newCacheEntry.CreatedAt = time.Now()
	c.CacheMap[entry] = newCacheEntry
	c.Unlock()
}

func (c *Pokecache) Get(entry string) ([]byte, bool) {
	c.RLock()
	defer c.RUnlock()
	v, ok := c.CacheMap[entry]
	return v.Val, ok
}

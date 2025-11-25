package pokecache_test

import (
	"testing"
	"time"

	"github.com/kahnaisehC/pokedex/internal/pokecache"
)

func TestCache(t *testing.T) {
	baseInterval := time.Millisecond * 5
	longInterval := baseInterval + baseInterval
	cache := pokecache.NewPokecache(baseInterval)

	type testCase struct {
		key      string
		expected []byte
	}

	testCases := []testCase{
		{
			key:      "banana",
			expected: []byte("Im an apple akshually"),
		},
		{},
	}

	for _, c := range testCases {
		cache.Add(c.key, c.expected)
		actual, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected a value in cache, none found. \nkey: %s\nexpected: %s\n", c.key, c.expected)
		}
		if string(actual) != string(c.expected) {
			t.Errorf("wrong value stored in cache. \nkey: %s\nexpected: %s\nactual: %s", c.key, c.expected, actual)
		}
		time.Sleep(longInterval)

		actual, ok = cache.Get(c.key)
		if ok {
			t.Errorf("expected NO value in cache, found one. \nkey: %s\nexpected: %s\nactual: %s", c.key, c.expected, actual)
		}

	}
}

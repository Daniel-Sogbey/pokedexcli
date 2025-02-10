package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	pokecache := NewCache(interval)

	for _, c := range cases {
		pokecache.Add(c.key, c.val)

		if len(pokecache.cache) == 0 {
			t.Errorf("expected cache length to be %d but go length of zero", len(pokecache.cache))
			return
		}

		bytes, ok := pokecache.Get(c.key)
		if !ok {
			t.Errorf("expected to find %s but got nothing", c.key)
			return
		}

		if string(bytes) != string(c.val) {
			t.Errorf("expected to find byte of content %s but got nothing", string(bytes))
			return
		}

	}

}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Second
	const waitTime = baseTime + 5*time.Millisecond
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		// {
		// 	key: "https://example.com/path",
		// 	val: []byte("moretestdata"),
		// },
	}

	for _, c := range cases {
		pokecache := NewCache(baseTime)
		pokecache.Add(c.key, c.val)

		_, ok := pokecache.Get(c.key)
		if !ok {
			t.Errorf("expected to find %s but got nothing", c.key)
			return
		}

		time.Sleep(waitTime)

		_, ok = pokecache.Get(c.key)

		if ok {
			t.Errorf("expected to find nothing but got %s", c.key)
			return
		}
	}
}

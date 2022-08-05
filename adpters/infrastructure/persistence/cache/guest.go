package cache

import (
	"errors"
	"sync"
)

type Cache struct {
	Item map[string]*Item
	sync.Mutex
}

type Item struct {
	Object interface{}
}

func (c *Cache) Get(key string) (*Item, error) {

	rsp, found := c.Item[key]
	if !found {
		return nil, errors.New("Not Found key")
	}

	return rsp, nil
}

func (c *Cache) Set(key string, x interface{}) {
	c.Lock()
	defer c.Unlock()

	if key == "" {
		return
	}

	c.Item[key] = &Item{Object: x}
}

func (c *Cache) Delete(key string) {
	c.Lock()
	defer c.Unlock()

	delete(c.Item, key)
}

func (c *Cache) Purge() {
	c.Lock()
	defer c.Unlock()

	c.Item = make(map[string]*Item)
}

func (c *Cache) Keys() []string {
	c.Lock()
	defer c.Unlock()
	keys := []string{}
	for key := range c.Item {
		keys = append(keys, key)
	}
	return keys
}

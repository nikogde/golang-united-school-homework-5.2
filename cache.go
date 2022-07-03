package cache

import "time"

type Dictionary struct {
	Value      string
	Deadline   time.Time
	CanExpired bool
}

type Cache struct {
	Key map[string]Dictionary
}

func NewCache() Cache {
	return Cache{Key: map[string]Dictionary{}}
}

func (c *Cache) Put(key, value string) {
	c.Key[key] = Dictionary{
		Value:      value,
		CanExpired: false,
	}
}

func (c Cache) Keys() []string {
	var keys []string

	for key, value := range c.Key {
		if !value.CanExpired || (value.CanExpired && value.Deadline.After(time.Now())) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c Cache) Get(key string) (string, bool) {
	dict := make(map[string]string)

	for _, val := range c.Keys() {
		dict[val] = c.Key[val].Value
	}

	k, ok := dict[key]
	return k, ok
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.Key[key] = Dictionary{
		Value:      value,
		Deadline:   deadline,
		CanExpired: true,
	}
}

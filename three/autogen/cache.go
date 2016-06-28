package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Cache represents a cache.
type Cache struct{ p *js.Object }

// Cache returns a cache object.
func (t *Three) Cache() *Cache {
	p := t.ctx.Get("Cache")
	return &Cache{p: p}
}

// NewCache returns a new cache object.
func (t *Cache) New() *Cache {
	p := t.p.New()
	return &Cache{p: p}
}

// Add TODO description.
func (c *Cache) Add(key, file float64) *Cache {
	c.p.Call("add", key, file)
	return c
}

// Get TODO description.
func (c *Cache) Get(key float64) *Cache {
	c.p.Call("get", key)
	return c
}

// Remove TODO description.
func (c *Cache) Remove(key float64) *Cache {
	c.p.Call("remove", key)
	return c
}

// Clear TODO description.
func (c *Cache) Clear() *Cache {
	c.p.Call("clear")
	return c
}


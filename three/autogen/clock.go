package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Clock represents a clock.
type Clock struct{ p *js.Object }

// Clock returns a Clock object.
func (t *Three) Clock() *Clock {
	p := t.ctx.Get("Clock")
	return &Clock{p: p}
}

// New returns a new Clock object.
func (t *Clock) New(autoStart float64) *Clock {
	p := t.p.New(autoStart)
	return &Clock{p: p}
}

// Start TODO description.
func (c *Clock) Start() *Clock {
	c.p.Call("start")
	return c
}

// Stop TODO description.
func (c *Clock) Stop() *Clock {
	c.p.Call("stop")
	return c
}

// GetElapsedTime TODO description.
func (c *Clock) GetElapsedTime() *Clock {
	c.p.Call("getElapsedTime")
	return c
}

// GetDelta TODO description.
func (c *Clock) GetDelta() *Clock {
	c.p.Call("getDelta")
	return c
}


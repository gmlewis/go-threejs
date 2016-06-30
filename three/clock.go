// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Clock represents a clock.
type Clock struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *Clock) JSObject() *js.Object { return c.p }

// Clock returns a Clock object.
func (t *Three) Clock() *Clock {
	p := t.ctx.Get("Clock")
	return &Clock{p: p}
}

// New returns a new Clock object.
func (c *Clock) New(autoStart float64) *Clock {
	p := c.p.New(autoStart)
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

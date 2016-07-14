// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import "github.com/gopherjs/gopherjs/js"

// Clock is an object for keeping track of time.
//
// http://threejs.org/docs/index.html#Reference/Core/Clock
type Clock struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (c *Clock) JSObject() *js.Object { return c.p }

// Clock returns a Clock JavaScript class.
func (t *Three) Clock() *Clock {
	p := t.ctx.Get("Clock")
	return ClockFromJSObject(p)
}

// ClockFromJSObject returns a wrapped Clock JavaScript class.
func ClockFromJSObject(p *js.Object) *Clock {
	return &Clock{p: p}
}

// NewClock returns a new Clock object.
func (t *Three) NewClock(autoStart bool) *Clock {
	p := t.ctx.Get("Clock").New(autoStart)
	return ClockFromJSObject(p)
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
func (c *Clock) GetElapsedTime() float64 {
	return c.p.Call("getElapsedTime").Float()
}

// GetDelta TODO description.
func (c *Clock) GetDelta() float64 {
	return c.p.Call("getDelta").Float()
}

// Running returns the property of the same name.
func (c *Clock) Running() bool {
	return c.p.Get("running").Bool()
}

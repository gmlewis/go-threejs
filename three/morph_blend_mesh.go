// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MorphBlendMesh represents a morphblendmesh.
type MorphBlendMesh struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (m *MorphBlendMesh) JSObject() *js.Object { return m.p }

// MorphBlendMesh returns a MorphBlendMesh JavaScript class.
func (t *Three) MorphBlendMesh() *MorphBlendMesh {
	p := t.ctx.Get("MorphBlendMesh")
	return &MorphBlendMesh{p: p}
}

// NewMorphBlendMesh returns a new MorphBlendMesh object.
func (t *Three) NewMorphBlendMesh(geometry, material float64) *MorphBlendMesh {
	p := t.ctx.Get("MorphBlendMesh").New(geometry, material)
	return &MorphBlendMesh{p: p}
}

// CreateAnimation TODO description.
func (m *MorphBlendMesh) CreateAnimation(name, start, end, fps float64) *MorphBlendMesh {
	m.p.Call("createAnimation", name, start, end, fps)
	return m
}

// AutoCreateAnimations TODO description.
func (m *MorphBlendMesh) AutoCreateAnimations(fps float64) *MorphBlendMesh {
	m.p.Call("autoCreateAnimations", fps)
	return m
}

// SetAnimationDirectionForward TODO description.
func (m *MorphBlendMesh) SetAnimationDirectionForward(name float64) *MorphBlendMesh {
	m.p.Call("setAnimationDirectionForward", name)
	return m
}

// SetAnimationDirectionBackward TODO description.
func (m *MorphBlendMesh) SetAnimationDirectionBackward(name float64) *MorphBlendMesh {
	m.p.Call("setAnimationDirectionBackward", name)
	return m
}

// SetAnimationFPS TODO description.
func (m *MorphBlendMesh) SetAnimationFPS(name, fps float64) *MorphBlendMesh {
	m.p.Call("setAnimationFPS", name, fps)
	return m
}

// SetAnimationDuration TODO description.
func (m *MorphBlendMesh) SetAnimationDuration(name, duration float64) *MorphBlendMesh {
	m.p.Call("setAnimationDuration", name, duration)
	return m
}

// SetAnimationWeight TODO description.
func (m *MorphBlendMesh) SetAnimationWeight(name, weight float64) *MorphBlendMesh {
	m.p.Call("setAnimationWeight", name, weight)
	return m
}

// SetAnimationTime TODO description.
func (m *MorphBlendMesh) SetAnimationTime(name, time float64) *MorphBlendMesh {
	m.p.Call("setAnimationTime", name, time)
	return m
}

// GetAnimationTime TODO description.
func (m *MorphBlendMesh) GetAnimationTime(name float64) *MorphBlendMesh {
	m.p.Call("getAnimationTime", name)
	return m
}

// GetAnimationDuration TODO description.
func (m *MorphBlendMesh) GetAnimationDuration(name float64) *MorphBlendMesh {
	m.p.Call("getAnimationDuration", name)
	return m
}

// PlayAnimation TODO description.
func (m *MorphBlendMesh) PlayAnimation(name float64) *MorphBlendMesh {
	m.p.Call("playAnimation", name)
	return m
}

// StopAnimation TODO description.
func (m *MorphBlendMesh) StopAnimation(name float64) *MorphBlendMesh {
	m.p.Call("stopAnimation", name)
	return m
}

// Update TODO description.
func (m *MorphBlendMesh) Update(delta float64) *MorphBlendMesh {
	m.p.Call("update", delta)
	return m
}

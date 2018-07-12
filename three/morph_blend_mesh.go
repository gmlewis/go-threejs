// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	return MorphBlendMeshFromJSObject(p)
}

// MorphBlendMeshFromJSObject returns a wrapped MorphBlendMesh JavaScript class.
func MorphBlendMeshFromJSObject(p *js.Object) *MorphBlendMesh {
	return &MorphBlendMesh{p: p}
}

// NewMorphBlendMesh returns a new MorphBlendMesh object.
func (t *Three) NewMorphBlendMesh(geometry, material float64) *MorphBlendMesh {
	p := t.ctx.Get("MorphBlendMesh").New(geometry, material)
	return MorphBlendMeshFromJSObject(p)
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

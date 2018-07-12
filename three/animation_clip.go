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

// AnimationClip represents an animationclip.
type AnimationClip struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AnimationClip) JSObject() *js.Object { return a.p }

// AnimationClip returns an AnimationClip JavaScript class.
func (t *Three) AnimationClip() *AnimationClip {
	p := t.ctx.Get("AnimationClip")
	return AnimationClipFromJSObject(p)
}

// AnimationClipFromJSObject returns a wrapped AnimationClip JavaScript class.
func AnimationClipFromJSObject(p *js.Object) *AnimationClip {
	return &AnimationClip{p: p}
}

// NewAnimationClip returns a new AnimationClip object.
func (t *Three) NewAnimationClip(name, duration, tracks float64) *AnimationClip {
	p := t.ctx.Get("AnimationClip").New(name, duration, tracks)
	return AnimationClipFromJSObject(p)
}

// ResetDuration TODO description.
func (a *AnimationClip) ResetDuration() *AnimationClip {
	a.p.Call("resetDuration")
	return a
}

// Trim TODO description.
func (a *AnimationClip) Trim() *AnimationClip {
	a.p.Call("trim")
	return a
}

// Optimize TODO description.
func (a *AnimationClip) Optimize() *AnimationClip {
	a.p.Call("optimize")
	return a
}

// Parse TODO description.
func (a *AnimationClip) Parse(json float64) *AnimationClip {
	a.p.Call("parse", json)
	return a
}

// ToJSON TODO description.
func (a *AnimationClip) ToJSON(clip float64) *AnimationClip {
	a.p.Call("toJSON", clip)
	return a
}

// CreateFromMorphTargetSequence TODO description.
func (a *AnimationClip) CreateFromMorphTargetSequence(name, morphTargetSequence, fps float64) *AnimationClip {
	a.p.Call("CreateFromMorphTargetSequence", name, morphTargetSequence, fps)
	return a
}

// FindByName TODO description.
func (a *AnimationClip) FindByName(clipArray, name float64) *AnimationClip {
	a.p.Call("findByName", clipArray, name)
	return a
}

// CreateClipsFromMorphTargetSequences TODO description.
func (a *AnimationClip) CreateClipsFromMorphTargetSequences(morphTargets, fps float64) *AnimationClip {
	a.p.Call("CreateClipsFromMorphTargetSequences", morphTargets, fps)
	return a
}

// ParseAnimation TODO description.
func (a *AnimationClip) ParseAnimation(animation, bones, nodeName float64) *AnimationClip {
	a.p.Call("parseAnimation", animation, bones, nodeName)
	return a
}

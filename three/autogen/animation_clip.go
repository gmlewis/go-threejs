package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AnimationClip represents an animationclip.
type AnimationClip struct{ p *js.Object }

// AnimationClip returns an AnimationClip object.
func (t *Three) AnimationClip() *AnimationClip {
	p := t.ctx.Get("AnimationClip")
	return &AnimationClip{p: p}
}

// New returns a new AnimationClip object.
func (t *AnimationClip) New(name, duration, tracks float64) *AnimationClip {
	p := t.p.New(name, duration, tracks)
	return &AnimationClip{p: p}
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


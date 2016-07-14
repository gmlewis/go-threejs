// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AnimationMixer represents an animationmixer.
type AnimationMixer struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (a *AnimationMixer) JSObject() *js.Object { return a.p }

// AnimationMixer returns an AnimationMixer JavaScript class.
func (t *Three) AnimationMixer() *AnimationMixer {
	p := t.ctx.Get("AnimationMixer")
	return AnimationMixerFromJSObject(p)
}

// AnimationMixerFromJSObject returns a wrapped AnimationMixer JavaScript class.
func AnimationMixerFromJSObject(p *js.Object) *AnimationMixer {
	return &AnimationMixer{p: p}
}

// NewAnimationMixer returns a new AnimationMixer object.
func (t *Three) NewAnimationMixer(root float64) *AnimationMixer {
	p := t.ctx.Get("AnimationMixer").New(root)
	return AnimationMixerFromJSObject(p)
}

// Total returns the total-component of the AnimationMixer.
func (a *AnimationMixer) Total() float64 {
	return a.p.Get("total").Float()
}

// InUse returns the inUse-component of the AnimationMixer.
func (a *AnimationMixer) InUse() float64 {
	return a.p.Get("inUse").Float()
}

// ClipAction TODO description.
func (a *AnimationMixer) ClipAction(clip, optionalRoot float64) *AnimationMixer {
	a.p.Call("clipAction", clip, optionalRoot)
	return a
}

// ExistingAction TODO description.
func (a *AnimationMixer) ExistingAction(clip, optionalRoot float64) *AnimationMixer {
	a.p.Call("existingAction", clip, optionalRoot)
	return a
}

// StopAllAction TODO description.
func (a *AnimationMixer) StopAllAction() *AnimationMixer {
	a.p.Call("stopAllAction")
	return a
}

// Update TODO description.
func (a *AnimationMixer) Update(deltaTime float64) *AnimationMixer {
	a.p.Call("update", deltaTime)
	return a
}

// GetRoot TODO description.
func (a *AnimationMixer) GetRoot() *AnimationMixer {
	a.p.Call("getRoot")
	return a
}

// UncacheClip TODO description.
func (a *AnimationMixer) UncacheClip(clip float64) *AnimationMixer {
	a.p.Call("uncacheClip", clip)
	return a
}

// UncacheRoot TODO description.
func (a *AnimationMixer) UncacheRoot(root float64) *AnimationMixer {
	a.p.Call("uncacheRoot", root)
	return a
}

// UncacheAction TODO description.
func (a *AnimationMixer) UncacheAction(clip, optionalRoot float64) *AnimationMixer {
	a.p.Call("uncacheAction", clip, optionalRoot)
	return a
}

// Play TODO description.
func (a *AnimationMixer) Play() *AnimationMixer {
	a.p.Call("play")
	return a
}

// Stop TODO description.
func (a *AnimationMixer) Stop() *AnimationMixer {
	a.p.Call("stop")
	return a
}

// Reset TODO description.
func (a *AnimationMixer) Reset() *AnimationMixer {
	a.p.Call("reset")
	return a
}

// IsRunning TODO description.
func (a *AnimationMixer) IsRunning() *AnimationMixer {
	a.p.Call("isRunning")
	return a
}

// IsScheduled TODO description.
func (a *AnimationMixer) IsScheduled() *AnimationMixer {
	a.p.Call("isScheduled")
	return a
}

// StartAt TODO description.
func (a *AnimationMixer) StartAt(time float64) *AnimationMixer {
	a.p.Call("startAt", time)
	return a
}

// SetLoop TODO description.
func (a *AnimationMixer) SetLoop(mode, repetitions float64) *AnimationMixer {
	a.p.Call("setLoop", mode, repetitions)
	return a
}

// SetEffectiveWeight TODO description.
func (a *AnimationMixer) SetEffectiveWeight(weight float64) *AnimationMixer {
	a.p.Call("setEffectiveWeight", weight)
	return a
}

// GetEffectiveWeight TODO description.
func (a *AnimationMixer) GetEffectiveWeight() *AnimationMixer {
	a.p.Call("getEffectiveWeight")
	return a
}

// FadeIn TODO description.
func (a *AnimationMixer) FadeIn(duration float64) *AnimationMixer {
	a.p.Call("fadeIn", duration)
	return a
}

// FadeOut TODO description.
func (a *AnimationMixer) FadeOut(duration float64) *AnimationMixer {
	a.p.Call("fadeOut", duration)
	return a
}

// CrossFadeFrom TODO description.
func (a *AnimationMixer) CrossFadeFrom(fadeOutAction, duration, warp float64) *AnimationMixer {
	a.p.Call("crossFadeFrom", fadeOutAction, duration, warp)
	return a
}

// CrossFadeTo TODO description.
func (a *AnimationMixer) CrossFadeTo(fadeInAction, duration, warp float64) *AnimationMixer {
	a.p.Call("crossFadeTo", fadeInAction, duration, warp)
	return a
}

// StopFading TODO description.
func (a *AnimationMixer) StopFading() *AnimationMixer {
	a.p.Call("stopFading")
	return a
}

// SetEffectiveTimeScale TODO description.
func (a *AnimationMixer) SetEffectiveTimeScale(timeScale float64) *AnimationMixer {
	a.p.Call("setEffectiveTimeScale", timeScale)
	return a
}

// GetEffectiveTimeScale TODO description.
func (a *AnimationMixer) GetEffectiveTimeScale() *AnimationMixer {
	a.p.Call("getEffectiveTimeScale")
	return a
}

// SetDuration TODO description.
func (a *AnimationMixer) SetDuration(duration float64) *AnimationMixer {
	a.p.Call("setDuration", duration)
	return a
}

// SyncWith TODO description.
func (a *AnimationMixer) SyncWith(action float64) *AnimationMixer {
	a.p.Call("syncWith", action)
	return a
}

// Halt TODO description.
func (a *AnimationMixer) Halt(duration float64) *AnimationMixer {
	a.p.Call("halt", duration)
	return a
}

// Warp TODO description.
func (a *AnimationMixer) Warp(startTimeScale, endTimeScale, duration float64) *AnimationMixer {
	a.p.Call("warp", startTimeScale, endTimeScale, duration)
	return a
}

// StopWarping TODO description.
func (a *AnimationMixer) StopWarping() *AnimationMixer {
	a.p.Call("stopWarping")
	return a
}

// GetMixer TODO description.
func (a *AnimationMixer) GetMixer() *AnimationMixer {
	a.p.Call("getMixer")
	return a
}

// GetClip TODO description.
func (a *AnimationMixer) GetClip() *AnimationMixer {
	a.p.Call("getClip")
	return a
}

// _update TODO description.
func (a *AnimationMixer) _update(time, deltaTime, timeDirection, accuIndex float64) *AnimationMixer {
	a.p.Call("_update", time, deltaTime, timeDirection, accuIndex)
	return a
}

// _updateWeight TODO description.
func (a *AnimationMixer) _updateWeight(time float64) *AnimationMixer {
	a.p.Call("_updateWeight", time)
	return a
}

// _updateTimeScale TODO description.
func (a *AnimationMixer) _updateTimeScale(time float64) *AnimationMixer {
	a.p.Call("_updateTimeScale", time)
	return a
}

// _updateTime TODO description.
func (a *AnimationMixer) _updateTime(deltaTime float64) *AnimationMixer {
	a.p.Call("_updateTime", deltaTime)
	return a
}

// _setEndings TODO description.
func (a *AnimationMixer) _setEndings(atStart, atEnd, pingPong float64) *AnimationMixer {
	a.p.Call("_setEndings", atStart, atEnd, pingPong)
	return a
}

// _scheduleFading TODO description.
func (a *AnimationMixer) _scheduleFading(duration, weightNow, weightThen float64) *AnimationMixer {
	a.p.Call("_scheduleFading", duration, weightNow, weightThen)
	return a
}

// _bindAction TODO description.
func (a *AnimationMixer) _bindAction(action, prototypAction float64) *AnimationMixer {
	a.p.Call("_bindAction", action, prototypAction)
	return a
}

// _activateAction TODO description.
func (a *AnimationMixer) _activateAction(action float64) *AnimationMixer {
	a.p.Call("_activateAction", action)
	return a
}

// _deactivateAction TODO description.
func (a *AnimationMixer) _deactivateAction(action float64) *AnimationMixer {
	a.p.Call("_deactivateAction", action)
	return a
}

// _initMemoryManager TODO description.
func (a *AnimationMixer) _initMemoryManager() *AnimationMixer {
	a.p.Call("_initMemoryManager")
	return a
}

// _isActiveAction TODO description.
func (a *AnimationMixer) _isActiveAction(action float64) *AnimationMixer {
	a.p.Call("_isActiveAction", action)
	return a
}

// _addInactiveAction TODO description.
func (a *AnimationMixer) _addInactiveAction(action, clipName, rootUUID float64) *AnimationMixer {
	a.p.Call("_addInactiveAction", action, clipName, rootUUID)
	return a
}

// _removeInactiveAction TODO description.
func (a *AnimationMixer) _removeInactiveAction(action float64) *AnimationMixer {
	a.p.Call("_removeInactiveAction", action)
	return a
}

// _removeInactiveBindingsForAction TODO description.
func (a *AnimationMixer) _removeInactiveBindingsForAction(action float64) *AnimationMixer {
	a.p.Call("_removeInactiveBindingsForAction", action)
	return a
}

// _lendAction TODO description.
func (a *AnimationMixer) _lendAction(action float64) *AnimationMixer {
	a.p.Call("_lendAction", action)
	return a
}

// _takeBackAction TODO description.
func (a *AnimationMixer) _takeBackAction(action float64) *AnimationMixer {
	a.p.Call("_takeBackAction", action)
	return a
}

// _addInactiveBinding TODO description.
func (a *AnimationMixer) _addInactiveBinding(binding, rootUUID, trackName float64) *AnimationMixer {
	a.p.Call("_addInactiveBinding", binding, rootUUID, trackName)
	return a
}

// _removeInactiveBinding TODO description.
func (a *AnimationMixer) _removeInactiveBinding(binding float64) *AnimationMixer {
	a.p.Call("_removeInactiveBinding", binding)
	return a
}

// _lendBinding TODO description.
func (a *AnimationMixer) _lendBinding(binding float64) *AnimationMixer {
	a.p.Call("_lendBinding", binding)
	return a
}

// _takeBackBinding TODO description.
func (a *AnimationMixer) _takeBackBinding(binding float64) *AnimationMixer {
	a.p.Call("_takeBackBinding", binding)
	return a
}

// _lendControlInterpolant TODO description.
func (a *AnimationMixer) _lendControlInterpolant() *AnimationMixer {
	a.p.Call("_lendControlInterpolant")
	return a
}

// _takeBackControlInterpolant TODO description.
func (a *AnimationMixer) _takeBackControlInterpolant(interpolant float64) *AnimationMixer {
	a.p.Call("_takeBackControlInterpolant", interpolant)
	return a
}

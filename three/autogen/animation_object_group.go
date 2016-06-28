package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AnimationObjectGroup represents an animationobjectgroup.
type AnimationObjectGroup struct{ p *js.Object }

// AnimationObjectGroup returns an animationobjectgroup object.
func (t *Three) AnimationObjectGroup() *AnimationObjectGroup {
	p := t.ctx.Get("AnimationObjectGroup")
	return &AnimationObjectGroup{p: p}
}

// NewAnimationObjectGroup returns a new animationobjectgroup object.
func (t *AnimationObjectGroup) New(var_args float64) *AnimationObjectGroup {
	p := t.p.New(var_args)
	return &AnimationObjectGroup{p: p}
}

// Total returns the total-component of the AnimationObjectGroup.
func (a *AnimationObjectGroup) Total() float64 {
	return a.p.Get("total").Float()
}

// InUse returns the inUse-component of the AnimationObjectGroup.
func (a *AnimationObjectGroup) InUse() float64 {
	return a.p.Get("inUse").Float()
}

// BindingsPerObject returns the bindingsPerObject-component of the AnimationObjectGroup.
func (a *AnimationObjectGroup) BindingsPerObject() float64 {
	return a.p.Get("bindingsPerObject").Float()
}

// Add TODO description.
func (a *AnimationObjectGroup) Add(var_args float64) *AnimationObjectGroup {
	a.p.Call("add", var_args)
	return a
}

// Remove TODO description.
func (a *AnimationObjectGroup) Remove(var_args float64) *AnimationObjectGroup {
	a.p.Call("remove", var_args)
	return a
}

// Uncache TODO description.
func (a *AnimationObjectGroup) Uncache(var_args float64) *AnimationObjectGroup {
	a.p.Call("uncache", var_args)
	return a
}

// Subscribe_ TODO description.
func (a *AnimationObjectGroup) Subscribe_(path, parsedPath float64) *AnimationObjectGroup {
	a.p.Call("subscribe_", path, parsedPath)
	return a
}

// Unsubscribe_ TODO description.
func (a *AnimationObjectGroup) Unsubscribe_(path float64) *AnimationObjectGroup {
	a.p.Call("unsubscribe_", path)
	return a
}


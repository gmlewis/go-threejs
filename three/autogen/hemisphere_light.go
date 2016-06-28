package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// HemisphereLight represents a hemispherelight.
type HemisphereLight struct{ p *js.Object }

// HemisphereLight returns a hemispherelight object.
func (t *Three) HemisphereLight() *HemisphereLight {
	p := t.ctx.Get("HemisphereLight")
	return &HemisphereLight{p: p}
}

// NewHemisphereLight returns a new hemispherelight object.
func (t *HemisphereLight) New(skyColor, groundColor, intensity float64) *HemisphereLight {
	p := t.p.New(skyColor, groundColor, intensity)
	return &HemisphereLight{p: p}
}

// Copy TODO description.
func (h *HemisphereLight) Copy(source float64) *HemisphereLight {
	h.p.Call("copy", source)
	return h
}


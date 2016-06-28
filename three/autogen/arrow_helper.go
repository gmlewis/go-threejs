package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ArrowHelper represents an arrowhelper.
type ArrowHelper struct{ p *js.Object }

// ArrowHelper returns an arrowhelper object.
func (t *Three) ArrowHelper() *ArrowHelper {
	p := t.ctx.Get("ArrowHelper")
	return &ArrowHelper{p: p}
}

// NewArrowHelper returns a new arrowhelper object.
func (t *ArrowHelper) New(dir, origin, length, color, headLength, headWidth float64) *ArrowHelper {
	p := t.p.New(dir, origin, length, color, headLength, headWidth)
	return &ArrowHelper{p: p}
}

// SetLength TODO description.
func (a *ArrowHelper) SetLength(length, headLength, headWidth float64) *ArrowHelper {
	a.p.Call("setLength", length, headLength, headWidth)
	return a
}

// SetColor TODO description.
func (a *ArrowHelper) SetColor(color float64) *ArrowHelper {
	a.p.Call("setColor", color)
	return a
}

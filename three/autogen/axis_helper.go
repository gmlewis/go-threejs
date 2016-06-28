package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AxisHelper represents an axishelper.
type AxisHelper struct{ p *js.Object }

// AxisHelper returns an axishelper object.
func (t *Three) AxisHelper() *AxisHelper {
	p := t.ctx.Get("AxisHelper")
	return &AxisHelper{p: p}
}

// NewAxisHelper returns a new axishelper object.
func (t *AxisHelper) New(size float64) *AxisHelper {
	p := t.p.New(size)
	return &AxisHelper{p: p}
}


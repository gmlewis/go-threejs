package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LineCurve represents a linecurve.
type LineCurve struct{ p *js.Object }

// LineCurve returns a LineCurve object.
func (t *Three) LineCurve() *LineCurve {
	p := t.ctx.Get("LineCurve")
	return &LineCurve{p: p}
}

// New returns a new LineCurve object.
func (t *LineCurve) New(v1, v2 float64) *LineCurve {
	p := t.p.New(v1, v2)
	return &LineCurve{p: p}
}

// GetPoint TODO description.
func (l *LineCurve) GetPoint(t float64) *LineCurve {
	l.p.Call("getPoint", t)
	return l
}

// GetPointAt TODO description.
func (l *LineCurve) GetPointAt(u float64) *LineCurve {
	l.p.Call("getPointAt", u)
	return l
}

// GetTangent TODO description.
func (l *LineCurve) GetTangent(t float64) *LineCurve {
	l.p.Call("getTangent", t)
	return l
}


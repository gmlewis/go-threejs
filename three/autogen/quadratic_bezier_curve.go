package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// QuadraticBezierCurve represents a quadraticbeziercurve.
type QuadraticBezierCurve struct{ p *js.Object }

// QuadraticBezierCurve returns a quadraticbeziercurve object.
func (t *Three) QuadraticBezierCurve() *QuadraticBezierCurve {
	p := t.ctx.Get("QuadraticBezierCurve")
	return &QuadraticBezierCurve{p: p}
}

// NewQuadraticBezierCurve returns a new quadraticbeziercurve object.
func (t *QuadraticBezierCurve) New(v0, v1, v2 float64) *QuadraticBezierCurve {
	p := t.p.New(v0, v1, v2)
	return &QuadraticBezierCurve{p: p}
}

// GetPoint TODO description.
func (q *QuadraticBezierCurve) GetPoint(t float64) *QuadraticBezierCurve {
	q.p.Call("getPoint", t)
	return q
}

// GetTangent TODO description.
func (q *QuadraticBezierCurve) GetTangent(t float64) *QuadraticBezierCurve {
	q.p.Call("getTangent", t)
	return q
}


package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// EllipseCurve represents an ellipsecurve.
type EllipseCurve struct{ p *js.Object }

// EllipseCurve returns an ellipsecurve object.
func (t *Three) EllipseCurve() *EllipseCurve {
	p := t.ctx.Get("EllipseCurve")
	return &EllipseCurve{p: p}
}

// NewEllipseCurve returns a new ellipsecurve object.
func (t *EllipseCurve) New(aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation float64) *EllipseCurve {
	p := t.p.New(aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
	return &EllipseCurve{p: p}
}

// GetPoint TODO description.
func (e *EllipseCurve) GetPoint(t float64) *EllipseCurve {
	e.p.Call("getPoint", t)
	return e
}


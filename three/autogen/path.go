package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Path represents a path.
type Path struct{ p *js.Object }

// Path returns a path object.
func (t *Three) Path() *Path {
	p := t.ctx.Get("Path")
	return &Path{p: p}
}

// NewPath returns a new path object.
func (t *Path) New(points float64) *Path {
	p := t.p.New(points)
	return &Path{p: p}
}

// FromPoints TODO description.
func (p *Path) FromPoints(vectors float64) *Path {
	p.p.Call("fromPoints", vectors)
	return p
}

// MoveTo TODO description.
func (p *Path) MoveTo(x, y float64) *Path {
	p.p.Call("moveTo", x, y)
	return p
}

// LineTo TODO description.
func (p *Path) LineTo(x, y float64) *Path {
	p.p.Call("lineTo", x, y)
	return p
}

// QuadraticCurveTo TODO description.
func (p *Path) QuadraticCurveTo(aCPx, aCPy, aX, aY float64) *Path {
	p.p.Call("quadraticCurveTo", aCPx, aCPy, aX, aY)
	return p
}

// BezierCurveTo TODO description.
func (p *Path) BezierCurveTo(aCP1x, aCP1y, aCP2x, aCP2y, aX, aY float64) *Path {
	p.p.Call("bezierCurveTo", aCP1x, aCP1y, aCP2x, aCP2y, aX, aY)
	return p
}

// SplineThru TODO description.
func (p *Path) SplineThru(pts /*Array of Vector*/ float64) *Path {
	p.p.Call("splineThru", pts /*Array of Vector*/)
	return p
}

// Arc TODO description.
func (p *Path) Arc(aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise float64) *Path {
	p.p.Call("arc", aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
	return p
}

// Absarc TODO description.
func (p *Path) Absarc(aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise float64) *Path {
	p.p.Call("absarc", aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
	return p
}

// Ellipse TODO description.
func (p *Path) Ellipse(aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation float64) *Path {
	p.p.Call("ellipse", aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
	return p
}

// Absellipse TODO description.
func (p *Path) Absellipse(aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation float64) *Path {
	p.p.Call("absellipse", aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
	return p
}

// GetSpacedPoints TODO description.
func (p *Path) GetSpacedPoints(divisions float64) *Path {
	p.p.Call("getSpacedPoints", divisions)
	return p
}

// GetPoints TODO description.
func (p *Path) GetPoints(divisions float64) *Path {
	p.p.Call("getPoints", divisions)
	return p
}

// ToShapes TODO description.
func (p *Path) ToShapes(isCCW, noHoles float64) *Path {
	p.p.Call("toShapes", isCCW, noHoles)
	return p
}


package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// TubeGeometry represents a tubegeometry.
type TubeGeometry struct{ p *js.Object }

// TubeGeometry returns a tubegeometry object.
func (t *Three) TubeGeometry() *TubeGeometry {
	p := t.ctx.Get("TubeGeometry")
	return &TubeGeometry{p: p}
}

// NewTubeGeometry returns a new tubegeometry object.
func (t *TubeGeometry) New(path, segments, radius, radialSegments, closed, taper float64) *TubeGeometry {
	p := t.p.New(path, segments, radius, radialSegments, closed, taper)
	return &TubeGeometry{p: p}
}


package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DirectGeometry represents a directgeometry.
type DirectGeometry struct{ p *js.Object }

// DirectGeometry returns a directgeometry object.
func (t *Three) DirectGeometry() *DirectGeometry {
	p := t.ctx.Get("DirectGeometry")
	return &DirectGeometry{p: p}
}

// NewDirectGeometry returns a new directgeometry object.
func (t *DirectGeometry) New() *DirectGeometry {
	p := t.p.New()
	return &DirectGeometry{p: p}
}

// ComputeFaceNormals TODO description.
func (d *DirectGeometry) ComputeFaceNormals() *DirectGeometry {
	d.p.Call("computeFaceNormals")
	return d
}

// ComputeVertexNormals TODO description.
func (d *DirectGeometry) ComputeVertexNormals() *DirectGeometry {
	d.p.Call("computeVertexNormals")
	return d
}

// ComputeGroups TODO description.
func (d *DirectGeometry) ComputeGroups(geometry float64) *DirectGeometry {
	d.p.Call("computeGroups", geometry)
	return d
}

// FromGeometry TODO description.
func (d *DirectGeometry) FromGeometry(geometry float64) *DirectGeometry {
	d.p.Call("fromGeometry", geometry)
	return d
}

// Dispose TODO description.
func (d *DirectGeometry) Dispose() *DirectGeometry {
	d.p.Call("dispose")
	return d
}

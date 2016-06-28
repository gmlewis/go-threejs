package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BoundingBoxHelper represents a boundingboxhelper.
type BoundingBoxHelper struct{ p *js.Object }

// BoundingBoxHelper returns a boundingboxhelper object.
func (t *Three) BoundingBoxHelper() *BoundingBoxHelper {
	p := t.ctx.Get("BoundingBoxHelper")
	return &BoundingBoxHelper{p: p}
}

// NewBoundingBoxHelper returns a new boundingboxhelper object.
func (t *BoundingBoxHelper) New(object, hex float64) *BoundingBoxHelper {
	p := t.p.New(object, hex)
	return &BoundingBoxHelper{p: p}
}


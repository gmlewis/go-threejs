package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// FaceNormalsHelper represents a facenormalshelper.
type FaceNormalsHelper struct{ p *js.Object }

// FaceNormalsHelper returns a FaceNormalsHelper object.
func (t *Three) FaceNormalsHelper() *FaceNormalsHelper {
	p := t.ctx.Get("FaceNormalsHelper")
	return &FaceNormalsHelper{p: p}
}

// New returns a new FaceNormalsHelper object.
func (t *FaceNormalsHelper) New(object, size, hex, linewidth float64) *FaceNormalsHelper {
	p := t.p.New(object, size, hex, linewidth)
	return &FaceNormalsHelper{p: p}
}


package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// FaceNormalsHelper represents a facenormalshelper.
type FaceNormalsHelper struct{ p *js.Object }

// FaceNormalsHelper returns a facenormalshelper object.
func (t *Three) FaceNormalsHelper() *FaceNormalsHelper {
	p := t.ctx.Get("FaceNormalsHelper")
	return &FaceNormalsHelper{p: p}
}

// NewFaceNormalsHelper returns a new facenormalshelper object.
func (t *FaceNormalsHelper) New(object, size, hex, linewidth float64) *FaceNormalsHelper {
	p := t.p.New(object, size, hex, linewidth)
	return &FaceNormalsHelper{p: p}
}


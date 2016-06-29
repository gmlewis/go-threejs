package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CanvasTexture represents a canvastexture.
type CanvasTexture struct{ p *js.Object }

// CanvasTexture returns a CanvasTexture object.
func (t *Three) CanvasTexture() *CanvasTexture {
	p := t.ctx.Get("CanvasTexture")
	return &CanvasTexture{p: p}
}

// New returns a new CanvasTexture object.
func (t *CanvasTexture) New(canvas, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *CanvasTexture {
	p := t.p.New(canvas, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return &CanvasTexture{p: p}
}


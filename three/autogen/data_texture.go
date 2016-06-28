package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// DataTexture represents a datatexture.
type DataTexture struct{ p *js.Object }

// DataTexture returns a datatexture object.
func (t *Three) DataTexture() *DataTexture {
	p := t.ctx.Get("DataTexture")
	return &DataTexture{p: p}
}

// NewDataTexture returns a new datatexture object.
func (t *DataTexture) New(data, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy float64) *DataTexture {
	p := t.p.New(data, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy)
	return &DataTexture{p: p}
}

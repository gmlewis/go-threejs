package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CompressedTexture represents a compressedtexture.
type CompressedTexture struct{ p *js.Object }

// CompressedTexture returns a compressedtexture object.
func (t *Three) CompressedTexture() *CompressedTexture {
	p := t.ctx.Get("CompressedTexture")
	return &CompressedTexture{p: p}
}

// NewCompressedTexture returns a new compressedtexture object.
func (t *CompressedTexture) New(mipmaps, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy float64) *CompressedTexture {
	p := t.p.New(mipmaps, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy)
	return &CompressedTexture{p: p}
}

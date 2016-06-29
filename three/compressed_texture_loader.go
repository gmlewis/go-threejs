package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CompressedTextureLoader represents a compressedtextureloader.
type CompressedTextureLoader struct{ p *js.Object }

// CompressedTextureLoader returns a CompressedTextureLoader object.
func (t *Three) CompressedTextureLoader() *CompressedTextureLoader {
	p := t.ctx.Get("CompressedTextureLoader")
	return &CompressedTextureLoader{p: p}
}

// New returns a new CompressedTextureLoader object.
func (t *CompressedTextureLoader) New(manager float64) *CompressedTextureLoader {
	p := t.p.New(manager)
	return &CompressedTextureLoader{p: p}
}

// Load TODO description.
func (c *CompressedTextureLoader) Load(url, onLoad, onProgress, onError float64) *CompressedTextureLoader {
	c.p.Call("load", url, onLoad, onProgress, onError)
	return c
}

// SetPath TODO description.
func (c *CompressedTextureLoader) SetPath(value float64) *CompressedTextureLoader {
	c.p.Call("setPath", value)
	return c
}


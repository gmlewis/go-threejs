package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BinaryTextureLoader represents a binarytextureloader.
type BinaryTextureLoader struct{ p *js.Object }

// BinaryTextureLoader returns a BinaryTextureLoader object.
func (t *Three) BinaryTextureLoader() *BinaryTextureLoader {
	p := t.ctx.Get("BinaryTextureLoader")
	return &BinaryTextureLoader{p: p}
}

// New returns a new BinaryTextureLoader object.
func (t *BinaryTextureLoader) New() *BinaryTextureLoader {
	p := t.p.New()
	return &BinaryTextureLoader{p: p}
}

// Load TODO description.
func (b *BinaryTextureLoader) Load(url, onLoad, onProgress, onError float64) *BinaryTextureLoader {
	b.p.Call("load", url, onLoad, onProgress, onError)
	return b
}


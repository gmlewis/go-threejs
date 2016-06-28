package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// BufferGeometryLoader represents a buffergeometryloader.
type BufferGeometryLoader struct{ p *js.Object }

// BufferGeometryLoader returns a BufferGeometryLoader object.
func (t *Three) BufferGeometryLoader() *BufferGeometryLoader {
	p := t.ctx.Get("BufferGeometryLoader")
	return &BufferGeometryLoader{p: p}
}

// New returns a new BufferGeometryLoader object.
func (t *BufferGeometryLoader) New(manager float64) *BufferGeometryLoader {
	p := t.p.New(manager)
	return &BufferGeometryLoader{p: p}
}

// Load TODO description.
func (b *BufferGeometryLoader) Load(url, onLoad, onProgress, onError float64) *BufferGeometryLoader {
	b.p.Call("load", url, onLoad, onProgress, onError)
	return b
}

// Parse TODO description.
func (b *BufferGeometryLoader) Parse(json float64) *BufferGeometryLoader {
	b.p.Call("parse", json)
	return b
}


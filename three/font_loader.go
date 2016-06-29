package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// FontLoader represents a fontloader.
type FontLoader struct{ p *js.Object }

// FontLoader returns a FontLoader object.
func (t *Three) FontLoader() *FontLoader {
	p := t.ctx.Get("FontLoader")
	return &FontLoader{p: p}
}

// New returns a new FontLoader object.
func (t *FontLoader) New(manager float64) *FontLoader {
	p := t.p.New(manager)
	return &FontLoader{p: p}
}

// Load TODO description.
func (f *FontLoader) Load(url, onLoad, onProgress, onError float64) *FontLoader {
	f.p.Call("load", url, onLoad, onProgress, onError)
	return f
}


package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// FontLoader represents a fontloader.
type FontLoader struct{ p *js.Object }

// FontLoader returns a fontloader object.
func (t *Three) FontLoader() *FontLoader {
	p := t.ctx.Get("FontLoader")
	return &FontLoader{p: p}
}

// NewFontLoader returns a new fontloader object.
func (t *FontLoader) New(manager float64) *FontLoader {
	p := t.p.New(manager)
	return &FontLoader{p: p}
}

// Load TODO description.
func (f *FontLoader) Load(url, onLoad, onProgress, onError float64) *FontLoader {
	f.p.Call("load", url, onLoad, onProgress, onError)
	return f
}


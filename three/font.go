package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Font represents a font.
type Font struct{ p *js.Object }

// Font returns a Font object.
func (t *Three) Font() *Font {
	p := t.ctx.Get("Font")
	return &Font{p: p}
}

// New returns a new Font object.
func (t *Font) New(data float64) *Font {
	p := t.p.New(data)
	return &Font{p: p}
}

// GenerateShapes TODO description.
func (f *Font) GenerateShapes(text, size, divisions float64) *Font {
	f.p.Call("generateShapes", text, size, divisions)
	return f
}


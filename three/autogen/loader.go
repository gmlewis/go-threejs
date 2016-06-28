package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Loader represents a loader.
type Loader struct{ p *js.Object }

// Loader returns a loader object.
func (t *Three) Loader() *Loader {
	p := t.ctx.Get("Loader")
	return &Loader{p: p}
}

// NewLoader returns a new loader object.
func (t *Loader) New() *Loader {
	p := t.p.New()
	return &Loader{p: p}
}

// ExtractUrlBase TODO description.
func (l *Loader) ExtractUrlBase(url float64) *Loader {
	l.p.Call("extractUrlBase", url)
	return l
}

// InitMaterials TODO description.
func (l *Loader) InitMaterials(materials, texturePath, crossOrigin float64) *Loader {
	l.p.Call("initMaterials", materials, texturePath, crossOrigin)
	return l
}

// Add TODO description.
func (l *Loader) Add(regex, loader float64) *Loader {
	l.p.Call("add", regex, loader)
	return l
}

// Get TODO description.
func (l *Loader) Get(file float64) *Loader {
	l.p.Call("get", file)
	return l
}

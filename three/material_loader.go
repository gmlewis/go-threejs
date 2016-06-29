package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MaterialLoader represents a materialloader.
type MaterialLoader struct{ p *js.Object }

// MaterialLoader returns a MaterialLoader object.
func (t *Three) MaterialLoader() *MaterialLoader {
	p := t.ctx.Get("MaterialLoader")
	return &MaterialLoader{p: p}
}

// New returns a new MaterialLoader object.
func (t *MaterialLoader) New(manager float64) *MaterialLoader {
	p := t.p.New(manager)
	return &MaterialLoader{p: p}
}

// Load TODO description.
func (m *MaterialLoader) Load(url, onLoad, onProgress, onError float64) *MaterialLoader {
	m.p.Call("load", url, onLoad, onProgress, onError)
	return m
}

// SetTextures TODO description.
func (m *MaterialLoader) SetTextures(value float64) *MaterialLoader {
	m.p.Call("setTextures", value)
	return m
}

// GetTexture TODO description.
func (m *MaterialLoader) GetTexture(name float64) *MaterialLoader {
	m.p.Call("getTexture", name)
	return m
}

// Parse TODO description.
func (m *MaterialLoader) Parse(json float64) *MaterialLoader {
	m.p.Call("parse", json)
	return m
}


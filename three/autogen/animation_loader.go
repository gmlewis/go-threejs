package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// AnimationLoader represents an animationloader.
type AnimationLoader struct{ p *js.Object }

// AnimationLoader returns an AnimationLoader object.
func (t *Three) AnimationLoader() *AnimationLoader {
	p := t.ctx.Get("AnimationLoader")
	return &AnimationLoader{p: p}
}

// New returns a new AnimationLoader object.
func (t *AnimationLoader) New(manager float64) *AnimationLoader {
	p := t.p.New(manager)
	return &AnimationLoader{p: p}
}

// Load TODO description.
func (a *AnimationLoader) Load(url, onLoad, onProgress, onError float64) *AnimationLoader {
	a.p.Call("load", url, onLoad, onProgress, onError)
	return a
}

// Parse TODO description.
func (a *AnimationLoader) Parse(json, onLoad float64) *AnimationLoader {
	a.p.Call("parse", json, onLoad)
	return a
}


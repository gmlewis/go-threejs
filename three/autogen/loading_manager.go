package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// LoadingManager represents a loadingmanager.
type LoadingManager struct{ p *js.Object }

// LoadingManager returns a loadingmanager object.
func (t *Three) LoadingManager() *LoadingManager {
	p := t.ctx.Get("LoadingManager")
	return &LoadingManager{p: p}
}

// NewLoadingManager returns a new loadingmanager object.
func (t *LoadingManager) New(onLoad, onProgress, onError float64) *LoadingManager {
	p := t.p.New(onLoad, onProgress, onError)
	return &LoadingManager{p: p}
}


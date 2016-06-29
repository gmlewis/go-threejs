package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CubeTextureLoader represents a cubetextureloader.
type CubeTextureLoader struct{ p *js.Object }

// CubeTextureLoader returns a CubeTextureLoader object.
func (t *Three) CubeTextureLoader() *CubeTextureLoader {
	p := t.ctx.Get("CubeTextureLoader")
	return &CubeTextureLoader{p: p}
}

// New returns a new CubeTextureLoader object.
func (t *CubeTextureLoader) New(manager float64) *CubeTextureLoader {
	p := t.p.New(manager)
	return &CubeTextureLoader{p: p}
}

// Load TODO description.
func (c *CubeTextureLoader) Load(urls, onLoad, onProgress, onError float64) *CubeTextureLoader {
	c.p.Call("load", urls, onLoad, onProgress, onError)
	return c
}

// SetCrossOrigin TODO description.
func (c *CubeTextureLoader) SetCrossOrigin(value float64) *CubeTextureLoader {
	c.p.Call("setCrossOrigin", value)
	return c
}

// SetPath TODO description.
func (c *CubeTextureLoader) SetPath(value float64) *CubeTextureLoader {
	c.p.Call("setPath", value)
	return c
}


package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CubeTexture represents a cubetexture.
type CubeTexture struct{ p *js.Object }

// CubeTexture returns a cubetexture object.
func (t *Three) CubeTexture() *CubeTexture {
	p := t.ctx.Get("CubeTexture")
	return &CubeTexture{p: p}
}

// NewCubeTexture returns a new cubetexture object.
func (t *CubeTexture) New(images, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *CubeTexture {
	p := t.p.New(images, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return &CubeTexture{p: p}
}

// Get TODO description.
func (c *CubeTexture) Get() *CubeTexture {
	c.p.Call("get")
	return c
}

// Set TODO description.
func (c *CubeTexture) Set(value float64) *CubeTexture {
	c.p.Call("set", value)
	return c
}

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Texture represents a texture.
type Texture struct{ p *js.Object }

// Texture returns a texture object.
func (t *Three) Texture() *Texture {
	p := t.ctx.Get("Texture")
	return &Texture{p: p}
}

// NewTexture returns a new texture object.
func (t *Texture) New(image, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy float64) *Texture {
	p := t.p.New(image, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)
	return &Texture{p: p}
}

// SetNeedsUpdate sets the needsUpdate-component of the Texture.
func (t *Texture) SetNeedsUpdate(value float64) *Texture {
	t.p.Set("needsUpdate", value)
	return t
}

// Clone TODO description.
func (t *Texture) Clone() *Texture {
	t.p.Call("clone")
	return t
}

// Copy TODO description.
func (t *Texture) Copy(source float64) *Texture {
	t.p.Call("copy", source)
	return t
}

// ToJSON TODO description.
func (t *Texture) ToJSON(meta float64) *Texture {
	t.p.Call("toJSON", meta)
	return t
}

// Dispose TODO description.
func (t *Texture) Dispose() *Texture {
	t.p.Call("dispose")
	return t
}

// TransformUv TODO description.
func (t *Texture) TransformUv(uv float64) *Texture {
	t.p.Call("transformUv", uv)
	return t
}

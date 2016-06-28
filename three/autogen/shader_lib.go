package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ShaderLib represents a shaderlib.
type ShaderLib struct{ p *js.Object }

// ShaderLib returns a ShaderLib object.
func (t *Three) ShaderLib() *ShaderLib {
	p := t.ctx.Get("ShaderLib")
	return &ShaderLib{p: p}
}

// New returns a new ShaderLib object.
func (t *ShaderLib) New() *ShaderLib {
	p := t.p.New()
	return &ShaderLib{p: p}
}


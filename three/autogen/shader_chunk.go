package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// ShaderChunk represents a shaderchunk.
type ShaderChunk struct{ p *js.Object }

// ShaderChunk returns a shaderchunk object.
func (t *Three) ShaderChunk() *ShaderChunk {
	p := t.ctx.Get("ShaderChunk")
	return &ShaderChunk{p: p}
}

// NewShaderChunk returns a new shaderchunk object.
func (t *ShaderChunk) New() *ShaderChunk {
	p := t.p.New()
	return &ShaderChunk{p: p}
}


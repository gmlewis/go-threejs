package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// SceneUtils represents a sceneutils.
type SceneUtils struct{ p *js.Object }

// SceneUtils returns a SceneUtils object.
func (t *Three) SceneUtils() *SceneUtils {
	p := t.ctx.Get("SceneUtils")
	return &SceneUtils{p: p}
}

// New returns a new SceneUtils object.
func (t *SceneUtils) New() *SceneUtils {
	p := t.p.New()
	return &SceneUtils{p: p}
}

// CreateMultiMaterialObject TODO description.
func (s *SceneUtils) CreateMultiMaterialObject(geometry, materials float64) *SceneUtils {
	s.p.Call("createMultiMaterialObject", geometry, materials)
	return s
}

// Detach TODO description.
func (s *SceneUtils) Detach(child, parent, scene float64) *SceneUtils {
	s.p.Call("detach", child, parent, scene)
	return s
}

// Attach TODO description.
func (s *SceneUtils) Attach(child, scene, parent float64) *SceneUtils {
	s.p.Call("attach", child, scene, parent)
	return s
}


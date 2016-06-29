package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// MeshPhongMaterial represents a material for shiny surfaces, evaluated per pixel.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshPhongMaterial
type MeshPhongMaterial struct{ p *js.Object }

// MeshPhongMaterial returns a MeshPhongMaterial object.
func (t *Three) MeshPhongMaterial() *MeshPhongMaterial {
	p := t.ctx.Get("MeshPhongMaterial")
	return &MeshPhongMaterial{p: p}
}

// New returns a new MeshPhongMaterial object.
//
// parameters is an object with one or more of the material's properties defining its appearance:
//     color — geometry color in hexadecimal. Default is 0xffffff.
//     map — Set texture map. Default is null
//     lightMap — Set light map. Default is null.
//     aoMap — Set ao map. Default is null.
//     emissiveMap — Set emissive map. Default is null.
//     specularMap — Set specular map. Default is null.
//     alphaMap — Set alpha map. Default is null.
//     displacementMap — Set displacement map. Default is null.
//     displacementScale — Set displacement scale. Default is 1.
//     displacementBias — Set displacement offset. Default is 0.
//     envMap — Set env map. Default is null.
//     fog — Define whether the material color is affected by global fog settings. Default is true.
//     shading — Define shading type. Default is THREE.SmoothShading.
//     wireframe — render geometry as wireframe. Default is false.
//     wireframeLinewidth — Line thickness. Default is 1.
//     wireframeLinecap — Define appearance of line ends. Default is 'round'.
//     wireframeLinejoin — Define appearance of line joints. Default is 'round'.
//     vertexColors — Define how the vertices gets colored. Default is THREE.NoColors.
//     skinning — Define whether the material uses skinning. Default is false.
//     morphTargets — Define whether the material uses morphTargets. Default is false.
func (t *MeshPhongMaterial) New(parameters map[string]interface{}) *MeshPhongMaterial {
	p := t.p.New(parameters)
	return &MeshPhongMaterial{p: p}
}

// Copy TODO description.
func (m *MeshPhongMaterial) Copy(source *MeshPhongMaterial) *MeshPhongMaterial {
	m.p.Call("copy", source.p)
	return m
}

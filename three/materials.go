package three

import "github.com/gopherjs/gopherjs/js"

// LineBasicMaterial creates a material for drawing wireframe-style geometries.
//     parameters is an object with one or more properties defining the material's appearance:
//     color — Line color in hexadecimal. Default is 0xffffff.
//     linewidth — Line thickness. Default is 1.
//     linecap — Define appearance of line ends. Default is 'round'.
//     linejoin — Define appearance of line joints. Default is 'round'.
//     vertexColors — Define how the vertices gets colored. Default is THREE.NoColors.
//     fog — Define whether the material color is affected by global fog settings. Default is false.
//
// http://threejs.org/docs/index.html#Reference/Materials/LineBasicMaterial
func (t *Three) LineBasicMaterial(parameters map[string]interface{}) *js.Object {
	return t.ctx.Get("LineBasicMaterial").New(parameters)
}

// LineDashedMaterial creates a material for drawing wireframe-style geometries with dashed lines.
//     parameters is an object with one or more properties defining the material's appearance:
//     color — Line color in hexadecimal. Default is 0xffffff.
//     linewidth — Line thickness. Default is 1.
//     scale — The scale of the dashed part of a line. Default is 1.
//     dashSize — The size of the dash. Default is 3.
//     gapSize - The size of the gap. Default is 1.
//     vertexColors — Define how the vertices gets colored. Default is THREE.NoColors.
//     fog — Define whether the material color is affected by global fog settings. Default is false.
//
// http://threejs.org/docs/index.html#Reference/Materials/LineDashedMaterial
func (t *Three) LineDashedMaterial(parameters map[string]interface{}) *js.Object {
	return t.ctx.Get("LineDashedMaterial").New(parameters)
}

// Material creates a generic material.
func (t *Three) Material() *js.Object {
	return t.ctx.Get("Material").New()
}

// MeshBasicMaterial creates a material for drawing geometries in a simple shaded (flat or wireframe) way.
//     parameters is an object with one or more properties defining the material's appearance:
//     color — geometry color in hexadecimal. Default is 0xffffff.
//     map — Set texture map. Default is null
//     aoMap — Set ambient occlusion map. Default is null
//     specularMap — Set specular map. Default is null.
//     alphaMap — Set alpha map. Default is null.
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
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshBasicMaterial
func (t *Three) MeshBasicMaterial(parameters map[string]interface{}) *js.Object {
	return t.ctx.Get("MeshBasicMaterial").New(parameters)
}

// MeshDepthMaterial creates a material for drawing geometry by depth.
// Depth is based off of the camera near and far plane. White is nearest, black is farthest.
//     parameters is an object with one or more properties defining the material's appearance:
//     morphTargets -- Define whether the material uses morphTargets. Default is false.
//     wireframe -- Render geometry as wireframe. Default is false (i.e. render as smooth shaded).
//     wireframeLinewidth -- Controls wireframe thickness. Default is 1.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshDepthMaterial
func (t *Three) MeshDepthMaterial(parameters map[string]interface{}) *js.Object {
	return t.ctx.Get("MeshDepthMaterial").New(parameters)
}

// MeshLambertMaterial creates a material for non-shiny (Lambertian) surfaces, evaluated per vertex.
//     parameters is an object with one or more properties defining the material's appearance:
//     color — Line color in hexadecimal. Default is 0xffffff.
//     map — Sets the texture map. Default is null
//     lightMap — Set light map. Default is null.
//     aoMap — Set ao map. Default is null.
//     emissiveMap — Set emissive map. Default is null.
//     specularMap — Set specular map. Default is null.
//     alphaMap — Set alpha map. Default is null.
//     envMap — Set env map. Default is null.
//     fog — Define whether the material color is affected by global fog settings. Default is false.
//     wireframe — Render geometry as wireframe. Default is false (i.e. render as smooth shaded).
//     wireframeLinewidth — Controls wireframe thickness. Default is 1.
//     wireframeLinecap — Define appearance of line ends. Default is 'round'.
//     wireframeLinejoin — Define appearance of line joints. Default is 'round'.
//     vertexColors — Define how the vertices gets colored. Default is THREE.NoColors.
//     skinning — Define whether the material uses skinning. Default is false.
//     morphTargets — Define whether the material uses morphTargets. Default is false.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshLambertMaterial
func (t *Three) MeshLambertMaterial(parameters map[string]interface{}) *js.Object {
	return t.ctx.Get("MeshLambertMaterial").New(parameters)
}

// MeshNormalMaterial creates a material that maps the normal vectors to RGB colors.
//     parameters is an object with one or more properties defining the material's appearance:
//     wireframe -- Render geometry as wireframe. Default is false (i.e. render as smooth shaded).
//     wireframeLinewidth -- Controls wireframe thickness. Default is 1.
//     morphTargets -- Define whether the material uses morphTargets. Default is false.
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshNormalMaterial
func (t *Three) MeshNormalMaterial(parameters map[string]interface{}) *js.Object {
	return t.ctx.Get("MeshNormalMaterial").New(parameters)
}

// MeshPhongMaterial creates a material for shiny surfaces, evaluated per pixel.
//     parameters is an object with one or more of the material's properties defining its appearance:
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
//
// http://threejs.org/docs/index.html#Reference/Materials/MeshPhongMaterial
func (t *Three) MeshPhongMaterial(parameters map[string]interface{}) *js.Object {
	return t.ctx.Get("MeshPhongMaterial").New(parameters)
}

// MeshStandardMaterial.js

// MultiMaterial.js

// PointsMaterial creates a material used by particle systems.
//     parameters is an object with one or more properties defining the material's appearance:
//     color — Particle color in hexadecimal. Default is 0xffffff.
//     map — a texture.If defined, then a point has the data from texture as colors. Default is null.
//     size — Define size of particles. Default is 1.0.
//     sizeAttenuation — Enable/disable size attenuation with distance.
//     vertexColors — Define whether the material uses vertex colors, or not. Default is false.
//     fog — Define whether the material color is affected by global fog settings. Default is true.
//
// http://threejs.org/docs/index.html#Reference/Materials/PointsMaterial
func (t *Three) PointsMaterial(parameters map[string]interface{}) *js.Object {
	return t.ctx.Get("PointsMaterial").New(parameters)
}

// RawShaderMaterial.js

// ShaderMaterial.js

// SpriteMaterial creates a material for a sprite.
//     parameters is an object defining the the default properties:
//     color - color of the sprite
//     map - the texture map
//     rotation - the rotation of the sprite
//     fog - whether or not to use the scene fog
//
// http://threejs.org/docs/index.html#Reference/Materials/SpriteMaterial
func (t *Three) SpriteMaterial(parameters map[string]interface{}) *js.Object {
	return t.ctx.Get("SpriteMaterial").New(parameters)
}

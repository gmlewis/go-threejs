package three

import "github.com/gopherjs/gopherjs/js"

// Bone creates a bone which is part of a skeleton.
//
// http://threejs.org/docs/index.html#Reference/Objects/Bone
func (t *Three) Bone() *js.Object { return t.ctx.Get("Bone").New() }

// Group creates a group.
func (t *Three) Group() *js.Object { return t.ctx.Get("Group").New() }

// LensFlare creates a simulated lens flare that tracks a light.
//     texture -- THREE.Texture (optional)
//     size -- size in pixels (-1 = use texture.width)
//     distance -- (0-1) from light source (0 = at light source)
//     blending -- Blending Mode - Defaults to THREE.NormalBlending
//     color -- The color of the lens flare
//
// http://threejs.org/docs/index.html#Reference/Objects/LensFlare
func (t *Three) LensFlare(texture *js.Object, size, distance float64, blending, color int) *js.Object {
	return t.ctx.Get("LensFlare").New(texture, size, distance, blending, color)
}

// Line creates a continuous line.
//     geometry — Vertices representing the line segment(s).
//     material — Material for the line. Default is LineBasicMaterial.
//
// http://threejs.org/docs/index.html#Reference/Objects/Line
func (t *Three) Line(geometry []*js.Object, material *js.Object) *js.Object {
	return t.ctx.Get("Line").New(geometry, material)
}

// LineSegments creates a series of lines.
//     geometry — Vertices representing the line segment(s).
//     material — Material for the line. Default is LineBasicMaterial.
//
// http://threejs.org/docs/index.html#Reference/Objects/LineSegments
func (t *Three) LineSegments(geometry []*js.Object, material *js.Object) *js.Object {
	return t.ctx.Get("LineSegments").New(geometry, material)
}

// LOD shows meshes with more or less geometry based on distance.
//
// http://threejs.org/docs/index.html#Reference/Objects/LOD
func (t *Three) LOD() *js.Object {
	return t.ctx.Get("LOD").New()
}

// Mesh is the base class for mesh objects.
//     geometry — an instance of geometry.
//     material — an instance of material (optional).
//
// http://threejs.org/docs/index.html#Reference/Objects/Mesh
func (t *Three) Mesh(geometry, material *js.Object) *js.Object {
	return t.ctx.Get("Mesh").New(geometry, material)
}

// Points is a class for displaying particles in the form of variable size points.
//     geometry — an instance of geometry.
//     material — an instance of material (optional).
//
// http://threejs.org/docs/index.html#Reference/Objects/Points
func (t *Three) Points(geometry, material *js.Object) *js.Object {
	return t.ctx.Get("Points").New(geometry, material)
}

// Skeleton uses an array of bones to create a skeleton that can be used by a SkinnedMesh.
//     bones — The array of bones
//     boneInverses — (optional) An array of Matrix4s
//     useVertexTexture — (optional) Whether or not to use a vertex texture in the shader.
//
// http://threejs.org/docs/index.html#Reference/Objects/Skeleton
func (t *Three) Skeleton(bones, boneInverses []*js.Object, useVertexTexture bool) *js.Object {
	return t.ctx.Get("Skeleton").New(bones, boneInverses, useVertexTexture)
}

// SkinnedMesh creates a mesh that has a Skeleton with bones that can then be used
// to animate the vertices of the geometry.
//     geometry — An instance of Geometry. Geometry.skinIndices and Geometry.skinWeights should be set.
//     material — An instance of Material (optional).
//     useVertexTexture -- Defines whether a vertex texture can be used (optional).
//
// http://threejs.org/docs/index.html#Reference/Objects/SkinnedMesh
func (t *Three) SkinnedMesh(geometry, material *js.Object, useVertexTexture bool) *js.Object {
	return t.ctx.Get("SkinnedMesh").New(geometry, material, useVertexTexture)
}

// Sprite creates a plane in an 3d scene which faces always towards the camera.
//     material — An instance of Material (optional).
//
// http://threejs.org/docs/index.html#Reference/Objects/Sprite
func (t *Three) Sprite(material *js.Object) *js.Object {
	return t.ctx.Get("Sprite").New(material)
}

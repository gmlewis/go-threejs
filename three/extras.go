package three

import "github.com/gopherjs/gopherjs/js"

// core/Curve.js

// core/CurvePath.js

// core/Font.js

// core/Path.js

// core/Shape.js

// curves/ArcCurve.js

// curves/CatmullRomCurve3.js

// curves/ClosedSplineCurve3.js

// curves/CubicBezierCurve3.js

// curves/CubicBezierCurve.js

// curves/EllipseCurve.js

// curves/LineCurve3.js

// curves/LineCurve.js

// curves/QuadraticBezierCurve3.js

// curves/QuadraticBezierCurve.js

// curves/SplineCurve3.js

// curves/SplineCurve.js

// CurveUtils.js

type BoxGeometryOpts struct {
	widthSegments  int // Number of segmented faces along the width of the sides. Default is 1.
	heightSegments int // Number of segmented faces along the height of the sides. Default is 1.
	depthSegments  int // Number of segmented faces along the depth of the sides. Default is 1.
}

// BoxGeometry creates a quadrilateral geometry primitive.
//     width — Width of the sides on the X axis.
//     height — Height of the sides on the Y axis.
//     depth — Depth of the sides on the Z axis.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/BoxGeometry
func (t *Three) BoxGeometry(width, height, depth float64, opts *BoxGeometryOpts) *js.Object {
	if opts != nil {
		return t.ctx.Get("BoxGeometry").New(width, height, depth, opts.widthSegments, opts.heightSegments, opts.depthSegments)
	}
	return t.ctx.Get("BoxGeometry").New(width, height, depth)
}

// CircleGeometry creates a simple circular shape.
//     radius — Radius of the circle, default = 50.
//     segments — Number of segments (triangles), minimum = 3, default = 8.
//     thetaStart — Start angle for first segment, default = 0 (three o'clock position).
//     thetaLength — The central angle, often called theta, of the circular sector.
//         The default is 2*Pi, which makes for a complete circle.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/CircleGeometry
func (t *Three) CircleGeometry(radius float64, segments int, thetaStart, thetaLength float64) *js.Object {
	return t.ctx.Get("CircleGeometry").New(radius, segments, thetaStart, thetaLength)
}

// CylinderGeometry creates a cylinder geometry primitive.
//     radiusTop — Radius of the cylinder at the top. Default is 20.
//     radiusBottom — Radius of the cylinder at the bottom. Default is 20.
//     height — Height of the cylinder. Default is 100.
//     radiusSegments — Number of segmented faces around the circumference of the cylinder. Default is 8
//     heightSegments — Number of rows of faces along the height of the cylinder. Default is 1.
//     openEnded — A Boolean indicating whether the ends of the cylinder are open or capped.
//         Default is false, meaning capped.
//     thetaStart — Start angle for first segment, default = 0 (three o'clock position).
//     thetaLength — The central angle, often called theta, of the circular sector.
//         The default is 2*Pi, which makes for a complete cylinder.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/CylinderGeometry
func (t *Three) CylinderGeometry(radiusTop, radiusBottom, height float64, radiusSegments, heightSegments int, openEnded bool, thetaStart, thetaLength float64) *js.Object {
	return t.ctx.Get("CylinderGeometry").New(radiusTop, radiusBottom, height, radiusSegments, heightSegments, openEnded, thetaStart, thetaLength)
}

// DodecahedronGeometry creates a dodecagedron.
//     radius — Radius of the dodecahedron. Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds vertices making it no longer a dodecahedron.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/DodecahedronGeometry
func (t *Three) DodecahedronGeometry(radius, detail float64) *js.Object {
	return t.ctx.Get("DodecahedronGeometry").New(radius, detail)
}

// geometries/EdgesGeometry.js

// geometries/ExtrudeGeometry.js

// IcosahedronGeometry creates an icosahedron.
//     radius — Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds more vertices making it no longer an icosahedron. When detail is greater than 1, it's effectively a sphere.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/IcosahedronGeometry
func (t *Three) IcosahedronGeometry(radius, detail float64) *js.Object {
	return t.ctx.Get("IcosahedronGeometry").New(radius, detail)
}

// LatheGeometry creates a mesh with axial symmetry.
//     points — Slice of Vector2s.
//     segments — the number of circumference segments to generate. Default is 12.
//     phiStart — the starting angle in radians. Default is 0.
//     phiLength — the radian (0 to 2PI) range of the lathed section 2PI is a closed lathe, less than 2PI is a portion. Default is 2*PI
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/LatheGeometry
func (t *Three) LatheGeometry(points []*js.Object, segments int, phiStart, phiLength float64) *js.Object {
	return t.ctx.Get("LatheGeometry").New(points, segments, phiStart, phiLength)
}

// OctahedronGeometry creates an octahedron.
//     radius — Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds more vertices making it no longer an icosahedron. When detail is greater than 1, it's effectively a sphere.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/OctahedronGeometry
func (t *Three) OctahedronGeometry(radius, detail float64) *js.Object {
	return t.ctx.Get("OctahedronGeometry").New(radius, detail)
}

// ParametricGeometry creates geometry representing a parametric surface.
//     f — A function that takes in a u and v value each between 0 and 1 and returns a Vector3
//     slices — The count of slices to use for the parametric function
//     stacks — The count of stacks to use for the parametric function
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/ParametricGeometry
func (t *Three) ParametricGeometry(f func(u, v float64) *js.Object, slices, stacks int) *js.Object {
	return t.ctx.Get("ParametricGeometry").New(f, slices, stacks)
}

// PlaneGeometry creates a plane.
//     width — Width along the X axis.
//     height — Height along the Y axis.
//     widthSegments — Optional. Default is 1.
//     heightSegments — Optional. Default is 1.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/PlaneGeometry
func (t *Three) PlaneGeometry(width, height float64, widthSegments, heightSegments int) *js.Object {
	return t.ctx.Get("PlaneGeometry").New(width, height, widthSegments, heightSegments)
}

// PolyhedronGeometry creates a solid with flat faces.
//     vertices — Array of points of the form [1,1,1, -1,-1,-1, ... ]
//     faces — Array of indices that make up the faces of the form [0,1,2, 2,3,0, ... ]
//     radius — Float - The radius of the final shape
//     detail — Integer - How many levels to subdivide the geometry. The more detail, the smoother the shape.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/PolyhedronGeometry
func (t *Three) PolyhedronGeometry(vertices []float64, faces []int, radius float64, detail int) *js.Object {
	return t.ctx.Get("PolyhedronGeometry").New(vertices, faces, radius, detail)
}

// RingGeometry creates two-dimensional ring geometry.
//     innerRadius — Default is 0, but it doesn't work right when innerRadius is set to 0.
//     outerRadius — Default is 50.
//     thetaSegments — Number of segments. A higher number means the ring will be more round. Minimum is 3. Default is 8.
//     phiSegments — Minimum is 1. Default is 8.
//     thetaStart — Starting angle. Default is 0.
//     thetaLength — Central angle. Default is Math.PI * 2.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/RingGeometry
func (t *Three) RingGeometry(innerRadius, outerRadius float64, thetaSegments, phiSegments int, thetaStart, thetaLength float64) *js.Object {
	return t.ctx.Get("RingGeometry").New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)
}

// geometries/ShapeGeometry.js

// SphereGeometry creates a sphere.
//     radius — sphere radius. Default is 50.
//     widthSegments — number of horizontal segments. Minimum value is 3, and the default is 8.
//     heightSegments — number of vertical segments. Minimum value is 2, and the default is 6.
//     phiStart — specify horizontal starting angle. Default is 0.
//     phiLength — specify horizontal sweep angle size. Default is Math.PI * 2.
//     thetaStart — specify vertical starting angle. Default is 0.
//     thetaLength — specify vertical sweep angle size. Default is Math.PI.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/SphereGeometry
func (t *Three) SphereGeometry(radius float64, widthSegments, heightSegments int, phiStart, phiLength, thetaStart, thetaLength float64) *js.Object {
	return t.ctx.Get("SphereGeometry").New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength)
}

// TetrahedronGeometry creates a tetrahedron.
//     radius — Radius of the tetrahedron. Default is 1.
//     detail — Default is 0. Setting this to a value greater than 0 adds vertices making it no longer a tetrahedron.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TetrahedronGeometry
func (t *Three) TetrahedronGeometry(radius, detail float64) *js.Object {
	return t.ctx.Get("TetrahedronGeometry").New(radius, detail)
}

// TextGeometry creates a 3D object of text as a single object.
//     text — The text to be shown.
//     parameters — Object that can contains the following parameters.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TextGeometry
func (t *Three) TextGeometry(text string, parameters map[string]interface{}) *js.Object {
	return t.ctx.Get("TextGeometry").New(text, parameters)
}

// TorusGeometry creates a torus.
//     radius — Default is 100.
//     tube — Diameter of the tube. Default is 40.
//     radialSegments — Default is 8
//     tubularSegments — Default is 6.
//     arc — Central angle. Default is Math.PI * 2.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TorusGeometry
func (t *Three) TorusGeometry(radius, tube float64, radialSegments, tubularSegments int, arc float64) *js.Object {
	return t.ctx.Get("TorusGeometry").New(radius, tube, radialSegments, tubularSegments, arc)
}

// TorusKnotGeometry creates a torus knot.
//     radius — Default is 100.
//     tube — Default is 40.
//     radialSegments — Default is 64.
//     tubularSegments — Default is 8.
//     p — Default is 2.
//     q — Default is 3.
//     heightScale — Default is 1.
//
// http://threejs.org/docs/index.html#Reference/Extras.Geometries/TorusKnotGeometry
func (t *Three) TorusKnotGeometry(radius, tube float64, radialSegments, tubularSegments int, p, q, heightScale float64) *js.Object {
	return t.ctx.Get("TorusKnotGeometry").New(radius, tube, radialSegments, tubularSegments, p, q, heightScale)
}

// geometries/TubeGeometry.js

// geometries/WireframeGeometry.js

// helpers/ArrowHelper.js

// helpers/AxisHelper.js

// helpers/BoundingBoxHelper.js

// helpers/BoxHelper.js

// helpers/CameraHelper.js

// helpers/DirectionalLightHelper.js

// helpers/EdgesHelper.js

// helpers/FaceNormalsHelper.js

// helpers/GridHelper.js

// helpers/HemisphereLightHelper.js

// helpers/PointLightHelper.js

// helpers/SkeletonHelper.js

// helpers/SpotLightHelper.js

// helpers/VertexNormalsHelper.js

// helpers/WireframeHelper.js

// objects/ImmediateRenderObject.js

// objects/MorphBlendMesh.js

// SceneUtils.js

// ShapeUtils.js

package three

import "github.com/gopherjs/gopherjs/js"

// CubeCamera creates 6 cameras that render to a WebGLRenderTargetCube.
//
// http://threejs.org/docs/index.html#Reference/Cameras/CubeCamera
func (t *Three) CubeCamera(near, far, cubeResolution float64) *js.Object {
	return t.ctx.Get("CubeCamera").New(near, far, cubeResolution)
}

// OrthographicCamera creates a camera with orthographic projection.
//
// http://threejs.org/docs/index.html#Reference/Cameras/OrthographicCamera
func (t *Three) OrthographicCamera(left, right, top, bottom, near, far float64) *js.Object {
	return t.ctx.Get("OrthographicCamera").New(left, right, top, bottom, near, far)
}

// PerspectiveCamera creates a camera with perspective projection.
//
// http://threejs.org/docs/index.html#Reference/Cameras/PerspectiveCamera
func (t *Three) PerspectiveCamera(fov, aspect, near, far float64) *js.Object {
	return t.ctx.Get("PerspectiveCamera").New(fov, aspect, near, far)
}

// StereoCamera creates a stereo projection camera with left and right eyes.
func (t *Three) StereoCamera() *js.Object { return t.ctx.Get("StereoCamera").New() }

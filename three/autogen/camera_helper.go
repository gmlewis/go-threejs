package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// CameraHelper represents a camerahelper.
type CameraHelper struct{ p *js.Object }

// CameraHelper returns a camerahelper object.
func (t *Three) CameraHelper() *CameraHelper {
	p := t.ctx.Get("CameraHelper")
	return &CameraHelper{p: p}
}

// NewCameraHelper returns a new camerahelper object.
func (t *CameraHelper) New(camera float64) *CameraHelper {
	p := t.p.New(camera)
	return &CameraHelper{p: p}
}


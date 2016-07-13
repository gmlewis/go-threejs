// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Texture represents a texture.
type Texture struct{ p *js.Object }

// JSObject returns the underlying *js.Object.
func (t *Texture) JSObject() *js.Object { return t.p }

// Texture returns a Texture JavaScript class.
func (t *Three) Texture() *Texture {
	p := t.ctx.Get("Texture")
	return TextureFromJSObject(p)
}

// TextureFromJSObject returns a wrapped Texture JavaScript class.
func TextureFromJSObject(p *js.Object) *Texture {
	return &Texture{p: p}
}

// TextureOpts representions optional arguments for creating a Texture.
//
//     mapping - How the image is applied to the object. An object
//         type of THREE.UVMapping is the default, where the U,V
//         coordinates are used to apply the map, and a single
//         texture is expected. The other types are
//         THREE.CubeReflectionMapping, for cube maps used as a
//         reflection map; THREE.CubeRefractionMapping, refraction
//         mapping; and THREE.SphericalReflectionMapping, a spherical
//         reflection map projection.
//
//     wrapS - The default is THREE.ClampToEdgeWrapping, where the
//         edge is clamped to the outer edge texels. The other two
//         choices are THREE.RepeatWrapping and
//         THREE.MirroredRepeatWrapping.
//
//     wrapT - The default is THREE.ClampToEdgeWrapping, where the
//         edge is clamped to the outer edge texels. The other two
//         choices are THREE.RepeatWrapping and
//         THREE.MirroredRepeatWrapping. NOTE: tiling of images in
//         textures only functions if image dimensions are powers of
//         two (2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, ...)
//         in terms of pixels. Individual dimensions need not be
//         equal, but each must be a power of two. This is a
//         limitation of WebGL, not Three.js.
//
//     magFilter - How the texture is sampled when a texel covers
//         more than one pixel. The default is THREE.LinearFilter,
//         which takes the four closest texels and bilinearly
//         interpolates among them. The other option is
//         THREE.NearestFilter, which uses the value of the closest
//         texel.
//
//     minFilter - How the texture is sampled when a texel covers
//         less than one pixel. The default is
//         THREE.LinearMipMapLinearFilter, which uses mipmapping and
//         a trilinear filter. Other choices are THREE.NearestFilter,
//         THREE.NearestMipMapNearestFilter,
//         THREE.NearestMipMapLinearFilter, THREE.LinearFilter, and
//         THREE.LinearMipMapNearestFilter. These vary whether the
//         nearest texel or nearest four texels are retrieved on the
//         nearest mipmap or nearest two mipmaps. Interpolation
//         occurs among the samples retrieved.
//
//     format - The default is THREE.RGBAFormat for the
//         texture. Other formats are: THREE.AlphaFormat,
//         THREE.RGBFormat, THREE.LuminanceFormat, and
//         THREE.LuminanceAlphaFormat. There are also compressed
//         texture formats, if the S3TC extension is supported:
//         THREE.RGB_S3TC_DXT1_Format, THREE.RGBA_S3TC_DXT1_Format,
//         THREE.RGBA_S3TC_DXT3_Format, and
//         THREE.RGBA_S3TC_DXT5_Format.
//
//     type - The default is THREE.UnsignedByteType. Other valid
//         types (as WebGL allows) are THREE.ByteType,
//         THREE.ShortType, THREE.UnsignedShortType, THREE.IntType,
//         THREE.UnsignedIntType, THREE.FloatType,
//         THREE.UnsignedShort4444Type, THREE.UnsignedShort5551Type,
//         and THREE.UnsignedShort565Type.
//
//     anisotropy - The number of samples taken along the axis
//         through the pixel that has the highest density of
//         texels. By default, this value is 1. A higher value gives
//         a less blurry result than a basic mipmap, at the cost of
//         more texture samples being used. Use
//         renderer.getMaxAnisotropy() to find the maximum valid
//         anisotropy value for the GPU; this value is usually a
//         power of 2.
type TextureOpts struct {
	mapping    int
	wrapS      int
	wrapT      int
	magFilter  int
	minFilter  int
	format     int
	typ        int
	anisotropy int
}

// NewTexture returns a new Texture object.
func (t *Three) NewTexture(image *js.Object, opts *TextureOpts) *Texture {
	p := t.ctx.Get("Texture")
	if opts != nil {
		p = p.New(image, opts.mapping, opts.wrapS, opts.wrapT, opts.magFilter, opts.minFilter, opts.format, opts.typ, opts.anisotropy)
	} else {
		p = p.New(image)
	}
	return TextureFromJSObject(p)
}

// SetNeedsUpdate sets the needsUpdate-component of the Texture.
func (t *Texture) SetNeedsUpdate(value bool) *Texture {
	t.p.Set("needsUpdate", value)
	return t
}

// Clone TODO description.
func (t *Texture) Clone() *Texture {
	t.p.Call("clone")
	return t
}

// Copy TODO description.
func (t *Texture) Copy(source *Texture) *Texture {
	t.p.Call("copy", source.p)
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

// SetWrapS sets the wrapS-component of the Texture.
func (t *Texture) SetWrapS(value int) *Texture {
	t.p.Set("wrapS", value)
	return t
}

// SetWrapT sets the wrapT-component of the Texture.
func (t *Texture) SetWrapT(value int) *Texture {
	t.p.Set("wrapT", value)
	return t
}

// SetAnisotropy sets the anisotropy-component of the Texture.
func (t *Texture) SetAnisotropy(value int) *Texture {
	t.p.Set("anisotropy", value)
	return t
}

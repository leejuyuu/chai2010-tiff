// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Auto Generated By 'go generate', DONOT EDIT!!!

package tiff

import (
	"image"
	"image/color"
	"reflect"
)

type imageRGB192f struct {
	M struct {
		Pix    []uint8
		Stride int
		Rect   image.Rectangle
	}
}

// newImageRGB192f returns a new imageRGB192f with the given bounds.
func newImageRGB192f(r image.Rectangle) *imageRGB192f {
	return new(imageRGB192f).Init(make([]uint8, 24*r.Dx()*r.Dy()), 24*r.Dx(), r)
}

func (p *imageRGB192f) Init(pix []uint8, stride int, rect image.Rectangle) *imageRGB192f {
	*p = imageRGB192f{
		M: struct {
			Pix    []uint8
			Stride int
			Rect   image.Rectangle
		}{
			Pix:    pix,
			Stride: stride,
			Rect:   rect,
		},
	}
	return p
}

func (p *imageRGB192f) Pix() []byte           { return p.M.Pix }
func (p *imageRGB192f) Stride() int           { return p.M.Stride }
func (p *imageRGB192f) Rect() image.Rectangle { return p.M.Rect }
func (p *imageRGB192f) Channels() int         { return 3 }
func (p *imageRGB192f) Depth() reflect.Kind   { return reflect.Float64 }

func (p *imageRGB192f) ColorModel() color.Model { return colorRGB192fModel }

func (p *imageRGB192f) Bounds() image.Rectangle { return p.M.Rect }

func (p *imageRGB192f) At(x, y int) color.Color {
	return p.RGB192fAt(x, y)
}

func (p *imageRGB192f) RGB192fAt(x, y int) colorRGB192f {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return colorRGB192f{}
	}
	i := p.PixOffset(x, y)
	return pRGB192fAt(p.M.Pix[i:])
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *imageRGB192f) PixOffset(x, y int) int {
	return (y-p.M.Rect.Min.Y)*p.M.Stride + (x-p.M.Rect.Min.X)*24
}

func (p *imageRGB192f) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := colorRGB192fModel.Convert(c).(colorRGB192f)
	pSetRGB192f(p.M.Pix[i:], c1)
	return
}

func (p *imageRGB192f) SetRGB192f(x, y int, c colorRGB192f) {
	if !(image.Point{x, y}.In(p.M.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	pSetRGB192f(p.M.Pix[i:], c)
	return
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *imageRGB192f) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.M.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &imageRGB192f{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return new(imageRGB192f).Init(
		p.M.Pix[i:],
		p.M.Stride,
		r,
	)
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *imageRGB192f) Opaque() bool {
	return true
}
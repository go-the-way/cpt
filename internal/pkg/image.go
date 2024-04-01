// Copyright 2023 cpt Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pkg

import (
	"image"
	"image/color"
	"math/rand"
)

func CutOut(bgImg, bcImg, newBcImg image.Image, x int) {
	bcWidth := bcImg.Bounds().Dx()
	bcHeight := bcImg.Bounds().Dy()
	var values [9]color.RGBA64
	for i := 0; i < bcWidth; i++ {
		for j := 0; j < bcHeight; j++ {
			pixel := getRGBA(bcImg, i, j)
			if pixel.A > 0 {
				setRGBA(newBcImg, i, j, getRGBA(bgImg, x+i, j))
				readNeighborPixel(bgImg, x+i, j, &values)
				setRGBA(bgImg, x+i, j, gaussianBlur(&values))
			}
			if i == (bcWidth-1) || j == (bcHeight-1) {
				continue
			}
			rightPixel := getRGBA(bcImg, i+1, j)
			bottomPixel := getRGBA(bcImg, i, j+1)
			if (pixel.A > 0 && rightPixel.A == 0) ||
				(pixel.A == 0 && rightPixel.A > 0) ||
				(pixel.A > 0 && bottomPixel.A == 0) ||
				(pixel.A == 0 && bottomPixel.A > 0) {
				white := color.White
				setRGBA(newBcImg, i, j, white)
				setRGBA(bgImg, x+i, j, white)
			}
		}
	}
}

func Interfere(bgImg, itImg image.Image, x int) {
	var values [9]color.RGBA64
	itWidth := itImg.Bounds().Dx()
	itHeight := itImg.Bounds().Dy()
	for i := 0; i < itWidth; i++ {
		for j := 0; j < itHeight; j++ {
			pixel := getRGBA(itImg, i, j)
			if pixel.A > 0 {
				readNeighborPixel(bgImg, x+i, j, &values)
				setRGBA(bgImg, x+i, j, gaussianBlur(&values))
			}
			if i == (itWidth-1) || j == (itHeight-1) {
				continue
			}
			rightPixel := getRGBA(itImg, i+1, j)
			bottomPixel := getRGBA(itImg, i, j+1)
			if (pixel.A > 0 && rightPixel.A == 0) ||
				(pixel.A == 0 && rightPixel.A > 0) ||
				(pixel.A > 0 && bottomPixel.A == 0) ||
				(pixel.A == 0 && bottomPixel.A > 0) {
				white := color.White
				setRGBA(bgImg, x+i, j, white)
			}
		}
	}
}

func RandPoint(bgWidth, bgHeight, bcWidth, _ int) *image.Point {
	wDiff := bgWidth - bcWidth
	hDiff := bgHeight - bcWidth
	var x, y int
	if wDiff <= 0 {
		x = 5
	} else {
		x = rand.Intn(wDiff-100) + 100
	}
	if hDiff <= 0 {
		y = 5
	} else {
		y = rand.Intn(hDiff) + 5
	}
	return &image.Point{X: x, Y: y}
}

func readNeighborPixel(img image.Image, x, y int, pixels *[9]color.RGBA64) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	xStart := x - 1
	yStart := y - 1
	current := 0
	for i := xStart; i < 3+xStart; i++ {
		for j := yStart; j < 3+yStart; j++ {
			tx := i
			if tx < 0 {
				tx = -tx
			} else if tx >= width {
				tx = x
			}
			ty := j
			if ty < 0 {
				ty = -ty
			} else if ty >= height {
				ty = y
			}
			pixels[current] = getRGBA(img, tx, ty)
			current++
		}
	}
}

func gaussianBlur(values *[9]color.RGBA64) color.RGBA64 {
	var r uint32
	var g uint32
	var b uint32
	var a uint32
	for i := 0; i < len(values); i++ {
		if i == 4 {
			continue
		}
		x := values[i]
		r += uint32(x.R)
		g += uint32(x.G)
		b += uint32(x.B)
		a += uint32(x.A)
	}
	return color.RGBA64{R: uint16(r / 8), G: uint16(g / 8), B: uint16(b / 8), A: uint16(a / 8)}
}

func getRGBA(img image.Image, x, y int) color.RGBA64 {
	r, g, b, a := img.At(x, y).RGBA()
	return color.RGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: uint16(a)}
}

func setRGBA(img image.Image, x, y int, c color.Color) {
	switch img.(type) {
	case *image.RGBA:
		img.(*image.RGBA).Set(x, y, c)
	case *image.NRGBA:
		img.(*image.NRGBA).Set(x, y, c)
	}
}

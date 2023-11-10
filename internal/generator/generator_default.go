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

package generator

import (
	"github.com/rwscode/cpt/internal/conf"
	"github.com/rwscode/cpt/internal/pkg"
	"github.com/rwscode/cpt/internal/resloader"
	"github.com/rwscode/cpt/internal/verifier"
	"image"
	"math/rand"
)

type defaultGenerator struct {
	resloader.ResLoader
	verifier.Verifier
}

func DefaultGenerator(resLoader resloader.ResLoader, ver verifier.Verifier) *defaultGenerator {
	return &defaultGenerator{resLoader, ver}
}

func (d *defaultGenerator) Generate() (imageToken ImageToken, err error) {
	imageToken, err = d.generateImage()
	if err != nil {
		return
	}
	d.Store(imageToken.Token, imageToken.X)
	return
}

func (d *defaultGenerator) generateImage() (imageToken ImageToken, err error) {
	var (
		bgImg, bcImg, itImg image.Image
	)
	token := d.Token(conf.GetTokenLength())
	if bgImg, err = d.BGImg(); err != nil {
		return
	}
	if bcImg, itImg, err = d.BCImg(); err != nil {
		return
	}
	bgWidth := bgImg.Bounds().Dx()
	bgHeight := bgImg.Bounds().Dy()
	bcWidth := bcImg.Bounds().Dx()
	bcHeight := bcImg.Bounds().Dy()
	point := pkg.RandPoint(bgWidth, bgHeight, bcWidth, bcHeight)
	// 抠图
	x := point.X
	newBcImg := image.NewNRGBA(image.Rect(0, 0, bgWidth, bgHeight))
	pkg.CutOut(bgImg, bcImg, newBcImg, x)
	// 插入干扰图片
	position := 0
	if bgWidth-x-5 > bcWidth*2 {
		// 在原扣图右边插入干扰图
		position = rand.Intn((bgWidth-bcWidth)-(x+bcWidth+5)) + (x + bcWidth + 5)
	} else {
		// 在原扣图左边插入干扰图
		position = rand.Intn((x-bcWidth-5)-100) + 100
	}
	// 干扰图
	pkg.Interfere(bgImg, itImg, position)
	var (
		bgImageBase64, bcImageBase64 string
	)
	if bgImageBase64, err = pkg.Image2Base64Img(bgImg); err != nil {
		return
	}
	if bcImageBase64, err = pkg.Image2Base64Img(newBcImg); err != nil {
		return
	}
	return ImageToken{BgImage: bgImg, BcImage: bcImg, X: x, BgImageBase64: bgImageBase64, BcImageBase64: bcImageBase64, Token: token}, nil
}

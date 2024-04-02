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

package resloader

import (
	"bytes"
	"embed"
	"errors"
	"image"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	_ "image/png"
)

var (
	//go:embed images/backgrounds/*.png
	bundleBgDirFs embed.FS
	//go:embed images/blocks/*.png
	bundleBcDirFs embed.FS
	logger        = log.New(os.Stdout, "[cpt_resloader] ", log.LstdFlags|log.Lshortfile)
	once          = &sync.Once{}
)

const (
	bundleBgDirFsName = "images/backgrounds"
	bundleBcDirFsName = "images/blocks"
)

type fsResLoader struct {
	bgFs, bcFs         *embed.FS
	bgFsName, bcFsName string
	bgs, bcs           [][]byte
}

func FsResLoader(bgFs, bcFs *embed.FS, bgFsName, bcFsName string) *fsResLoader {
	fr := &fsResLoader{bgFs: bgFs, bcFs: bcFs, bgFsName: bgFsName, bcFsName: bcFsName}
	once.Do(fr.init)
	return fr
}

func FsResLoaderDefault() *fsResLoader {
	return FsResLoader(&bundleBgDirFs, &bundleBcDirFs, bundleBgDirFsName, bundleBcDirFsName)
}

func (f *fsResLoader) init() {
	initFn := func(fs *embed.FS, name string) [][]byte {
		list := make([][]byte, 0)
		files, err := fs.ReadDir(name)
		if err != nil {
			logger.Panicln(err)
		}
		for _, file := range files {
			if !file.IsDir() {
				fileName := name + "/" + file.Name()
				buf, bErr := fs.ReadFile(fileName)
				if bErr != nil {
					logger.Panicln(bErr)
				}
				if buf != nil {
					list = append(list, buf)
					logger.Println("attached: " + fileName)
				}
			}
		}
		return list
	}
	f.bgs = initFn(f.bgFs, f.bgFsName)
	f.bcs = initFn(f.bcFs, f.bcFsName)
}

func (f *fsResLoader) BGImg() (bgImg image.Image, err error) {
	length := len(f.bgs)
	if f.bgs == nil || length <= 0 {
		err = errors.New("the background image list was empty, please set res fs first")
		return
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(length)
	bgImg, _, _ = image.Decode(bytes.NewReader(f.bgs[r]))
	return
}

func (f *fsResLoader) BCImg() (bcImg, itImg image.Image, err error) {
	length := len(f.bcs)
	if f.bgs == nil || length <= 0 {
		err = errors.New("the block image list was empty, please set res fs first")
		return
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(length)
	rr := (r + length/2) % (length - 1)
	bcImg, _, _ = image.Decode(bytes.NewReader(f.bcs[r]))
	itImg, _, _ = image.Decode(bytes.NewReader(f.bcs[rr]))
	return
}

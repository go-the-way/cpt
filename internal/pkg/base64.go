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
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
)

func Image2Base64(img image.Image) (str string, err error) {
	var buf bytes.Buffer
	if err = png.Encode(&buf, img); err != nil {
		return
	}
	str = base64.StdEncoding.EncodeToString(buf.Bytes())
	return
}

func Image2Base64Img(img image.Image) (str string, err error) {
	image2Base64, err := Image2Base64(img)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("data:image/png;base64,%s", image2Base64), nil
}

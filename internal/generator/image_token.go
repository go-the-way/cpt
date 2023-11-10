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
	"encoding/json"
	"image"
)

type ImageToken struct {
	BgImage image.Image `json:"-"`
	BcImage image.Image `json:"-"`
	X       int         `json:"-"`

	BgImageBase64 string `json:"bg_image_base_64"`
	BcImageBase64 string `json:"bc_image_base_64"`
	Token         string `json:"token"`
}

func (t *ImageToken) JSON() string {
	if t == nil {
		return `{"err":"missing image token"}`
	}
	buf, err := json.Marshal(t)
	if err != nil {
		return `{"err":"` + err.Error() + `"}`
	}
	return string(buf)
}

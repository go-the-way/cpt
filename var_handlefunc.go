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

package cpt

import (
	"github.com/rwscode/cpt/middleware"
	"net/http"
	"strconv"
)

func generateHandlerFunc(middlewares ...http.HandlerFunc) http.HandlerFunc {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		imageToken, err := defGtr.Generate()
		if err != nil {
			writeJSON(w, `{"err":"`+err.Error()+`"}`)
			return
		}
		writeJSON(w, imageToken.JSON())
		return
	})
	middlewares = append(middlewares, middleware.GzipHandler(h).ServeHTTP)
	return walkHandlerFunc(middlewares...)
}

func verifyHandlerFunc(middlewares ...http.HandlerFunc) http.HandlerFunc {
	h := func(w http.ResponseWriter, r *http.Request) {
		captchaToken := r.URL.Query().Get("captcha_token")
		if captchaToken == "" {
			writeJSON(w, `{"err":"missing captcha_token"}`)
			return
		}
		captchaXStr := r.URL.Query().Get("captcha_xs")
		if captchaXStr == "" {
			writeJSON(w, `{"err":"missing captcha_xs"}`)
			return
		}
		captchaX, err := strconv.Atoi(captchaXStr)
		if err != nil {
			writeJSON(w, `{"err":"invalid captcha_xs"}`)
			return
		}
		if ok := Verify(captchaToken, captchaX); !ok {
			writeJSON(w, `{"err":"captcha_err"}`)
			return
		}
		writeJSON(w, `{"verified":"ok"}`)
		return
	}
	middlewares = append(middlewares, h)
	return walkHandlerFunc(middlewares...)
}

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
	"net/http"
	"strconv"

	"github.com/go-the-way/cpt/internal/generator"
	"github.com/go-the-way/cpt/middleware"
)

func generateHandlerFunc(middlewares ...http.HandlerFunc) http.HandlerFunc {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions || r.Method == http.MethodHead {
			return
		}
		imageToken, err := defGtr.Generate()
		if err != nil {
			writeJSON(w, `{"err":"`+err.Error()+`"}`)
			return
		}
		if r.URL.Query().Has("html") {
			writeHTML(w, generateHTML(imageToken))
			return
		}
		writeJSON(w, imageToken.JSON())
	})
	middlewares = append(middlewares, middleware.GzipHandler(h).ServeHTTP)
	return walkHandlerFunc(middlewares...)
}

func verifyHandlerFunc(middlewares ...http.HandlerFunc) http.HandlerFunc {
	h := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions || r.Method == http.MethodHead {
			return
		}
		captchaToken := r.URL.Query().Get("token")
		if captchaToken == "" {
			writeJSON(w, `{"err":"missing token"}`)
			return
		}
		captchaXStr := r.URL.Query().Get("x")
		if captchaXStr == "" {
			writeJSON(w, `{"err":"missing x"}`)
			return
		}
		captchaX, err := strconv.Atoi(captchaXStr)
		if err != nil {
			writeJSON(w, `{"err":"invalid x"}`)
			return
		}
		if ok := Verify(captchaToken, captchaX); !ok {
			writeJSON(w, `{"err":"captcha_err"}`)
			return
		}
		writeJSON(w, `{"verified":"ok"}`)
	}
	middlewares = append(middlewares, h)
	return walkHandlerFunc(middlewares...)
}

func generateHTML(token generator.ImageToken) string {
	return `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Document</title>
</head>
<body>
<h4>` + token.Token + `</h4>
<hr/>
<img src="` + token.BgImageBase64 + `"/>                 
<hr/>
<img src="` + token.BcImageBase64 + `"/>                
</body>
</html>`
}

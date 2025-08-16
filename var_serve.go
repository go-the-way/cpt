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
	"log"
	"net/http"
	"os"

	"github.com/go-the-way/cpt/middleware"
)

func serve(addr string) { ServeRouter(addr, "/cpt/generate", "/cpt/verify") }

func serveRouter(addr, generateRouter, verifyRouter string) {
	logger := log.New(os.Stdout, "[cpt_server] ", log.LstdFlags|log.Lshortfile)
	logger.Println("served on " + addr)
	http.HandleFunc(generateRouter, WrappedGenerateHandlerFunc(middleware.Cors()))
	http.HandleFunc(verifyRouter, WrappedVerifyHandlerFunc(middleware.Cors()))
	logger.Println(http.ListenAndServe(addr, nil))
}

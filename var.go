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
	"github.com/rwscode/cpt/internal/conf"
	"github.com/rwscode/cpt/internal/generator"
	"github.com/rwscode/cpt/internal/resloader"
	"github.com/rwscode/cpt/internal/verifier"
)

var (
	defResLoader = resloader.FsResLoaderDefault()
	defVerifier  = verifier.DefaultVerifier()
	defGtr       = generator.DefaultGenerator(defResLoader, defVerifier)
	Generate     = defGtr.Generate
	Verify       = defGtr.Verify
	Delete       = defGtr.Delete

	SetTokenExpiration       = conf.SetTokenExpiration
	SetTokenClearJobExecTick = conf.SetTokenClearJobExecTick
	SetTokenDeviation        = conf.SetTokenDeviation
	SetTokenLength           = conf.SetTokenLength
	SetResLoaderDefaultOpts  = setResLoaderDefaultOpts

	WrappedGenerateHandlerFunc = generateHandlerFunc
	WrappedVerifyHandlerFunc   = verifyHandlerFunc

	Serve       = serve
	ServeRouter = serveRouter
)

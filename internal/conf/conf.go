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

package conf

import "time"

var (
	tokenExpiration       = time.Minute * 5
	tokenClearJobExecTick = time.Second
	tokenDeviation        = 5
	tokenLength           = 64
)

func GetTokenExpiration() time.Duration           { return tokenExpiration }
func SetTokenExpiration(expiration time.Duration) { tokenExpiration = expiration }
func GetTokenClearJobExecTick() time.Duration     { return tokenClearJobExecTick }
func SetTokenClearJobExecTick(tick time.Duration) { tokenClearJobExecTick = tick }
func GetTokenDeviation() int                      { return tokenDeviation }
func SetTokenDeviation(deviation int)             { tokenDeviation = deviation }
func GetTokenLength() int                         { return tokenLength }
func SetTokenLength(len int)                      { tokenLength = len }

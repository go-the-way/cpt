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

package verifier

import (
	"log"
	"math"
	"os"
	"sync"
	"time"
	
	"github.com/go-the-way/cpt/internal/pkg"
	"github.com/go-the-way/cpt/internal/conf"
)

var (
	logger = log.New(os.Stdout, "[cpt_verifier] ", log.LstdFlags|log.Lshortfile)
)

type expireAt struct {
	x      int
	expire time.Time
}

type defaultVerifier struct {
	mu *sync.RWMutex
	m  map[string]expireAt
}

func DefaultVerifier() *defaultVerifier {
	return (&defaultVerifier{&sync.RWMutex{}, map[string]expireAt{}}).clean()
}

func (d *defaultVerifier) clean() *defaultVerifier { go d.cleanJob(); return d }

func (d *defaultVerifier) cleanJob() {
	logger.Println("token clean job started")
	ticker := time.NewTicker(conf.GetTokenClearJobExecTick())
	defer ticker.Stop()
	for {
		<-ticker.C
		d.mu.Lock()
		for k, v := range d.m {
			// expired
			if time.Now().After(v.expire) {
				delete(d.m, k)
				logger.Println("expired: " + k)
			}
		}
		d.mu.Unlock()
	}
}

func (d *defaultVerifier) Token(len int) (token string) { return pkg.RandomStr(len) }

func (d *defaultVerifier) Store(token string, x int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	expiration := time.Now().Add(conf.GetTokenExpiration())
	d.m[token] = expireAt{x: x, expire: expiration}
	logger.Printf("stored: "+token+" x: %d\n", x)
}

func (d *defaultVerifier) Verify(token string, x int) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	eAt, ok := d.m[token]
	deviation := conf.GetTokenDeviation()
	b := int(math.Abs(float64(eAt.x-x))) <= deviation
	logger.Printf("token: %s x: %d ok: %v eAt.x: %d b: %v deviation: %d\n", token, x, ok, eAt.x, b, deviation)
	return ok && b
}

func (d *defaultVerifier) Delete(token string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.m, token)
	logger.Println("deleted:", token)
}

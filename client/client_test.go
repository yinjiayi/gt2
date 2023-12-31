// Copyright (c) 2022 Institute of Software, Chinese Academy of Sciences (ISCAS)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"testing"
	"time"
)

func TestClientWaitUntilReady(t *testing.T) {
	c, err := New(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	err = c.WaitUntilReady(2 * time.Second)
	if err != errTimeout {
		t.Fatal("err != timeout")
	}
	go func() {
		time.Sleep(time.Second)
		c.addTunnel(&conn{})
	}()
	err = c.WaitUntilReady(30 * time.Second)
	if err == errTimeout {
		t.Fatal("err == timeout")
	}
}

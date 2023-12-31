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

package server

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/isrc-cas/gt/bufio"
)

func TestPeekTLSHost(t *testing.T) {
	bufHexStr := "1603010255010002510303be8eccdd54ca3147ca8a55b52d8d2845b36f114497ef9ac4fc55abd2d9fcf4c020126c4d026115eb79acec6d09228464c2625726468bd7d89974b6922fd05742e70020fafa130113021303c02bc02fc02cc030cca9cca8c013c014009c009d002f0035010001e8baba000000000012001000000d6173736574732e6d736e2e636e00170000ff01000100000a000a0008baba001d00170018000b00020100002300000010000e000c02683208687474702f312e31000500050100000000000d0012001004030804040105030805050108060601001200000033002b0029baba000100001d0020065bc8a4e837100b29783c739d00bdb017ea471ca56a0f7122222c49f5f73311002d00020101002b000b0adada0304030303020301001b0003020002446900050003026832fafa0001000029011b00e600e000004b28939eefd45f2dac0a262083c164a34426309c554007e7336ffbe4dac8d54fd34e1eb219a443821ad8b777b3948e48a2bb6bfa1d73492f3bae723bdd7e1eb29bfc39bbe069ed1c5ac86af6768997752ff37fcc7807a94ad78957af47e6a1a7b3d9989c7996494d5fae7013e32b8ec9058c154943c6de98d1f1dfc578add8e957bd6431d493c854a3a90fe07311be7715f86732a147628b4ce716cf9804d9c0de90ce1604678fdf1f2807711c79c743a84f84931352922a866e1a7ddd4e5f31eb1b2479175be62a1bfb9cabf0ef0813680aac5c61e348d396fd93c4b11c3dfc80b500313018ba4c08126c0a92c9b758254a140bff167022c842e3430541c862ae526dac56a74c280771afd3fdface4deb5655e35a"
	buf, err := hex.DecodeString(bufHexStr)
	if err != nil {
		t.Fatal(err)
	}
	bufReader := bufio.NewReader(bytes.NewReader(buf))
	_, err = bufReader.Peek(len(buf))
	if err != nil {
		t.Fatal(err)
	}

	host, err := peekTLSHost(bufReader)
	if err != nil {
		t.Fatal(err)
	}

	expectedHost := []byte("assets.msn.cn")
	if !bytes.Equal(host, expectedHost) {
		t.Fatalf("unexpected result: %v, expected: %v", string(host), string(expectedHost))
	}
}

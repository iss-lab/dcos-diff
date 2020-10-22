//  MIT License
//
//  Copyright (c) 2020 Institutional Shareholder Services. All other rights reserved.
//
//  Permission is hereby granted, free of charge, to any person obtaining a copy
//  of this software and associated documentation files (the "Software"), to deal
//  in the Software without restriction, including without limitation the rights
//  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:
//
//  The above copyright notice and this permission notice shall be included in all
//  copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//  SOFTWARE.

package util

import (
	"testing"
)

const testAppIDsBlock1 = `/data/arangodb3
/infra/consul
/infra/edgelb/api
/infra/edgelb/pools/main
/kubernetes`

const testAppIDsBlock2 = `/data/arangodb3
/infra/consul
/infra/edgelb/pools/main
/kubernetes
/zookeeper`

const testDiffText = `/data/arangodb3
/infra/consul
/infra/edgelb/api
/infra/edgelb/pools/main
/kubernetes
/zookeeper`

const testDiffHtml = `<span>/data/arangodb3&para;<br>/infra/consul&para;<br>/infra/edgelb/</span><del style="background:#ffe6e6;">api&para;<br>/infra/edgelb/</del><span>pools/main&para;<br>/kubernetes</span><ins style="background:#e6ffe6;">&para;<br>/zookeeper</ins>`

func TestDiffText(t *testing.T) {
	_, diffHtml := DiffText(testAppIDsBlock1, testAppIDsBlock2)
	if diffHtml != testDiffHtml {
		t.Fatalf("Incorrect pretty diff text: %+v", diffHtml)
	}
}

func TestDiffMissing(t *testing.T) {
	missingLeft, missingRight := DiffMissing(BlockToSlice(testAppIDsBlock1), BlockToSlice(testAppIDsBlock2))
	if missingLeft[0] != "/zookeeper" {
		t.Fatalf("Incorrect missing left: %+v", missingLeft)
	}

	if missingRight[0] != "/infra/edgelb/api" {
		t.Fatalf("Incorrect missing right: %+v", missingRight)
	}
}

func TestEnvBytesToMap(t *testing.T) {
	envBytes := GetFileBytes("./testService.env")
	envMap := EnvBytesToMap(envBytes)

	if envMap["PACKAGE_NAME"] != "consul" {
		t.Fatalf("Incorrect env map contents: %+v", envMap)
	}

	if envMap["PACKAGE_VERSION"] != "1.7.3" {
		t.Fatalf("Incorrect env map contents: %+v", envMap)
	}
}

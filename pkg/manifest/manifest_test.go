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

package manifest

import (
	"testing"
)

const testAppIDsBlock = `/data/arangodb3
/infra/consul
/infra/edgelb/api
/infra/edgelb/pools/main
/kubernetes`

func TestNew(t *testing.T) {
	m := New("./testManifest")

	if len(m.AppIDs) != 5 {
		t.Fatalf("Incorrect number of appIDs: %+v", len(m.AppIDs))
	}

	block := m.GetAppIDsBlock()

	if block != testAppIDsBlock {
		t.Fatalf("Incorrect AppIDs block: %+v", block)
	}
}

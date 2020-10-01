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
	"io/ioutil"
	"sort"

	"github.com/iss-lab/dcos-diff/pkg/util"
)

// Manifest is a representation of a total service spec manifest
type Manifest struct {
	AppIDs []string
}

// New returns a populated Manifest object pointer
func New(path string) *Manifest {
	dirs := getDirNames(path, 4, "", nil)
	sort.Strings(dirs)
	manifest := &Manifest{
		AppIDs: dirs,
	}

	return manifest
}

// GetAppIDsBlock returns the joined list of appids
func (m *Manifest) GetAppIDsBlock() string {
	return util.SliceToBlock(m.AppIDs)
}

func getDirNames(basePath string, levels int, prefix string, accum []string) []string {
	if accum == nil {
		accum = []string{}
	}

	infos, err := ioutil.ReadDir(basePath + "/" + prefix)
	util.CheckError(err)

	if len(infos) == 0 {
		accum = append(accum, prefix)
	}

	if levels == 0 {
		return accum
	}

	for _, i := range infos {
		if i.IsDir() {
			accum = getDirNames(basePath, levels-1, prefix+"/"+i.Name(), accum)
		}
	}

	return accum
}

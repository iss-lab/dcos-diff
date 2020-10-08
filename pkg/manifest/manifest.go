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
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	"github.com/iss-lab/dcos-diff/pkg/util"
)

const MarathonFile = "marathon.json"
const FrameworkFile = "options.json"
const PoolFile = "pool.json"
const K8sFile = "k8s.json"

// MarathonSpec describes a service with a marathon.json
type MarathonSpec struct {
	ID     string
	Config []byte
}

// FrameworkSpec describes a service with an options.json
type FrameworkSpec struct {
	ID     string
	Config []byte
}

// PoolSpec describes an EdgeLB Pool definition from a pool.json
type PoolSpec struct {
	ID     string
	Config []byte
}

// K8sSpec describes a MKE Kubernetes cluster definition from a k8s.json
type K8sSpec struct {
	ID     string
	Config []byte
}

// Manifest is a representation of a total service spec manifest
type Manifest struct {
	AppIDs         []string
	MarathonSpecs  *[]MarathonSpec
	FrameworkSpecs *[]FrameworkSpec
	PoolSpecs      *[]PoolSpec
	K8sSpecs       *[]K8sSpec
}

// New returns a populated Manifest object pointer
func New(path string, maxDepth int) *Manifest {
	if maxDepth == 0 {
		maxDepth = 4
	}

	dirs := getDirNames(path, maxDepth, "", nil)
	sort.Strings(dirs)

	fmt.Println(fmt.Errorf("Path: %+v", path))
	fmt.Println(fmt.Errorf("Dirs: %+v", dirs))

	marathonSpecs := &[]MarathonSpec{}
	frameworkSpecs := &[]FrameworkSpec{}
	poolSpecs := &[]PoolSpec{}
	k8sSpecs := &[]K8sSpec{}

	appIDs := []string{}

	for _, d := range dirs {
		if fileExists(path, d, MarathonFile) {
			// Handle Marathon
			appIDs = append(appIDs, d)
			continue
		}
		if fileExists(path, d, FrameworkFile) {
			// Handle Framework
			appIDs = append(appIDs, d)
			continue
		}
		if fileExists(path, d, PoolFile) {
			// Handle Pool
			appIDs = append(appIDs, d)
			continue
		}
		if fileExists(path, d, K8sFile) {
			// Handle K8s
			appIDs = append(appIDs, d)
			continue
		}
	}

	manifest := &Manifest{
		AppIDs:         appIDs,
		MarathonSpecs:  marathonSpecs,
		FrameworkSpecs: frameworkSpecs,
		PoolSpecs:      poolSpecs,
		K8sSpecs:       k8sSpecs,
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

	// TODO: Change this to check a list of ignored dirs
	if prefix == "/.git" {
		return accum
	}

	infos, err := ioutil.ReadDir(basePath + "/" + prefix)
	util.CheckError(err)

	numDirs := 0

	for _, i := range infos {
		if i.IsDir() && levels > 0 {
			numDirs++
			accum = getDirNames(basePath, levels-1, prefix+"/"+i.Name(), accum)
		}
	}

	if numDirs == 0 {
		accum = append(accum, prefix)
	}

	return accum
}

func fileExists(basePath, path, name string) bool {
	if _, err := os.Stat(fmt.Sprintf("%s%s/%s", basePath, path, name)); os.IsNotExist(err) {
		return false
	}
	return true
}

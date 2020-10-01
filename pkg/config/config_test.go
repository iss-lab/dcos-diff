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

package config

import (
	"testing"

	"github.com/iss-lab/dcos-diff/pkg/util"
)

func getTestClusters() []*Cluster {
	b := util.GetFileBytes("./testClusters.json")
	return GetClustersFromBytes(b)
}

func getTestConfig() *Config {
	return New("a0f3a5f2-4465-4990-81bf-8915f7c95e8e", getTestClusters())
}

func TestGetClustersFromBytes(t *testing.T) {
	clusters := getTestClusters()
	if len(clusters) != 2 {
		t.Fatal("Clusters length should be 2")
	}
}

func TestNew(t *testing.T) {
	c := getTestConfig()
	if c.Cluster == nil {
		t.Fatal("Config Cluster should not be nil")
	}
}

func TestEnsureCluster(t *testing.T) {
	c := getTestConfig()
	c.EnsureCluster()

	if !c.Cluster.Attached {
		t.Fatal("Cluster was not attached properly")
	}
}

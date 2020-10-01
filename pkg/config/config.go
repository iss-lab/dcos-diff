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
	"encoding/json"
	"log"
	"os/exec"

	"github.com/iss-lab/dcos-diff/pkg/util"
)

// Config is a represenation of values needed to run dcos-diff
type Config struct {
	Cluster *Cluster
}

// Cluster is a representation of the dcos cli cluster config
type Cluster struct {
	Attached  bool   `json:"attached"`
	ClusterID string `json:"cluster_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	URL       string `json:"url"`
	Version   string `json:"version"`
}

// GetClusters returns list of Clusters
func GetClusters() []*Cluster {
	cmd := exec.Command("dcos", "cluster", "list", "--json")
	output, err := cmd.Output()
	util.CheckError(err)
	return GetClustersFromBytes(output)
}

// GetClustersFromBytes returns list of Clusters from bytes
func GetClustersFromBytes(input []byte) []*Cluster {
	var s []*Cluster
	json.Unmarshal(input, &s)
	return s
}

// New returns a config object given some parameters
func New(clusterID string, clusters []*Cluster) *Config {
	config := &Config{}

	if clusters == nil {
		clusters = GetClusters()
	}

	for _, c := range clusters {
		if c.ClusterID == clusterID {
			config.Cluster = c
		}
	}

	return config
}

// EnsureCluster performs some checks to ensure we're using the correct cluster
func (c *Config) EnsureCluster() {
	if c.Cluster == nil || len(c.Cluster.ClusterID) == 0 {
		log.Fatal("Cluster is not properly set")
	}

	if !c.Cluster.Attached {
		cmd := exec.Command("dcos", "cluster", "attach", c.Cluster.ClusterID)
		_, err := cmd.Output()
		util.CheckError(err)
		c.Cluster.Attached = true
	}

	return
}

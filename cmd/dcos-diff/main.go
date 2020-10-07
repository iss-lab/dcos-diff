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

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/iss-lab/dcos-diff/pkg/util"

	"github.com/iss-lab/dcos-diff/pkg/config"
)

var manifestPath string
var clusterID string

func init() {
	flag.StringVar(
		&manifestPath,
		"manifest_path",
		"./",
		"path to directory containing dcos service manifest")
	flag.StringVar(
		&clusterID,
		"cluster_id",
		"",
		"dcos cluster id to use with dcos cli commands")
}

func main() {
	flag.Parse()

	clusters := config.GetClusters()

	if len(clusterID) == 0 && len(clusters) > 1 {
		printClusters(clusters)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choose a cluster by index: ")

		clusterIndexStr, err := reader.ReadString('\n')
		clusterIndexStr = strings.TrimRight(clusterIndexStr, "\n")
		util.CheckError(err)

		if len(clusterIndexStr) == 0 {
			log.Fatal("empty cluster index given, exiting")
		}

		clusterIndex, err := strconv.Atoi(clusterIndexStr)
		util.CheckError(err)

		// Convert to 0 based index
		clusterIndex--
		if clusterIndex > len(clusters)-1 || clusterIndex < 0 {
			log.Fatal("invalid cluster index given, exiting")
		}
		clusterID = clusters[clusterIndex].ClusterID
	}

	cfg := config.New(clusterID, clusters)

	cfg.EnsureCluster()
}

func printClusters(clusters []*config.Cluster) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "Index", "Name", "URL", "Version", "Cluster ID")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t", "----", "----", "----", "----", "----")

	for i, c := range clusters {
		fmt.Fprintf(w, "\n %d\t%s\t%s\t%s\t%s\t", i+1, c.Name, c.URL, c.Version, c.ClusterID)
	}

	fmt.Fprint(w, "\n\n")
}

func dcosDiff() {

}

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

package marathon

import (
	"encoding/json"
	"os/exec"

	"github.com/iss-lab/dcos-diff/pkg/util"
)

// App json struct
type App struct {
	ID string `json:"id"`
}

// GetAppIDs returns list of App IDs as strings
func GetAppIDs(apps []*App) []string {
	appIDs := []string{}
	for _, a := range apps {
		appIDs = append(appIDs, a.ID)
	}
	return appIDs
}

// GetApps returns list of Apps
func GetApps() []*App {
	cmd := exec.Command("dcos", "marathon", "app", "list", "--json")
	output, err := cmd.Output()
	util.CheckError(err)
	return GetAppsFromBytes(output)
}

// GetAppsFromBytes returns list of Apps from bytes
func GetAppsFromBytes(input []byte) []*App {
	var s []*App
	json.Unmarshal(input, &s)
	return s
}

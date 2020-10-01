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
	"io/ioutil"
	"log"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// CheckError logs a fatal error and exits when an error is passed
func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// DiffText outputs a diff from two strings
func DiffText(left, right string) (string, string) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(left, right, false)
	return dmp.DiffPrettyText(diffs), dmp.DiffPrettyHtml(diffs)
}

// DiffMissing returns the left and right missing strings from 2 string slices
func DiffMissing(left, right []string) ([]string, []string) {
	mapLeft := sliceToMap(left)
	mapRight := sliceToMap(right)

	missingLeft := []string{}
	missingRight := []string{}

	for _, l := range left {
		if !mapRight[l] {
			missingRight = append(missingRight, l)
		}
	}

	for _, r := range right {
		if !mapLeft[r] {
			missingLeft = append(missingLeft, r)
		}
	}

	return missingLeft, missingRight
}

// GetFileBytes returns bytes from a file when there is no error
func GetFileBytes(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	CheckError(err)
	return dat
}

// SliceToBlock takes a slice of strings and joins them with newlines
func SliceToBlock(lines []string) string {
	return strings.Join(lines, "\n")
}

// BlockToSlice takes a slice of strings and joins them with newlines
func BlockToSlice(lines string) []string {
	return strings.Split(lines, "\n")
}

func sliceToMap(lines []string) map[string]bool {
	m := make(map[string]bool, len(lines))
	for _, l := range lines {
		m[l] = true
	}
	return m
}

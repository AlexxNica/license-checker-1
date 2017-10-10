/*
Copyright 2017 Heptio Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/heptiolabs/license-checker/license"
)

func main() {
	supportedLicenses := []string{"apache"}
	licenseType := "apache"
	flag.StringVar(&licenseType, "license", licenseType, fmt.Sprintf("type of license to validate (supported options are %q)", strings.Join(supportedLicenses, ",")))
	check := "Dockerfile,Makefile,.go,.c,.cpp,.py,.sh,.rb,.yaml"
	flag.StringVar(&check, "check", check, "comma-separated list of basenames and extensions to check")

	skip := ".git,vendor,generated"
	flag.StringVar(&skip, "skip", skip, "comma-separated list of directories to skip")

	flag.Parse()

	root := "."
	if len(flag.Args()) > 0 {
		root = flag.Arg(0)
	}

	checks := strings.Split(check, ",")
	skips := strings.Split(skip, ",")

	missing, err := license.ScanDir(root, licenseType, checks, skips)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error running: %v\n", err)
		os.Exit(1)
	}
	if len(missing) == 0 {
		return
	}
	for _, m := range missing {
		fmt.Println(m)
	}
	os.Exit(1)
}

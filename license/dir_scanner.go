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

package license

import (
	"os"
	"path/filepath"
	"strings"
)

// Check recursively examines all the files in root for valid license headers, returning a list of
// invalid files. An error is returned if there is a problem reading any of the files.
func ScanDir(root string, license string, checks, skips []string) ([]string, error) {
	var invalid []string

	var prefixes []string
	extensions := make(map[string]struct{})
	for _, c := range checks {
		if strings.HasPrefix(c, ".") {
			extensions[c] = struct{}{}
		} else {
			prefixes = append(prefixes, c)
		}
	}

	skipsMap := make(map[string]struct{}, len(skips))
	for _, e := range skips {
		skipsMap[e] = struct{}{}
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		base := filepath.Base(path)
		if info.IsDir() {
			if _, skip := skipsMap[base]; skip {
				return filepath.SkipDir
			}

			return nil
		}

		matchedBase := false
		for _, p := range prefixes {
			if strings.HasPrefix(base, p) {
				matchedBase = true
				break
			}
		}

		extension := filepath.Ext(path)
		_, matchedExtension := extensions[extension]

		if !matchedBase && !matchedExtension {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		valid, err := Validate(license, file)
		if err != nil {
			return err
		}
		if !valid {
			invalid = append(invalid, path)
		} else {
		}

		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return invalid, nil
}

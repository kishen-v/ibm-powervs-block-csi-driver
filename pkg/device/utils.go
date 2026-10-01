/*
Copyright 2023 The Kubernetes Authors.

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

package device

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// findStringSubmatchMap finds and builds a map of named capture groups.
func findStringSubmatchMap(s string, r *regexp.Regexp) map[string]string {
	captures := make(map[string]string)
	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}
	for i, name := range r.SubexpNames() {
		if i > 0 && name != "" {
			captures[name] = match[i]
		}
	}
	return captures
}

// readFirstLine reads the first line from filePath.
func readFirstLine(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", nil
}

func getMpathName(pathname string) (string, error) {
	return readFirstLine(filepath.Join("/sys/block", pathname, "dm/name"))
}

func getUUID(pathname string) (string, error) {
	return readFirstLine(filepath.Join("/sys/block", pathname, "dm/uuid"))
}

// deleteSdDevice deletes the SCSI device by writing "1" to its delete sysfs path.
func deleteSdDevice(deletePath string) error {
	if err := os.WriteFile(deletePath, []byte("1"), 0644); err != nil {
		return fmt.Errorf("error writing to file %s: %w", deletePath, err)
	}
	return nil
}

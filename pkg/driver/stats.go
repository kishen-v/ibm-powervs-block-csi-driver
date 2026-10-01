/*
Copyright 2026 The Kubernetes Authors.

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

package driver

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
	"k8s.io/kubernetes/pkg/volume/util/fs"
)

// StatsUtils provides volume and filesystem statistical utilities.
type StatsUtils interface {
	FSInfo(path string) (int64, int64, int64, int64, int64, int64, error)
	IsBlockDevice(devicePath string) (bool, error)
	DeviceInfo(devicePath string) (int64, error)
	IsPathNotExist(path string) bool
}

// VolumeStatUtils implements StatsUtils.
type VolumeStatUtils struct{}

// IsPathNotExist returns true if a particular path does not exist.
func (su *VolumeStatUtils) IsPathNotExist(path string) bool {
	_, err := os.Stat(path)
	return errors.Is(err, os.ErrNotExist) || os.IsNotExist(err)
}

// IsBlockDevice returns true if the path provided is a block device.
func (su *VolumeStatUtils) IsBlockDevice(devicePath string) (bool, error) {
	var stat unix.Stat_t
	if err := unix.Stat(devicePath, &stat); err != nil {
		return false, err
	}
	return (stat.Mode & unix.S_IFMT) == unix.S_IFBLK, nil
}

// DeviceInfo returns the size of the block device in bytes.
func (su *VolumeStatUtils) DeviceInfo(devicePath string) (int64, error) {
	output, err := exec.CommandContext(context.Background(), "blockdev", "--getsize64", devicePath).CombinedOutput()
	if err != nil {
		return 0, fmt.Errorf("failed to get size of block volume at path %s (output: %s): %w", devicePath, strings.TrimSpace(string(output)), err)
	}
	strOut := strings.TrimSpace(string(output))
	gotSizeBytes, err := strconv.ParseInt(strOut, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse size %q into int: %w", strOut, err)
	}

	return gotSizeBytes, nil
}

// FSInfo returns the information related to the FS.
func (su *VolumeStatUtils) FSInfo(path string) (int64, int64, int64, int64, int64, int64, error) {
	return fs.Info(path)
}

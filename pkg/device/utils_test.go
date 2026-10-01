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
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"testing"
)

func TestFindStringSubmatchMap(t *testing.T) {
	re := regexp.MustCompile(`^(?P<name>[a-z]+)-(?P<id>\d+)(?:-(unnamed))?$`)
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			name:     "no match",
			input:    "123-abc",
			expected: map[string]string{},
		},
		{
			name:  "match with named groups",
			input: "volume-12345",
			expected: map[string]string{
				"name": "volume",
				"id":   "12345",
			},
		},
		{
			name:  "match with unnamed group present",
			input: "volume-12345-unnamed",
			expected: map[string]string{
				"name": "volume",
				"id":   "12345",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := findStringSubmatchMap(tt.input, re)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Fatalf("expected %+v, got %+v", tt.expected, res)
			}
		})
	}
}

func TestReadFirstLine(t *testing.T) {
	tmpDir := t.TempDir()

	emptyFile := filepath.Join(tmpDir, "empty.txt")
	if err := os.WriteFile(emptyFile, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}

	singleLineFile := filepath.Join(tmpDir, "single.txt")
	if err := os.WriteFile(singleLineFile, []byte("first line content\n"), 0644); err != nil {
		t.Fatal(err)
	}

	multiLineFile := filepath.Join(tmpDir, "multi.txt")
	if err := os.WriteFile(multiLineFile, []byte("first line\nsecond line\n"), 0644); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		path        string
		expected    string
		expectError bool
	}{
		{
			name:        "non-existent file",
			path:        filepath.Join(tmpDir, "missing.txt"),
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty file",
			path:        emptyFile,
			expected:    "",
			expectError: false,
		},
		{
			name:        "single line",
			path:        singleLineFile,
			expected:    "first line content",
			expectError: false,
		},
		{
			name:        "multi line",
			path:        multiLineFile,
			expected:    "first line",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			line, err := readFirstLine(tt.path)
			if tt.expectError && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tt.expectError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if line != tt.expected {
				t.Fatalf("expected %q, got %q", tt.expected, line)
			}
		})
	}
}

func TestDeleteSdDevice(t *testing.T) {
	tmpDir := t.TempDir()
	deletePath := filepath.Join(tmpDir, "delete")
	if err := os.WriteFile(deletePath, []byte("0"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := deleteSdDevice(deletePath); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	content, err := os.ReadFile(deletePath)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "1" {
		t.Fatalf("expected file content '1', got %q", string(content))
	}

	// Non-existent path returns error
	if err := deleteSdDevice(filepath.Join(tmpDir, "does-not-exist/delete")); err == nil {
		t.Fatal("expected error writing to non-existent path, got nil")
	}
}

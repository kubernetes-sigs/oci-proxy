/*
Copyright 2022 The Kubernetes Authors.

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

package integration

import "testing"

func TestModuleRootDir(t *testing.T) {
	root, err := ModuleRootDir()
	if err != nil {
		t.Fatalf("unexpected error getting root dir: %v", err)
	}
	if root == "" {
		t.Fatalf("expected root dir to be non-empty string")
	}
}

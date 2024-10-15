/*
   Copyright The containerd Authors.

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

package reference

import (
	"github.com/distribution/reference"
)

// Sort sorts string references preferring higher information references.
//
// The precedence is as follows:
//
//  1. [Named] + [Tagged] + [Digested] (e.g., "docker.io/library/busybox:latest@sha256:<digest>")
//  2. [Named] + [Tagged]              (e.g., "docker.io/library/busybox:latest")
//  3. [Named] + [Digested]            (e.g., "docker.io/library/busybo@sha256:<digest>")
//  4. [Named]                         (e.g., "docker.io/library/busybox")
//  5. [Digested]                      (e.g., "docker.io@sha256:<digest>")
//  6. Parse error
func Sort(references []string) []string {
	return reference.Sort(references)
}

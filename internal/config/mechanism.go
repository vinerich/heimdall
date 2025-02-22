// Copyright 2022 Dimitrij Drus <dadrus@gmx.de>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package config

import "github.com/goccy/go-json"

type Mechanism struct {
	ID     string          `koanf:"id"`
	Type   string          `koanf:"type"`
	Config MechanismConfig `koanf:"config"`
}

type MechanismConfig map[string]any

func (in *MechanismConfig) DeepCopyInto(out *MechanismConfig) {
	if in == nil {
		return
	}

	jsonStr, _ := json.Marshal(in)

	// we cannot do anything with an error here as
	// the interface implemented here doesn't support
	// error responses
	json.Unmarshal(jsonStr, out) //nolint:errcheck
}

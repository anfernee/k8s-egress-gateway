/*
Copyright 2020 The Kubernetes authors.

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

package ipmasq

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents ip-masq-agent config file
type Config struct {
	NonMasqCIDRs  []string `json:"nonMasqueradeCIDRs"`
	MasqLinkLocal bool     `json:"masqLinkLocal"`
}

// IPMasqAgent represents a state of the ip-masq-agent
type IPMasqAgent struct {
	configPath string
}

// New creates  an IPMasqAgent instance
func New(configPath string) *IPMasqAgent {
	return &IPMasqAgent{
		configPath: configPath,
	}
}

// Update updates ip-masq-agent config file
func (a *IPMasqAgent) Update(config *Config) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(a.configPath, data, 0644)
}

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
	"os"
	"testing"
)

func TestIPMasqAgent(t *testing.T) {
	f, err := ioutil.TempFile("", "testfile-ip-masq-agent")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	agent := New(f.Name())

	c := &Config{
		NonMasqCIDRs: []string{
			"10.0.0.0/8",
			"192.168.0.0/16",
		},
	}
	if err := agent.Update(c); err != nil {
		t.Error(err)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	expectedData := `nonmasqcidrs:
- 10.0.0.0/8
- 192.168.0.0/16
masqlinklocal: false
`
	if string(data) != expectedData {
		t.Errorf("expected data %q; got %q", expectedData, string(data))
	}
}

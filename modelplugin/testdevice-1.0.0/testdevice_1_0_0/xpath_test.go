// Copyright 2022-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testdevice_1_0_0

import (
	"fmt"
	"github.com/antchfx/xpath"
	"github.com/onosproject/config-models/pkg/xpath/navigator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_XPathSelect(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice-1-config.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathSelect{
		{
			Name: "test key1",
			Path: "/t1:cont1a/t1e:list5/@t1e:key1",
			Expected: []string{
				"Iter Value: key1: eight",
				"Iter Value: key1: five",
				"Iter Value: key1: five",
				"Iter Value: key1: two",
			},
		},
		{
			Name: "test key2",
			Path: "/t1:cont1a/t1e:list5/@t1e:key2",
			Expected: []string{
				"Iter Value: key2: 8",
				"Iter Value: key2: 6",
				"Iter Value: key2: 7",
				"Iter Value: key2: 1",
			},
		},
		{
			Name: "test key2 eight",
			Path: "/t1:cont1a/t1e:list5[@t1e:key1='eight'][@t1e:key2=8]/t1e:leaf5a",
			Expected: []string{
				"Iter Value: leaf5a: 5a eight-8",
			},
		},
		{
			Name: "test list4",
			Path: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a1",
			},
		},
		{
			Name: "test list4",
			Path: "/t1:cont1a/t1e:list4[@t1e:id='l2a2']/t1e:leaf4b",
			Expected: []string{
				"Iter Value: leaf4b: this is list4-l2a2",
			},
		},
		{
			Name: "test list4 1 list4a",
			Path: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a",
			Expected: []string{
				"Iter Value: list4a: value of list4a",
				"Iter Value: list4a: value of list4a",
			},
		},
		{
			Name: "test list4 1 list4a displayname",
			Path: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-6",
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
		{
			Name: "test list4 1 list4a fives displayname",
			Path: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a[@t1e:fkey1='five']/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-6",
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
		{
			Name: "test list4 1 list4a 1 displayname",
			Path: "/t1:cont1a/t1e:list4[@t1e:id='l2a1']/t1e:list4a[@t1e:fkey1='five'][@t1e:fkey2=7]/t1e:displayname",
			Expected: []string{
				"Iter Value: displayname: Value l2a1-five-7",
			},
		},
	}

	for _, test := range tests {
		expr, err := xpath.Compile(test.Path)
		assert.NoError(t, err, test.Name)
		assert.NotNil(t, expr, test.Name)

		iter := expr.Select(ynn)
		resultCount := 0
		for iter.MoveNext() {
			assert.LessOrEqual(t, resultCount, len(test.Expected)-1, test.Name, ". More results than expected")
			assert.Equal(t, test.Expected[resultCount], fmt.Sprintf("Iter Value: %s: %s",
				iter.Current().LocalName(), iter.Current().Value()), test.Name)
			resultCount++
		}
		assert.Equal(t, len(test.Expected), resultCount, "%s. Did not receive all the expected results", test.Name)
	}
}

func Test_XPathEvaluate(t *testing.T) {
	sampleConfig, err := ioutil.ReadFile("../testdata/sample-testdevice-1-config.json")
	if err != nil {
		assert.NoError(t, err)
	}
	device := new(Device)

	schema, err := Schema()
	if err := schema.Unmarshal(sampleConfig, device); err != nil {
		assert.NoError(t, err)
	}
	schema.Root = device
	assert.NotNil(t, device)
	ynn := navigator.NewYangNodeNavigator(schema.RootSchema(), device)
	assert.NotNil(t, ynn)

	tests := []navigator.XpathEvaluate{
		{
			Name:     "test get key1",
			Path:     "count(/t1:cont1a/t1e:list5/@t1e:key1)",
			Expected: float64(4),
		},
		{
			Name:     "test get key1 for 'five'", // There are 2 fives
			Path:     "count(/t1:cont1a/t1e:list5[@t1e:key1='five']/@t1e:key1)",
			Expected: float64(2),
		},
		{
			Name:     "test extract key1 for five",
			Path:     "string(/t1:cont1a/t1e:list5[@t1e:key1='five'][@t1e:key2=7]/@t1e:key1)",
			Expected: "five",
		},
		{
			Name:     "test extract key2 for five",
			Path:     "number(/t1:cont1a/t1e:list5[@t1e:key1='five'][@t1e:key2=7]/@t1e:key2)",
			Expected: float64(7),
		},
		{
			Name: "test concat string",
			Path: "concat(concat('5e ', string(/t1:cont1a/t1e:list5[@t1e:key1='five'][@t1e:key2=7]/@t1e:key1)), " +
				"concat('-', string(/t1:cont1a/t1e:list5[@t1e:key1='five'][@t1e:key2=7]/@t1e:key2)))",
			Expected: "5e five-7",
		},
	}

	for _, test := range tests {
		expr, testErr := xpath.Compile(test.Path)
		assert.NoError(t, testErr, test.Name)
		assert.NotNil(t, expr, test.Name)

		result := expr.Evaluate(ynn)
		assert.Equal(t, test.Expected, result, test.Name)
	}
}

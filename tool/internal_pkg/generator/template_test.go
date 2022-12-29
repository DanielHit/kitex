// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"text/template"

	"github.com/cloudwego/kitex/internal/test"
)

func TestNilSafe(t *testing.T) {
	var q *TemplateExtension
	fn := "/tmp/kitex-template-nil.json"

	err := q.ToJSONFile(fn)
	test.Assert(t, err == nil, err)

	err = q.FromJSONFile(fn)
	test.Assert(t, err == nil, err)
}

var tmpl = `package {{ .Package }}

import (
	{{ range $_, $v := .Import -}}
		"{{ $v }}"
	{{ end -}}
)

func main() {
	fmt.Println("{{ .Content }}")
	for i := 0; i <  {{ .LoopTimes}};  i ++ {
		fmt.Println("{{ .Value }}")
	}
}`

func TestWriteFile(t *testing.T) {
	b := &bytes.Buffer{}

	data := struct {
		Package   string
		LoopTimes int64
		Value     string
		Import    []string
		Content   string
	}{
		Package:   "main",
		LoopTimes: 10,
		Value:     "what a beautiful world!",
		Import:    []string{"fmt"},
		Content:   "Hello, World!",
	}

	err := template.Must(template.New("test").Parse(tmpl)).Execute(b, data)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(b.String())
	helloFile, err := os.Create("/tmp/hello.go")
	if err != nil {
		return
	}

	// loop for write
	_, err = helloFile.Write(b.Bytes())
	if err != nil {
		return
	}

}

func TestMarshal(t *testing.T) {
	p := &TemplateExtension{
		FeatureNames:   []string{"f1", "f2"},
		EnableFeatures: []string{"f1"},
		Dependencies: map[string]string{
			"example.com/demo": "demo",
			"example.com/test": "test2",
		},
		ExtendClient: &APIExtension{
			ImportPaths:  []string{"example.com/demo"},
			ExtendOption: "option...",
			ExtendFile:   "file...",
		},
		ExtendServer: &APIExtension{
			ImportPaths:  []string{"example.com/demo"},
			ExtendOption: "option...",
			ExtendFile:   "file...",
		},
		ExtendInvoker: &APIExtension{
			ImportPaths:  []string{"example.com/demo"},
			ExtendOption: "option...",
			ExtendFile:   "file...",
		},
	}

	fn := "/tmp/kitex-template.json"
	err := p.ToJSONFile(fn)
	test.Assert(t, err == nil, err)

	q := new(TemplateExtension)
	err = q.FromJSONFile(fn)
	test.Assert(t, err == nil, err)

	test.DeepEqual(t, p, q)
}

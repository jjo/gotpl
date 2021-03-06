package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYamlTemplate(t *testing.T) {
	type io struct {
		Input    string
		Template string
		Output   string
	}

	tests := []io{
		io{
			Input:    "test: value",
			Template: "{{.test}}",
			Output:   "value",
		},
		io{
			Input:    "name: Max\nage: 15",
			Template: "Hello {{.name}}, of {{.age}} years old",
			Output:   "Hello Max, of 15 years old",
		},
		io{
			Input:    "legumes:\n  - potato\n  - onion\n  - cabbage",
			Template: "Legumes:{{ range $index, $el := .legumes}}{{if $index}},{{end}} {{$el}}{{end}}",
			Output:   "Legumes: potato, onion, cabbage",
		},
		// sprig:
		io{
			Input:    "legumes:\n  - potato\n  - onion\n  - cabbage",
			Template: `Legumes: {{ .legumes | join ", " }}`,
			Output:   "Legumes: potato, onion, cabbage",
		},
		io{
			Input:    "legumes: potato onion cabbage",
			Template: `Legumes: {{ .legumes | splitList " " | join ", " }}`,
			Output:   "Legumes: potato, onion, cabbage",
		},
		io{
			Input:    "legumes: potato onion cabbage",
			Template: "Legumes: {{ .legumes | b64enc | b64dec }}",
			Output:   "Legumes: potato onion cabbage",
		},
	}

	for _, test := range tests {
		tplFile, err := ioutil.TempFile("", "")
		assert.Nil(t, err)
		defer func() { os.Remove(tplFile.Name()) }()

		_, err = tplFile.WriteString(test.Template)
		assert.Nil(t, err)
		tplFile.Close()

		output := bytes.NewBuffer(nil)
		err = ExecuteTemplates(strings.NewReader(test.Input), output,
			[]string{tplFile.Name()})
		assert.Nil(t, err)

		assert.Equal(t, test.Output, output.String())

	}
}

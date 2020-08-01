package runner

import (
	"bytes"
	"text/template"
)

var __program_template string = `
package main

import (
{{ range $import := .imports }}
    "{{ $import }}"
{{ end }}
)

func main() {
{{ range $line := .lines }}
    {{ $line }}
{{ end }}
}
`

func buildProgram(imports, lines []string) ([]byte, error) {
	tmpl, err := template.New("program").Parse(__program_template)
	if err != nil {
		return nil, err
	}

	out := bytes.NewBuffer(nil)
	data := map[string]interface{}{
		"imports": imports,
		"lines":   lines,
	}

	if err := tmpl.Execute(out, data); err != nil {
		return nil, err
	}

	b, err := goimports(out.Bytes())
	if err != nil {
		return nil, err
	}

	// goimports do not always remove every unused package at once,
	// so using it twice is recommended.
	return goimports(b)
}

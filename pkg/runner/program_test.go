package runner

import (
	"fmt"
	"strings"
	"testing"
)

func TestBuildProgram(t *testing.T) {
	imports := []string{"fmt", "math"}
	lines := []string{"fmt.Println(123)", `fmt.Println("blablabla")`}

	expected := `
package main

import (
	"fmt"
)

func main() {

	fmt.Println(123)

	fmt.Println("blablabla")

}
`

	out, err := buildProgram(imports, lines)

	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if strings.TrimSpace(string(out)) != strings.TrimSpace(expected) {
		fmt.Println("Program is different")
		t.Fail()
	}

}

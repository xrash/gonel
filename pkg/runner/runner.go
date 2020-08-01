package runner

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Runner struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Tempdir string
}

func (r *Runner) Run(imports, lines []string) error {

	b, err := buildProgram(imports, lines)
	if err != nil {
		return err
	}

	basename := generateBasename()
	filename := filepath.Join(r.Tempdir, basename)

	err = ioutil.WriteFile(filename, b, os.ModePerm)
	if err != nil {
		return err
	}

	defer func() {
		if err := os.Remove(filename); err != nil {
			panic(err)
		}
	}()

	cmd := exec.Command("go", "run", filename)
	cmd.Stdin = r.Stdin
	cmd.Stdout = r.Stdout
	cmd.Stderr = r.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}

package runner

import (
	"bytes"
	"fmt"
	"os/exec"
)

func goimports(in []byte) ([]byte, error) {
	inreader := bytes.NewBuffer(in)
	outwriter := bytes.NewBuffer(nil)
	errwriter := bytes.NewBuffer(nil)

	cmd := exec.Command("goimports")
	cmd.Stdin = inreader
	cmd.Stdout = outwriter
	cmd.Stderr = errwriter

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	if errwriter.Len() > 0 {
		return nil, fmt.Errorf("%s", errwriter.String())
	}

	return outwriter.Bytes(), nil
}

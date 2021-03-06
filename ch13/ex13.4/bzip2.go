package bzip

import (
	"io"
	"os/exec"
)

type writer struct {
	cmd exec.Cmd
	stdin io.WriteCloser
}

func NewWriter(w io.Writer) (io.WriteCloser, error) {
	cmd := exec.Cmd{Path: "/bin/bzip2", Stdout: w}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return &writer{cmd, stdin}, nil
}

func (w *writer) Write(data []byte) (int, error) {
	return w.stdin.Write(data)
}

func (w *writer) Close() error {
	pipeErr := w.stdin.Close()
	cmdErr := w.cmd.Wait()
	if pipeErr != nil {
		return pipeErr
	}

	if cmdErr != nil {
		return cmdErr
	}

	return nil
}


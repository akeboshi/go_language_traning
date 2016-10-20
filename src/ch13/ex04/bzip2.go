package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	/*	filename := os.Args[1]
		err := exec.Command("/usr/bin/bzip2", filename).Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}*/
	w := NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}

type writer struct {
	w   io.Writer
	tmp *os.File
	cmd *exec.Cmd
}

func NewWriter(out io.Writer) io.WriteCloser {
	filename := "/tmp/__tmp_go"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	cmd := exec.Command("bzip2", "-c", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	w := &writer{out, file, cmd}
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	fileWriter := bufio.NewWriter(w.tmp)
	nn, err := fileWriter.Write(data)
	if err != nil {
		return nn, err
	}
	err = fileWriter.Flush()
	return nn, err
}

func (w *writer) Close() error {
	err := w.cmd.Start()
	if err != nil {
		return err
	}
	return w.tmp.Close()
}

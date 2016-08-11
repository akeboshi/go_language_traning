// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package mytar

import (
	"archive/tar"
	"bytes"
	"ch10/ex02/myarchive"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func init() {
	myarchive.RegisterFormat("tar", detectTar, Untar)
}

func detectTar(b []byte) bool {
	if len(b) < 265 {
		return false
	}
	if b[257] == 'u' && b[258] == 's' && b[259] == 't' && b[260] == 'a' && b[261] == 'r' {
		if b[262] == '\000' ||
			(b[262] == '\040' && b[263] == '\040') {
			return true
		}
	}
	return false
}

func Untar(b []byte, dest string) error {
	tr := tar.NewReader(bytes.NewReader(b))

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		path := filepath.Join(dest, hdr.Name)
		if hdr.FileInfo().IsDir() {
			os.MkdirAll(path, hdr.FileInfo().Mode())
		} else {
			fmt.Printf("extructing %s\n", path)
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, hdr.FileInfo().Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err = io.Copy(f, tr); err != nil {
				log.Fatalln(err)
			}
		}
	}
	return nil
}

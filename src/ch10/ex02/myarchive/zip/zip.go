// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package myzip

import (
	"archive/zip"
	"ch10/ex02/myarchive"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func init() {
	myarchive.RegisterFormat("zip", detectZip, Unzip)
}

func detectZip(b []byte) bool {
	if b[0] == 'P' && b[1] == 'K' {
		if (b[2] == '\003' && b[3] == '\004') ||
			(b[2] == '\005' && b[3] == '\006') ||
			(b[2] == '\007' && b[3] == '\010') {
			return true
		}
	}
	return false
}

type zipReaderAt []byte

func (zra zipReaderAt) ReadAt(b []byte, off int64) (int, error) {
	copy(b, zra[int(off):int(off)+len(b)])
	return len(b), nil
}

func Unzip(b []byte, dest string) error {

	zr, err := zip.NewReader(zipReaderAt(b), int64(len(b)))

	if err != nil {
		return err
	}
	for _, f := range zr.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			fmt.Printf("extructing %s\n", path)
			f, err := os.OpenFile(
				path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

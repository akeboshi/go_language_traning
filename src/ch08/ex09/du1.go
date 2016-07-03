// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

type fs struct {
	size int64
	root string
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan fs)

	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes, root)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles int64
	var nbytes = make(map[string]int64)
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}

			nfiles++
			nbytes[size.root] += size.size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	for size := range fileSizes {
		nfiles++
		nbytes[size.root] *= size.size
	}
	printDiskUsage(nfiles, nbytes)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- fs, root string) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes, root)
		} else {
			fileSizes <- fs{entry.Size(), root}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles int64, nbytes map[string]int64) {
	for k, v := range nbytes {
		fmt.Printf("%s\t:%d files %.1f GB\n", k, nfiles, float64(v)/1e9)
	}
}

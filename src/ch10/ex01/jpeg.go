// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

type availableMethod struct {
	kind    string
	encoder func(io.Writer, image.Image) error
	enable  *bool
}

func main() {
	var availableMethods = []availableMethod{
		{"jpeg", toJPEG, flag.Bool("f", false, "add output type is jpeg")},
		{"png", png.Encode, flag.Bool("p", false, "add output type is png")},
		{"gif", toGIF, flag.Bool("g", false, "add output type is gif")},
	}
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("Please input a image source.")
	img, err := decode(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	outputFiles(img, availableMethods)
}

func outputFiles(img image.Image, availableMethods []availableMethod) {
	for _, am := range availableMethods {
		if *am.enable {
			writeFile, err := os.Create("output." + am.kind)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				continue
			}
			defer writeFile.Close()
			writer := bufio.NewWriter(writeFile)
			if err := am.encoder(writer, img); err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", am.kind, err)
				continue
			}
			writer.Flush()
		}
	}
}

func decode(in io.Reader) (image.Image, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return img, nil
}

func toJPEG(out io.Writer, img image.Image) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toGIF(out io.Writer, img image.Image) error {
	return gif.Encode(out, img, &gif.Options{})
}

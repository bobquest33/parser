package main

import (
	"archive/zip"
	"fmt"
)

var (
	NoEPUBError     = fmt.Errorf("File does not seem to be an EPUB file")
	InvalidXMLError = fmt.Errorf("EPUB file contains invalid XML")
	UnexpectedError = fmt.Errorf("An unexpected error happened")
)

type EPUB struct{}

type epubFile struct {
	data *EPUB
	r    *zip.ReadCloser
}

func Parse(path string) (*EPUB, error) {
	r, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}

	ef := &epubFile{&EPUB{}, r}
	defer ef.r.Close()

	container, err := parseContainer(ef)
	if err != nil {
		return nil, err
	}
	fmt.Println(container.OEBPSPackagePath)

	return nil, nil
}

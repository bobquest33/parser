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

type EPUB struct {
	Version string
	Titles  []string
}

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

	c, err := parseContainer(ef)
	if err != nil {
		return nil, err
	}

	err = parseOEBPSPackage(ef, c)
	if err != nil {
		return nil, err
	}

	return ef.data, nil
}

func (e *EPUB) String() string {
	return fmt.Sprintf(
		"Version: %s\nTitles: %v",
		e.Version,
		e.Titles,
	)

}

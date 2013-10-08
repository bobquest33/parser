package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
)

var (
	NoEPUBError     = fmt.Errorf("File does not seem to be an EPUB file")
	InvalidXMLError = fmt.Errorf("EPUB file contains invalid XML")
	UnexpectedError = fmt.Errorf("An unexpected error happened")
)

type Person struct {
	Name   string
	FileAs string
	Role   string
}

type EPUB struct {
	Version  string
	Titles   []string
	Creators []*Person
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

func (c *Person) String() string {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("Could not convert Person to string: %s", err)
	}
	return fmt.Sprintf("&Person%s", data)
}

func (e *EPUB) String() string {
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return fmt.Sprintf("Could not convert EPUB to string: %s", err)
	}
	return fmt.Sprintf("&EPUB%s", data)
}

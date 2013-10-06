package main

import (
	"archive/zip"
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

	return nil, nil
}

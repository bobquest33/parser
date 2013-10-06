package main

import (
	"archive/zip"
	"testing"
)

func tEF() *epubFile {
	r, _ := zip.OpenReader("test/metamorphosis.epub")
	return &epubFile{&EPUB{}, r}
}

func TestParseGeneral(t *testing.T) {
	ef := tEF()
	defer ef.r.Close()

	_, err := parseContainer(ef)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOEBPSPackagePath(t *testing.T) {
	ef := tEF()
	defer ef.r.Close()

	c, _ := parseContainer(ef)

	actual := c.OEBPSPackagePath
	expected := "5200/content.opf"
	if actual != expected {
		t.Errorf("expected OEBPSPackagePath to be %s, got %s", expected, actual)
	}
}

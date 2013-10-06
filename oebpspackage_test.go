package main

import (
	"archive/zip"
	"testing"

	"github.com/reapub/parser/testutil"
)

func mEF() *epubFile {
	r, _ := zip.OpenReader("test/metamorphosis.epub")
	return &epubFile{&EPUB{}, r}
}

func mC() *container {
	return &container{"5200/content.opf"}
}

func TestParseMetadata(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	_, err := parseMetadata(ef, c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMetadataTitles(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	m, _ := parseMetadata(ef, c)

	actual := m.Titles
	expected := []string{"Metamorphosis"}
	if !testutil.StrSliceEquals(actual, expected) {
		t.Errorf("expected Titles to be %v, got %v", expected, actual)
	}
}

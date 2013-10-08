package main

import (
	"archive/zip"
	"testing"

	"github.com/bmizerany/assert"
)

func mEF() *epubFile {
	r, _ := zip.OpenReader("test/metamorphosis.epub")
	return &epubFile{&EPUB{}, r}
}

func mC() *container {
	return &container{"5200/content.opf"}
}

func TestParseOEBPSPackage(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	err := parseOEBPSPackage(ef, c)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOEBPSVersion(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	parseOEBPSPackage(ef, c)

	actual := ef.data.Version
	expected := "2.0"
	assert.Equal(t, expected, actual)
}

func TestMetadataTitles(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	parseOEBPSPackage(ef, c)

	actual := ef.data.Titles
	expected := []string{"Metamorphosis"}
	assert.Equal(t, expected, actual)
}

func TestMetadataCreators(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	parseOEBPSPackage(ef, c)

	actual := ef.data.Creators
	expected := []*Person{&Person{
		Name:   "Franz Kafka",
		FileAs: "Kafka, Franz",
	}}
	assert.Equal(t, expected, actual)
}

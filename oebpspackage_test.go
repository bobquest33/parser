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

func TestMetadataContributors(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	parseOEBPSPackage(ef, c)

	actual := ef.data.Contributors
	expected := []*Person{&Person{
		Name:   "David Wyllie",
		FileAs: "Wyllie, David",
		Role:   "trl",
	}}
	assert.Equal(t, expected, actual)
}

func TestMetadataSubjects(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	parseOEBPSPackage(ef, c)

	actual := ef.data.Subjects
	expected := []string{"Psychological fiction", "Metamorphosis -- Fiction"}
	assert.Equal(t, expected, actual)
}

func TestMetadataDescription(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	parseOEBPSPackage(ef, c)

	actual := ef.data.Description
	expected := "Classic story of self-discovery, told in a unique manner by Kafka."
	assert.Equal(t, expected, actual)
}

func TestMetadataPublisher(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	parseOEBPSPackage(ef, c)

	actual := ef.data.Publisher
	expected := "Random House"
	assert.Equal(t, expected, actual)
}

func TestMetadataDates(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	parseOEBPSPackage(ef, c)

	actual := ef.data.Dates
	expected := []*Date{
		&Date{Date: "2005-08-17", Event: "publication"},
		&Date{Date: "2013-03-13T11:10:00.924800+00:00", Event: "conversion"},
		&Date{Date: "2012-01-18T12:47:00Z", Event: "modified"},
	}
	assert.Equal(t, expected, actual)
}

func TestMetadataIdentifiers(t *testing.T) {
	ef, c := mEF(), mC()
	defer ef.r.Close()

	parseOEBPSPackage(ef, c)

	actual := ef.data.Identifiers
	expected := []*Identifier{
		&Identifier{Identifier: "http://www.gutenberg.org/ebooks/5200", Scheme: "URI"},
		&Identifier{Identifier: "9781479157303", Scheme: "ISBN"},
	}
	assert.Equal(t, expected, actual)
}

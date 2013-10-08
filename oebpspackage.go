package main

import (
	"fmt"
	"io/ioutil"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

type metadata struct {
	Titles []string
}

type oebpspackage struct {
	Version  string
	Metadata *metadata
}

func parseOEBPSPackage(ef *epubFile, c *container) (*oebpspackage, error) {
	m := &oebpspackage{}

	file := findZipFile(ef.r, c.OEBPSPackagePath)
	if file == nil {
		return nil, UnexpectedError
	}

	fr, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("could not open %s, %s", c.OEBPSPackagePath, err)
	}
	defer fr.Close()

	data, err := ioutil.ReadAll(fr)
	if err != nil {
		return nil, UnexpectedError
	}
	doc, err := gokogiri.ParseXml(data)
	if err != nil {
		return nil, InvalidXMLError
	}
	defer doc.Free()
	doc.RecursivelyRemoveNamespaces()

	m.Version = doc.Root().Attr("version")

	res, _ := doc.Search("/package/metadata")
	if len(res) != 1 {
		return nil, NoEPUBError
	}

	mn := res[0]
	m.Metadata = parseMetadata(mn)

	return m, nil
}

func parseMetadata(m xml.Node) *metadata {
	metadata := &metadata{}

	res, _ := m.Search("title")
	for _, n := range res {
		metadata.Titles = append(metadata.Titles, n.Content())
	}

	return metadata
}

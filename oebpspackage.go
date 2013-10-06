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

func parseMetadata(ef *epubFile, c *container) (*metadata, error) {
	m := &metadata{}

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

	res, _ := doc.Search("/package/metadata")
	if len(res) != 1 {
		return nil, NoEPUBError
	}

	mn := res[0]
	m.Titles = parseTitles(mn)

	return m, nil
}

func parseTitles(m xml.Node) []string {
	titles := []string{}

	res, _ := m.Search("title")
	for _, n := range res {
		titles = append(titles, n.Content())
	}

	return titles
}

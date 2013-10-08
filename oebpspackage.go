package main

import (
	"fmt"
	"io/ioutil"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

func parseOEBPSPackage(ef *epubFile, c *container) error {
	epub := ef.data

	file := findZipFile(ef.r, c.OEBPSPackagePath)
	if file == nil {
		return UnexpectedError
	}

	fr, err := file.Open()
	if err != nil {
		return fmt.Errorf("could not open %s, %s", c.OEBPSPackagePath, err)
	}
	defer fr.Close()

	data, err := ioutil.ReadAll(fr)
	if err != nil {
		return UnexpectedError
	}
	doc, err := gokogiri.ParseXml(data)
	if err != nil {
		return InvalidXMLError
	}
	defer doc.Free()
	doc.RecursivelyRemoveNamespaces()

	epub.Version = doc.Root().Attr("version")

	res, _ := doc.Search("/package/metadata")
	if len(res) != 1 {
		return NoEPUBError
	}

	mn := res[0]
	epub.Titles = parseTitles(mn)

	return nil
}

func parseTitles(m xml.Node) []string {
	titles := []string{}

	res, _ := m.Search("title")
	for _, n := range res {
		titles = append(titles, n.Content())
	}

	return titles
}

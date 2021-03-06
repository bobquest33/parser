package main

import (
	"fmt"
	"io/ioutil"

	"github.com/moovweb/gokogiri"
)

type container struct {
	OEBPSPackagePath string
}

func parseContainer(ef *epubFile) (*container, error) {
	c := &container{}

	rootpath := "META-INF/container.xml"
	rootfile := findZipFile(ef.r, rootpath)
	if rootfile == nil {
		return nil, NoEPUBError
	}

	fr, err := rootfile.Open()
	if err != nil {
		return nil, fmt.Errorf("could not open %s, %s", rootpath, err)
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

	res, _ := doc.Search("/container/rootfiles/rootfile")
	for _, node := range res {
		if node.Attr("media-type") == "application/oebps-package+xml" {
			c.OEBPSPackagePath = node.Attr("full-path")
		}
	}

	return c, nil
}

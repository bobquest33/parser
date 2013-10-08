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

	creators, _ := mn.Search("creator")
	contributors, _ := mn.Search("contributor")

	epub.Titles = parseTitles(mn)
	epub.Creators = parsePeople(creators)
	epub.Contributors = parsePeople(contributors)
	epub.Subjects = parseSubjects(mn)
	epub.Description = parseDescription(mn)

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

func parsePeople(s []xml.Node) []*Person {
	people := []*Person{}

	for _, n := range s {
		person := &Person{
			Name:   n.Content(),
			FileAs: n.Attr("file-as"),
			Role:   n.Attr("role"),
		}
		people = append(people, person)
	}
	return people
}

func parseSubjects(m xml.Node) []string {
	subjects := []string{}

	res, _ := m.Search("subject")
	for _, n := range res {
		subjects = append(subjects, n.Content())
	}

	return subjects
}

func parseDescription(m xml.Node) string {
	description := ""

	res, _ := m.Search("description")
	if len(res) > 0 {
		description = res[0].Content()
		fmt.Println("description:", description)
	}

	return description
}

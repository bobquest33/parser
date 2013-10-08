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
	epub.Publisher = parsePublisher(mn)
	epub.Dates = parseDates(mn)
	epub.Identifiers = parseIdentifiers(mn)
	epub.Source = parseSource(mn)
	epub.Languages = parseLanguages(mn)

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
	}

	return description
}

func parsePublisher(m xml.Node) string {
	publisher := ""

	res, _ := m.Search("publisher")
	if len(res) > 0 {
		publisher = res[0].Content()
	}

	return publisher
}

func parseDates(m xml.Node) []*Date {
	dates := []*Date{}

	res, _ := m.Search("date")
	for _, n := range res {
		date := Date{Date: n.Content(), Event: n.Attr("event")}
		dates = append(dates, &date)
	}

	res, _ = m.Search("meta[@property='dcterms:modified']")
	if len(res) > 0 {
		date := Date{Date: res[0].Content(), Event: "modified"}
		dates = append(dates, &date)
	}

	return dates
}

func parseIdentifiers(m xml.Node) []*Identifier {
	identifiers := []*Identifier{}

	res, _ := m.Search("identifier")
	for _, n := range res {
		identifier := Identifier{Identifier: n.Content(), Scheme: n.Attr("scheme")}
		identifiers = append(identifiers, &identifier)
	}

	return identifiers
}

func parseSource(m xml.Node) string {
	res, _ := m.Search("source")

	if len(res) > 0 {
		return res[0].Content()
	}

	return ""
}

func parseLanguages(m xml.Node) []string {
	languages := []string{}

	res, _ := m.Search("language")
	for _, n := range res {
		languages = append(languages, n.Content())
	}

	return languages
}

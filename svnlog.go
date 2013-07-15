package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os/exec"
)

type Data struct {
	XMLName xml.Name `xml:"log"`
	Entry   []Entry  `xml:"logentry"`
}

type Entry struct {
	Revision int    `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
	Message  string `xml:"msg"`
	Paths    Paths  `xml:"paths"`
}

type Paths struct {
	Path []Path `xml:"path"`
}

type Path struct {
	Action string `xml:"action,attr"`
	Kind   string `xml:"kind,attr"`
	Value  string `xml:",chardata"`
}

func main() {
	xmlFile, err := exec.Command("/usr/bin/svn", "log", "-v", "--xml").Output()
	if err != nil {
		log.Fatal(err)
	}
	var data Data
	xml.Unmarshal([]byte(xmlFile), &data)
	for index := range data.Entry {
		fmt.Printf("%+v\n", data.Entry[index])
	}
}

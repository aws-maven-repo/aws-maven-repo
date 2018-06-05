package main

import (
	"fmt"
	"encoding/xml"
    "io/ioutil"
)

type Project struct {
	GroupId       string `xml:"groupId"`
	ArtifactId    string `xml:"artifactId"`
	Version       string `xml:"version"`
}

func decodePom(file string, p Project) {
	xmlFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	xml.Unmarshal(xmlFile, &p)
	fmt.Println(p.GroupId)
	fmt.Println(p.ArtifactId)
	fmt.Println(p.Version)
}

func main() {
	var p Project
	decodePom("test/pom.xml", p)
}

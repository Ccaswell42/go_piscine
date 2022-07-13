package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Cakes struct {
	XMLName xml.Name `xml:"recipes" json:"-"`
	Cakes   []struct {
		Name        string `json:"name" xml:"name"`
		Time        string `json:"time" xml:"stovetime"`
		Ingredients []struct {
			Name  string `json:"ingredient_name" xml:"itemname"`
			Count string `json:"ingredient_count" xml:"itemcount"`
			Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit"`
		} `json:"ingredients" xml:"ingredients>item"`
	} `json:"cake" xml:"cake"`
}

type DBReader interface {
	convert() error
}

type JsonFile struct {
	File_name string
}

func (r JsonFile) convert() error {
	var str Cakes
	readFile, err := ioutil.ReadFile(r.File_name)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(readFile, &str); err != nil {
		log.Fatal(err)
	}

	writeFile, err := xml.MarshalIndent(str, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("output.xml", writeFile, 0600); err != nil {
		log.Fatal(err)
	}
	return nil

}

type XmlFile struct {
	File_name string
}

func (r XmlFile) convert() error {
	var str Cakes
	readFile, err := ioutil.ReadFile(r.File_name)
	if err != nil {
		log.Fatal(err)
	}
	if err := xml.Unmarshal(readFile, &str); err != nil {
		log.Fatal(err)
	}

	writeFile, err := json.MarshalIndent(str, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("out.json", writeFile, 0600); err != nil {
		log.Fatal(err)
	}

	return nil

}

func choice(p DBReader) {
	p.convert()
}
func main() {
	s := os.Args[1]
	if strings.HasSuffix(s, ".json") {
		reader := &JsonFile{File_name: s}
		choice(reader)

	} else if strings.HasSuffix(s, ".xml") {
		reader := &XmlFile{File_name: s}
		choice(reader)
	} else {
		err := errors.New("invalid format!")
		log.Fatal(err)
	}
}

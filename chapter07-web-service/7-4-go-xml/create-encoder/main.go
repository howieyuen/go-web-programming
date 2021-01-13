package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	ID      string   `xml:"id"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		ID:      "1",
		Content: "zzz",
		Author: Author{
			ID:   "2",
			Name: "xx",
		},
	}
	xmlFile, err := os.Create("post.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println(err)
		return
	}
}

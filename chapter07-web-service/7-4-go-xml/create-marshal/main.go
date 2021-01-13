package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("post.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

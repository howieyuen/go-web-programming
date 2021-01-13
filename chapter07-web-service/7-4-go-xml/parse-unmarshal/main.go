package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	ID       string    `xml:"id"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	ID   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	ID      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}

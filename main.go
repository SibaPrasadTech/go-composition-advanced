package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Author struct {
	Name string
	Country string
}

type Book struct {
	Name string
	PublishingYear int32
	Author
}

func (b Book) BookJSONWriter(wr io.Writer){
	WriteJSON(b,wr)
}

type Post struct {
	Title string
	Content string
	Author
} 

func (p Post) PostJSONWriter(wr io.Writer){
	WriteJSON(p,wr)
}

func WriteJSON(data interface{}, wr io.Writer) error {
	js,jerr := json.Marshal(data)
	if jerr != nil {
		return jerr
	}
	_,werr := wr.Write(js)
	if(werr) != nil {
		return werr
	}
	return nil
}

func main() {
	fmt.Println("Go Interface Use Case.... Create a common JSON Writer");
	author := Author{
		Name: "Siba",
		Country: "India",
	}
	post := Post{
		Author: author,
		Title: "Save Water",
		Content:"Water is life. Water is scarce. Save Water.",
	}
	book := Book{
		Author: author,
		Name: "Criminal Law",
		PublishingYear: 2023,
	}
	var postBytes bytes.Buffer
	post.PostJSONWriter(&postBytes)
	fmt.Println(postBytes)
	fmt.Printf("%s",&postBytes)
	var bookBytes bytes.Buffer
	book.BookJSONWriter(&bookBytes)
	fmt.Println(bookBytes)
	fmt.Printf("%s",&bookBytes)
}
package main

import (
	"fmt"
	"log"
)

type Author struct {
	Name string
	Country string
}

func (a Author) String() string {
	return fmt.Sprintf("AUTHOR... Name: %v ----- Country: %v \n",a.Name,a.Country)
}

type Book struct {
	Name string
	PublishingYear int32
	Author
}

func (b Book) String() string {
	return fmt.Sprintf("BOOK... Name: %v ----- Publishing Year: %v \n",b.Name,b.PublishingYear) + b.Author.String()
}

type Post struct {
	Title string
	Content string
	Author
}

func (p Post) String() string {
	return fmt.Sprintf("POST... Name: %v ----- Country: %v \n",p.Title,p.Content) + p.Author.String()
}

type Count int32

// Using the fmt.Stringer interface we can create a logger
func WriteLog(logObject fmt.Stringer){
	log.Print(logObject.String());
}

func main() {
	fmt.Println("Go Interface Use Case.... Create common Logger");
	author := Author{
		Name: "Siba",
		Country: "India",
	}
	book := Book{
		Author: author,
		Name: "Meaning of Life",
		PublishingYear: 2023,
	}
	post := Post{
		Author: author,
		Title: "Save Water",
		Content:"Water is life. Water is scarce. Save Water.",
	}
	WriteLog(author);
	WriteLog(book);
	WriteLog(post);
}
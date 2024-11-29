package main

import "fmt"

type Book struct {
	Title string
	Author string
	Series int
	Copies int
	PriceCents int

}

var books []Book

func AddBook(catalog []Book, book Book) []Book {
	catalog = append(catalog, book)
	return catalog

}
func main() {
	books = []Book{
		{
			Title: "The Promised Land",
			Author: "Grace Ogot",
			Series: 1,
			Copies: 1,
		},
		{
			Title: "Sunrise at Midnight",
			Author: "Ongoro Wa Munga",
			Series: 1,
			Copies: 1,
		},
	}

	books = AddBook(books, Book{Title: "For the love of Go"})
	fmt.Println(books)
}
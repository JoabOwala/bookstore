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

	toAdd := Book{Title: "Sunrise at Midnight", Author: "Ongoro Wa Munga", Series: 1, Copies:30}
	toPrint := AddBook(books, toAdd)
	fmt.Println(books)
	fmt.Println(toPrint)
}
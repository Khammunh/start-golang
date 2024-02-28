package main

import "github.com/gofiber/fiber/v2"


type Book struct {
	ID  int `json': "id"`
	Title string `json": "title"`
	Author string `json": "author"`

}
var books  []Book

func main() {
    app := fiber.New()

	books = append(books,Book{ID: 1, Title: "Mr John", Author: "Make"})
	books = append(books,Book{ID: 2, Title: "Mr Son", Author: "Jecket 3"})
    app.Get("/books",getBooks)	
	app.Get("/books/:id",getBook)

     app.Listen(":8080")
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}
func getBook(c *fiber.Ctx) error {
	bookId := c.Params("id")
	return c.SendString(bookId)
}


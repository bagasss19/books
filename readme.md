Problem Statement:
You are tasked to implement an HTTP server that serves a RESTful API for managing a collection of books.

The API should support the following operations:

GET /books: returns a list of all books in the collection.
GET /books/:id: returns a single book with the specified ID.
POST /books: adds a new book to the collection.
PUT /books/:id: updates an existing book with the specified ID.
DELETE /books/:id: deletes a book with the specified ID from the collection.
The book data should include the following fields:

id: a unique identifier for the book.
title: the title of the book.
author: the author of the book.
isbn: the ISBN number of the book.
The server should store the book data in memory or in a file.

Requirements:
The server should be implemented using Golang and the standard Golang libraries for HTTP handling.
The server should use a router package, such as Gorilla Mux, to handle HTTP requests.
The server should handle errors gracefully and return appropriate HTTP error codes and messages.
The server should include appropriate tests to validate its functionality.
The server should use an Object-Relational Mapping (ORM) package, such as GORM or XORM, to interact with the data source.
The server should validate input data, including the book title, author, and ISBN number.

Add support for filtering books by title, author, or ISBN number.
Add support for sorting books by title, author, or ISBN number.
Add support for pagination of book results.

Bonus:
Add support for handling concurrent requests using Goroutines and Channels.
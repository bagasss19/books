// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/books": {
            "get": {
                "description": "show list of Book",
                "tags": [
                    "Book"
                ],
                "summary": "Get Book List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Books Title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Books Author",
                        "name": "author",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Books ISBN",
                        "name": "isbn",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Page. Default is 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Limit. Default is 5",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Order data by field",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sorting. ASC or DESC",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "Create Book",
                "tags": [
                    "Book"
                ],
                "summary": "Create Book",
                "parameters": [
                    {
                        "description": "body payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/feature.BookPayload"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/books/{id}": {
            "get": {
                "description": "show Book by ID",
                "tags": [
                    "Book"
                ],
                "summary": "Get one Book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "book id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update Book with dynamic fields. editable field : isbn, author, title",
                "tags": [
                    "Book"
                ],
                "summary": "Update Book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "enter desired field that want to update",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.BookUpdateRequestJson"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete Book by ID",
                "tags": [
                    "Book"
                ],
                "summary": "Delete Book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Book ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "controller.BookUpdateRequestJson": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "feature.BookPayload": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "isbn": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Books API",
	Description:      "This is a collection of books API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

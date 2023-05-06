mock:
	mockgen -source=service/books/repository/repository.go -destination=service/books/repository/mock/bookmock.go

run:
	go run main.go

swagger:
	swag init
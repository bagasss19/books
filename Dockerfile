FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

# install and generate swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN apk add --update make
RUN make swagger

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/binary"]
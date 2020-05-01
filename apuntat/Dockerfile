FROM golang:1.13

WORKDIR /app

ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . .

# Build the application
RUN go build -o main .

EXPOSE 1323

CMD ["/app/main"]

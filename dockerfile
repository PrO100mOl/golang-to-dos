FROM golang:1.26.1-bookworm

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /app/exe main.go

CMD ["/app/exe"]

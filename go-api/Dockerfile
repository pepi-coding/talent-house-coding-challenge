FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o go-api .

EXPOSE 3000

CMD ["./go-api"]

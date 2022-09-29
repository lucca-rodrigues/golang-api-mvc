FROM golang:1.19-alpine

RUN apk add --no-cache git

WORKDIR /app/api-mvc

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# RUN go build -o ./out/api-mvc .

EXPOSE 3333

# CMD ["./out/api-mvc"]
CMD ["go", "run", "main.go"]
FROM golang:1.18
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main main.go

RUN adduser --disabled-password --gecos --quiet app
USER app
CMD ["./main"]
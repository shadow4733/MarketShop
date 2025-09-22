FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./

WORKDIR /app/cmd

RUN go mod tidy && go build -o /app/myapp .

CMD ["/app/myapp"]
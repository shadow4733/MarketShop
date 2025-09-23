FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./

RUN go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app
RUN swag init -g cmd/main.go -o cmd/docs

WORKDIR /app/cmd
RUN go mod tidy && go build -o /app/myapp .

CMD ["/app/myapp"]
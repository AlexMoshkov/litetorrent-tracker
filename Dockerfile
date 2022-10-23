FROM golang:1.19.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

RUN go build -o app ./cmd/main.go

CMD /wait && ./app

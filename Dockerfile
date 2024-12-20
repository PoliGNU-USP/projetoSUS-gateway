FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE ${PORT:-8080}

RUN go mod download

CMD ["go", "run", "./cmd/gateway/."]


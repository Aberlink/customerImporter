FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o customerimporter ./cmd/app

CMD ["sh", "-c", "echo 'Container started.' && sleep infinity"]

FROM golang

WORKDIR /WireguardHub

COPY . .

CMD go run ./...

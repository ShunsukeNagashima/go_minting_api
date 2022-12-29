# create a binary include in deploy container
FROM golang:1.19.4-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# --------------------------------------------------

# deploy
FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# --------------------------------------------------

# use for local development (hot reload)
FROM golang:1.19.4 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]

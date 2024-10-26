FROM golang:1.23.2 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o migration cmd/migration/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o seeder cmd/seeder/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go

FROM alpine:latest

COPY --from=build /app/migration /app/migration
COPY --from=build /app/seeder /app/seeder
COPY --from=build /app/server /app/server
COPY --from=build /app/.env /app/.env

WORKDIR /app

RUN chmod +x /app/migration /app/seeder /app/server

ENTRYPOINT ["/bin/sh", "-c", "sleep 5 && /app/migration && /app/seeder && /app/server"]
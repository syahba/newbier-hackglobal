# Stage 1: Build the Go binaries
FROM golang:1.23.2 AS build

WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the migration, seeder, and server binaries
RUN CGO_ENABLED=0 GOOS=linux go build -o migration cmd/migration/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o seeder cmd/seeder/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go

# Stage 2: Create the final container image
FROM alpine:latest

# Set working directory in the final image
WORKDIR /app

# Copy the binaries from the build stage
COPY --from=build /app/migration /app/migration
COPY --from=build /app/seeder /app/seeder
COPY --from=build /app/server /app/server
COPY --from=build /app/.env /app/.env

# Copy required data files and directories for the seeder
COPY --from=build /app/cmd/seeder/data /app/cmd/seeder/data

# Make sure the binaries are executable
RUN chmod +x /app/migration /app/seeder /app/server

# Wait for 5 seconds before starting migration, seeder, and server
ENTRYPOINT ["/bin/sh", "-c", "sleep 5 && /app/migration && /app/seeder && /app/server"]

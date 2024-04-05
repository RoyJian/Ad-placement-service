FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ad-placement-service .

FROM alpine:edge
WORKDIR /app
COPY --from=build /app/ad-placement-service .
COPY .env.release   ./.env
ENV GIN_MODE=release
ENTRYPOINT ["./ad-placement-service"]
EXPOSE 3000

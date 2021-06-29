# Devlopment
FROM golang:1.16.6-buster as development
ENV GOOS=linux \
    GO111MODULE='on'
WORKDIR /code
COPY go.mod .
RUN go mod download
CMD ["go", "run", "."]

# Migrations
FROM migrate/migrate:v4.14.1 as migrate
COPY ./migrations /migrations

# Production
FROM gcr.io/distroless/base:nonroot as production
COPY ./app /
CMD ["/app"]
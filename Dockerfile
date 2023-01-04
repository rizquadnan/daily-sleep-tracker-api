FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg
COPY config.env .

RUN go build -o /rest-server ./cmd/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY config.env .
COPY --from=build /rest-server /rest-server

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/rest-server"]



FROM golang:1.16 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg

RUN go build -o /rest-server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /rest-server /rest-server

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/rest-server"]



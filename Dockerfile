FROM golang:1.20.3 AS BUILDER

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app internal/cmd/main.go

FROM gcr.io/distroless/static-debian11

COPY --from=BUILDER /go/src/app/app /usr/local/bin/

EXPOSE 8000

USER nonroot:nonroot

CMD ["app"]

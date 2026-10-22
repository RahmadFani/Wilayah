FROM golang:1.27.1 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app ./cmd/api/main.go

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian13
COPY --from=build /go/bin/app /
CMD ["/app"]
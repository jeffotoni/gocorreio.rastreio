# tart by building the application.
# Build em gocorreio.rastreio com distroless
FROM golang:1.14.1 as builder

WORKDIR /go/src/gocorreio.rastreio

COPY gocorreio.rastreio .

ENV GO111MODULE=on

#RUN go install -v ./...
#RUN GOOS=linux go  build -ldflags="-s -w" -o gocorreio.rastreio main.go
RUN cp gocorreio.rastreio /go/bin/gocorreio.rastreio

RUN ls -lh

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/gocorreio.rastreio /
CMD ["/gocorreio.rastreio"]
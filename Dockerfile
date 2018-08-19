FROM golang:1.10

WORKDIR /
COPY demo.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o demo ./demo.go

FROM scratch
COPY --from=0 /demo /demo

ENTRYPOINT ["/demo"]

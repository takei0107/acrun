FROM golang:1.24 AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN make

FROM golang:1.24 AS runner

WORKDIR /work

COPY --from=builder /go/src/app/acrun .
CMD ["./acrun"]

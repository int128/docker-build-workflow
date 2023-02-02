FROM golang:1.20 as builder

WORKDIR /workspace
COPY go.* .
RUN go mod download

COPY main.go .
RUN go build -o main main.go

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/main .
USER 65532:65532
ENTRYPOINT ["/main"]

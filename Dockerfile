FROM golang:1.12.6 as builder
WORKDIR /go/src/app
ENV GO111MODULE on
COPY go.* ./
RUN go mod tidy
COPY . .
RUN PORT=8888 BUILD_MODE=debug ginkgo -v -r
RUN go build -o locrep

FROM alpine:3.9.3
WORKDIR /
COPY --from=builder /go/src/app/locrep /
ENTRYPOINT /locrep
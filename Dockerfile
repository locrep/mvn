FROM golang:1.13.4 AS builder
WORKDIR /go/src/app
ENV GO111MODULE on
ARG port
ARG mode
ARG redis
RUN go get github.com/onsi/ginkgo/ginkgo
COPY go.* ./
RUN go mod tidy
COPY . .
RUN REDIS_URL=$redis PORT=$port BUILD_MODE=$mode ginkgo -v -r
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o locrep .

FROM alpine:latest
WORKDIR /go/src/app/
COPY --from=builder /go/src/app/locrep .
ENTRYPOINT ./locrep
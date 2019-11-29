FROM golang:1.13.4 as builder
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
RUN go build -o locrep

FROM golang:1.13.4-alpine3.10
WORKDIR /
COPY --from=builder /go/src/app/locrep /
ENTRYPOINT /locrep
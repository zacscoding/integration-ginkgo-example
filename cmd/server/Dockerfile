FROM golang:1.13-alpine AS build

RUN mkdir -p /go/src/github.com/zacscoding/integration-ginkgo-example ~/.ssh && \
    apk add --no-cache git openssh-client make gcc libc-dev
WORKDIR /go/src/github.com/zacscoding/integration-ginkgo-example
COPY . .
RUN make build

FROM alpine:3
COPY --from=build /go/src/github.com/zacscoding/integration-ginkgo-example/integration-server /bin/integration-server
CMD /bin/integration-server
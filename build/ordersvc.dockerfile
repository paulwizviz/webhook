
# This is only for demonstration purpose it is not completely wired up
ARG GO_VER

FROM ${GO_VER}

WORKDIR /opt

RUN apt-get update && \
    apt-get install clang -y

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go build -o ./build/bin/ordersvc ./cmd/ordersvc
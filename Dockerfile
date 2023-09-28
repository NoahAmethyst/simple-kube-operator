FROM golang:1.19.7-alpine3.17 AS base

FROM base AS builder

RUN go env -w GO111MODULE=on \
  && go env -w CGO_ENABLED=0
#  && go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /opt

COPY ./ .

RUN go mod download

RUN go build

FROM alpine:3.17 AS app

WORKDIR /opt

COPY --from=builder /opt/kube-operator ./


EXPOSE 9090

ENTRYPOINT ["./kube-operator"]
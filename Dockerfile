FROM golang:1.13.5-alpine3.10

WORKDIR $GOPATH/src/pod-dependency-init-container
COPY  ./ ./

ENV GO111MODULE on

ENV GOPROXY https://goproxy.cn

RUN CGO_ENABLED=0 go build -mod=vendor -ldflags '-w -s' -o pod-dependency-init-container  main.go

FROM alpine:3.10.3

WORKDIR /app

COPY --from=0 /go/src/pod-dependency-init-container/pod-dependency-init-container ./

ENTRYPOINT ["./pod-dependency-init-container"]

FROM alpine:3.12.1

WORKDIR /app/

COPY ./bin/pod-dependency-init-container ./pod-dependency-init-container

ENTRYPOINT ["/app/pod-dependency-init-container"]

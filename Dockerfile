FROM alpine:latest
LABEL MAINTAINER="grin-ch"

WORKDIR /apps
COPY ./cfg/cfg.yaml ./cfg/cfg.yaml
COPY ./main ./main

EXPOSE 8080

ENTRYPOINT ["./main"]
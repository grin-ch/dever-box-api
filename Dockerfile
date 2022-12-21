FROM alpine:latest
LABEL MAINTAINER="grin-ch"

RUN echo 'http://mirrors.aliyun.com/alpine/v3.4/main/' > /etc/apk/repositories \
    && apk --update add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && rm -rf /var/cache/apk/*

WORKDIR /apps
COPY ./cfg/cfg.yaml ./cfg/cfg.yaml
COPY ./main ./main

EXPOSE 8080

ENTRYPOINT ["./main"]
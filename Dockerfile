FROM alpine:latest

LABEL maintainer="gwuhaolin <gwuhaolin@gmail.com>"

ENV LIGHTSOCKS_SERVER_PORT 12315
ENV LIGHTSOCKS_VERSION 1.0.8
ENV LIGHTSOCKS_DOWNLOAD_URL https://github.com/gwuhaolin/lightsocks/releases/download/${LIGHTSOCKS_VERSION}/lightsocks_${LIGHTSOCKS_VERSION}_linux_amd64.tar.gz

RUN apk upgrade --update
RUN apk add --no-cache curl tar
RUN curl -SLO ${LIGHTSOCKS_DOWNLOAD_URL}
RUN tar -zxf lightsocks_${LIGHTSOCKS_VERSION}_linux_amd64.tar.gz
RUN apk del curl tar
RUN rm -rf lightsocks-local lightsocks_${LIGHTSOCKS_VERSION}_linux_amd64.tar.gz readme.md

EXPOSE ${LIGHTSOCKS_SERVER_PORT}
CMD ./lightsocks-server
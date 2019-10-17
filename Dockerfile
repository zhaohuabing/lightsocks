FROM alpine:latest

LABEL maintainer="gwuhaolin <gwuhaolin@gmail.com>"
ENV LIGHTSOCKS_SERVER_PORT 12315
COPY ./dist/lightsocks-server_linux_amd64/lightsocks-server ./lightsocks-server
EXPOSE ${LIGHTSOCKS_SERVER_PORT}
CMD ./lightsocks-server
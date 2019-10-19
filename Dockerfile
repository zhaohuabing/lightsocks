FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o lightsocks-server ./cmd/lightsocks-server

FROM alpine:latest
LABEL maintainer="gwuhaolin <gwuhaolin@gmail.com>"
ENV LIGHTSOCKS_SERVER_PORT 12315
COPY --from=builder /app/lightsocks-server .
EXPOSE ${LIGHTSOCKS_SERVER_PORT}
CMD ./lightsocks-server
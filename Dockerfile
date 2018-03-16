FROM alpine:latest

ARG BUILD_NUM

RUN mkdir /service

WORKDIR /service

COPY $BUILD_NUM/* .

EXPOSE 8081 8081
FROM alpine:latest

ARG BUILD_NUM

RUN mkdir /service

COPY /$BUILD_NUM /service

WORKDIR /service

CMD service-b
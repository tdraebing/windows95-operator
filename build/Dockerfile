FROM alpine:3.8

RUN apk upgrade --update --no-cache

USER nobody

ADD build/_output/bin/win95-operator /usr/local/bin/win95-operator

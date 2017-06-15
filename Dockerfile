FROM golang:1.8
MAINTAINER Luc CHMIELOWSKI <luc.chmielowski@gmail.com>

# Envs
ENV GO15VENDOREXPERIMENT=1

EXPOSE 5004
RUN apt-get update -y
RUN mkdir -p /go/src/github.com/iochti/point-service
WORKDIR /go/src/github.com/iochti/point-service
COPY . /go/src/github.com/iochti/point-service

RUN go get github.com/tools/godep
RUN godep restore
RUN go install ./...

RUN rm -rf /go/src/github.com/iochti/point-service

CMD ["point-service"]

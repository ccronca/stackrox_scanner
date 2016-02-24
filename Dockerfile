FROM golang:1.5
MAINTAINER Quentin Machu <quentin.machu@coreos.com>

RUN apt-get update && \
    apt-get install -y bzr rpm xz-utils && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN mkdir /db
VOLUME /db
VOLUME /config

EXPOSE 6060 6061

ADD .   /go/src/github.com/coreos/clair/
WORKDIR /go/src/github.com/coreos/clair/

ENV GO15VENDOREXPERIMENT 1
RUN go install -v github.com/coreos/clair/cmd/clair

ENTRYPOINT ["clair"]

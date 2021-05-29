FROM golang:1.16-alpine as base

ENV GO111MODULE=on
ENV GIN_MODE=release
WORKDIR $GOPATH/src/github.com/manishiitg/sfu-coordinator

COPY go.mod go.sum ./
RUN cd $GOPATH/src/github.com/manishiitg/sfu-coordinator && go mod download

COPY . $GOPATH/src/github.com/manishiitg/sfu-coordinator
RUN GOOS=linux go build -o /sfu-coordinator .

FROM google/cloud-sdk:alpine
RUN apk --update add openjdk7-jre
COPY --from=base /sfu-coordinator /usr/local/bin/sfu-coordinator

RUN apk add --update \
 curl \
 which \
 bash

COPY ./steady-datum-291915-386b3c696fb0.json /usr/src/app/steady-datum-291915-386b3c696fb0.json

# RUN curl -sSL https://sdk.cloud.google.com | bash
# ENV PATH $PATH:/root/google-cloud-sdk/bin

RUN gcloud auth activate-service-account --key-file=/usr/src/app/steady-datum-291915-386b3c696fb0.json
RUN gcloud config list
RUN gcloud config set project steady-datum-291915

ENTRYPOINT ["/usr/local/bin/sfu-coordinator"]
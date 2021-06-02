# FROM golang:1.16 as golang
# FROM gcr.io/google.com/cloudsdktool/cloud-sdk:slim
# COPY --from=golang /usr/local/go/ /usr/local/go/
# ENV GOROOT=/usr/local/go/
# ENV PATH=$PATH:/usr/local/go/bin/
# RUN apt-get install -y make
# RUN apt-get install -y default-jre
# RUN apt-get install -y google-cloud-sdk-datastore-emulator google-cloud-sdk-app-engine-go
# RUN go version
FROM golang:1.16 as golang
# COPY --from=base /sfu-coordinator /usr/local/bin/sfu-coordinator

FROM gcr.io/google.com/cloudsdktool/cloud-sdk:slim
COPY --from=golang /usr/local/go/ /usr/local/go/

ENV GOROOT=/usr/local/go/
ENV PATH=$PATH:/usr/local/go/bin/
RUN apt-get install -y make
RUN apt-get install -y default-jre
RUN go version


ENV GO111MODULE=on
ENV GIN_MODE=release
WORKDIR $GOPATH/src/github.com/manishiitg/sfu-coordinator

COPY go.mod go.sum ./
RUN cd $GOPATH/src/github.com/manishiitg/sfu-coordinator && go mod download

COPY . $GOPATH/src/github.com/manishiitg/sfu-coordinator
RUN GOOS=linux go build -o /sfu-coordinator .


COPY ./steady-datum-291915-386b3c696fb0.json /usr/src/app/steady-datum-291915-386b3c696fb0.json

RUN gcloud auth activate-service-account --key-file=/usr/src/app/steady-datum-291915-386b3c696fb0.json
# RUN gcloud config list
RUN gcloud config set project steady-datum-291915
  
ENTRYPOINT ["/sfu-coordinator"]
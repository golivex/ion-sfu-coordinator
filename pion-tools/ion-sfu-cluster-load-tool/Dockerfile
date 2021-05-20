FROM golang:1.15-alpine as base 

ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/manishiitg/ion-sfu-cluster-load-tool

RUN apk add \
    gstreamer \
    gstreamer-dev \
    gst-plugins-base \
    gst-plugins-base-dev \
    gst-plugins-good \
    gst-plugins-bad \
    gst-plugins-ugly

FROM base as build

RUN apk add \
    build-base \
    pkgconfig

COPY go.mod go.sum ./
RUN cd $GOPATH/src/github.com/manishiitg/ion-sfu-cluster-load-tool && go mod download

COPY . $GOPATH/src/github.com/manishiitg/ion-sfu-cluster-load-tool
RUN GOOS=linux go build -o /ion-sfu-cluster-load-tool .

FROM base 
RUN apk --no-cache add ca-certificates
COPY --from=build /ion-sfu-cluster-load-tool /usr/local/bin/ion-sfu-cluster-load-tool

# COPY config.toml /configs/sfu.toml

ENTRYPOINT ["/usr/local/bin/ion-sfu-cluster-load-tool"]
CMD ["-session", "test", "-file", "test"]

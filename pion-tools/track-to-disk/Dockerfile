FROM golang:1.15-alpine as base 

ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/manishiitg/track-to-disk

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
RUN cd $GOPATH/src/github.com/manishiitg/track-to-disk && go mod download

COPY . $GOPATH/src/github.com/manishiitg/track-to-disk
RUN GOOS=linux go build -o /track-to-disk .

FROM base 
RUN apk --no-cache add ca-certificates
COPY --from=build /track-to-disk /usr/local/bin/track-to-disk

# COPY config.toml /configs/sfu.toml

ENTRYPOINT ["/usr/local/bin/track-to-disk"]
CMD ["-session", "test", "-file", "test"]

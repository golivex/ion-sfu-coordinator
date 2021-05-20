FROM golang:1.16-alpine as base 

ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/manishiitg/rtmp-to-webrtc

COPY go.mod go.sum ./
RUN cd $GOPATH/src/github.com/manishiitg/rtmp-to-webrtc && go mod download

COPY . $GOPATH/src/github.com/manishiitg/rtmp-to-webrtc
RUN GOOS=linux go build -o /rtmp-to-webrtc .

FROM base 
RUN apk --no-cache add ca-certificates
COPY --from=build /rtmp-to-webrtc /usr/local/bin/rtmp-to-webrtc

# COPY config.toml /configs/sfu.toml

ENTRYPOINT ["/usr/local/bin/rtmp-to-webrtc"]
CMD ["-session", "test", "-file", "test"]

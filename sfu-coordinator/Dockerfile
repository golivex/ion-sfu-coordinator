FROM golang:1.16-alpine as base

ENV GO111MODULE=on
ENV GIN_MODE=release
WORKDIR $GOPATH/src/github.com/manishiitg/sfu-coordinator

COPY go.mod go.sum ./
RUN cd $GOPATH/src/github.com/manishiitg/sfu-coordinator && go mod download

COPY . $GOPATH/src/github.com/manishiitg/sfu-coordinator
RUN GOOS=linux go build -o /sfu-coordinator .

FROM base 
RUN apk --no-cache add ca-certificates
COPY --from=base /sfu-coordinator /usr/local/bin/sfu-coordinator

ENTRYPOINT ["/usr/local/bin/sfu-coordinator"]
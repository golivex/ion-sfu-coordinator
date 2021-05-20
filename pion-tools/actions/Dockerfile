FROM golang:1.16-alpine as base 

ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/manishiitg/actions

COPY go.mod go.sum ./
RUN cd $GOPATH/src/github.com/manishiitg/actions && go mod download

COPY . $GOPATH/src/github.com/manishiitg/actions
RUN GOOS=linux go build -o /actions .

FROM base 
RUN apk --no-cache add ca-certificates
COPY --from=base /actions /usr/local/bin/actions

# COPY config.toml /configs/sfu.toml

ENTRYPOINT ["/usr/local/bin/actions"]
CMD [""]

# mbus
daocloud mbus implemented by rabbitmq

## Getting Started
    go get github.com/daocloud-hwu/mbus
    cd $GOPATH/src/github.com/daocloud-hwu/mbus/example
    go build receive.go
    go build emit.go
    ./receive "topic01"
    ./emit "topic01" "hello topic 01"

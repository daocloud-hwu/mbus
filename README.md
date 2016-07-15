# mbus
DaoCloud message bus sdk, written in go

## Subscribe
    func callback(body []byte) {
        fmt.Println(string(body))
    }

    client, _ := mbus.Connect("localhost:5672")
    client.Subscribe("topic01", callback)

## Publish
    client, _ := mbus.Connect("localhost:5672")
    client.Publish("topic01", []byte("message01"))

## Example
    Please make sure rabbitmq is installed
    go get github.com/daocloud-hwu/mbus
    cd $GOPATH/src/github.com/daocloud-hwu/mbus/example
    go run receive.go "topic01"
    go run emit.go "topic01" "message01"

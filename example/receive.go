package main

import (
	"fmt"
	"log"
	"os"

    "github.com/daocloud-hwu/mbus"
)

func callback(body []byte) {
    fmt.Println(string(body))
}

func main() {
    cli, err := mbus.Connect("localhost:5672")
    if err != nil {
        log.Fatalln(err)
    }

    if len(os.Args) < 2 {
        log.Printf("Usage: %s [topics]...", os.Args[0])
        os.Exit(0)
    }

    for _, s := range os.Args[1:] {
        log.Printf("Subscribe topic %s...\n", s)
        if err = cli.Subscribe(s, callback); err != nil {
            log.Fatalln(err)
        }
    }

	log.Printf(" [*] Waiting for topics. To exit press CTRL+C")

    forever := make(chan struct{})
	<-forever
}

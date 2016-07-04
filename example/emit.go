package main

import (
	"log"
	"os"
	"strings"

	"github.com/daocloud-hwu/mbus"
)

func main() {
    cli, err := mbus.Connect("localhost:5672")
    if err != nil {
        log.Fatalln(err)
    }

    body := bodyFrom(os.Args)
    if err := cli.Publish(severityFrom(os.Args), []byte(body)); err != nil {
        log.Println(err)
    } else {
    	log.Printf(" [x] Sent %s", body)
    }
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}

package main

import (
	"fmt"
	"github.com/heatxsink/go-logstash"
)

func main() {
	l := logstash.New("127.0.0.1", 9000, 5)
	_, err := l.Connect()
	if err != nil {
		fmt.Println(err)
	}
	err = l.Writeln("{ 'foo' : 'bar' }")
	if err != nil {
		fmt.Println(err)
	}
}

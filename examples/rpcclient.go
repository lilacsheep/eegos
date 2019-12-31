package main

import (
	"eegos/cluster"

	"log"
	"time"
)

func main() {
	cluster.Connect("test", "127.0.0.1:1234")

	for {
		time.Sleep(2 * time.Second)
		cluster.Send("test", "Test.PrintAB")
		log.Println("send:1: Test.PrintAB done")

		ret := cluster.Call("test", "Test.TestArgs", 10)
		log.Println("call:2: ret", ret[0])

		ret = cluster.Call("test", "Test.TestString", "aaaa")
		log.Println("call:3: ret", ret[0])
	}
}
package main

import (
	"fmt"
	"github.com/jim/demo/client"
	"github.com/jim/demo/server"
)

func main() {
	go server.StartServer()
	//time.Sleep(5 * time.Second)
	//go server.StartServer2()
	fmt.Println("00")
	for i := 0; i < 10; i++ {
		client.ClientRun()
	}
	//进入阻塞
	select {}
}

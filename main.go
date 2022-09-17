package main

import {
	server "github.com/KaratsubaLabs/onigiri-server-go/api"
}


func main() {
	go server.Start()
}

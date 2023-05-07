package main

import "github.com/bypepe77/Rajoy-Hits-API/server"

func main() {
	err := Run()
	if err != nil {
		panic("error running the server")
	}
}

func Run() error {
	config := server.NewConfig("Greatest hits", "localhost", "8080")
	server := server.NewServer(config)
	return server.Run()
}

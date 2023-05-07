package main

import (
	"errors"
	"fmt"

	"github.com/bypepe77/Rajoy-Hits-API/structs"
)

const (
	Variable = "test"
)

const Path = "url"

func main() {
	streamer := structs.NewStreamer("altas", "altaskur", 34)
	fmt.Println(streamer.GetName())
	fmt.Println(streamer.GetChannel())
	fmt.Println(streamer.GetViewers())
	streamer.SetName("test")
	fmt.Println(streamer.GetName())

	streamersList := []*structs.Streamer{
		{
			Name:    "Altaskur",
			Channel: "Altaskur",
			Viewers: 20,
		},
		{
			Name:    "David",
			Channel: "CursosDeDesarrollo",
			Viewers: 24,
		},
		{
			Name:    "Pythonesa",
			Channel: "Patonesafan",
			Viewers: 40,
		},
	}
	for index, streamer := range streamersList {
		fmt.Println(index, streamer.ToString())
	}
}

// sayhello dddd
func sayHello(name string) (string, error) {
	hello, err := sayHello2(name)
	if err != nil {
		return "", err
	}
	return hello, nil
}

func sayHello2(name string) (string, error) {
	return "", errors.New("Error")
}

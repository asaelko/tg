package main

import (
	"log"
	_ "./src/mtproto"
	"fmt"
)

var (
	configFile string = "./telegramData"
)

func main(){
	log.Println("Hello!")
	log.Println("Starting app")

	var app App
	app.LoadState(configFile)

	// connect to telegram servers

	/*m, err := mtproto.NewMTProto("./telegram_go")

	if err != nil {
		log.Fatalf("Creating of temporary file failed: %s\n", err)
	}
	*/
	/*
	err = m.Connect()
	if err != nil {
		log.Fatalf("Connect failed: %s\n", err)
	}
	log.Println("Connected")

	// check auth
	// if auth was found ...

	// no auth — lets create auth
	err = m.Auth()
	if err != nil {
		log.Fatal(err)
	}

	tgChannel, error := m.SearchChannel("nudesporno")
	if error != nil {
		log.Println(error)
	} else {
		tgChannel.GetMessages(0,0,0,0,0,0)
	}
	*/
	go app.WebserverStart()

	log.Println("Waiting for input..")

	var str string
	for {
		fmt.Scanf("%s", &str)

		go func(str string){
			GetCommand(str)
		}(str)
	}
}

func GetCommand(str string){
	log.Println(str)
}
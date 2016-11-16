package main

import (
	"log"
	"net/http"
	"fmt"
	"io"
)

func (app *App) WebserverStart(){
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/webhook/set", app.SetWebhook)

	log.Println("Web server started at localhost:8989..")
	http.ListenAndServe(":8989", nil)
}

func (app *App) SetWebhook(w http.ResponseWriter, r *http.Request) {
	// get data from request
	newHook := r.URL.Query().Get("hook")

	if newHook == "" {
		io.WriteString(w, "{\"result\": false}")
		return
	}

	app.Lock()
	oldHook := app.Webhook
	app.Webhook = newHook
	if err := app.SaveState(configFile); err != nil {
		app.Webhook = oldHook
		log.Println(err)
		io.WriteString(w, "{\"result\": false}")
		return
	}
	app.Unlock()

	io.WriteString(w, "{\"result\": true}")
	return
}


func viewHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Hello</h1>");
	log.Printf("Incoming request %q\n", r)
}

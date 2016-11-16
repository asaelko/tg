package main

import (
	"./src/mtproto"
	"bytes"
	"encoding/gob"
	"io"
	"os"
	"sync"
)

type App struct {
	mutex sync.Mutex
	Webhook  string
	Channels []mtproto.Channel
	Auths    []mtproto.Auth
}

func (app *App) LoadState(configFile string) error {
	config, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer config.Close()

	buf := make([]byte, 1024*4)
	nTotal := 0
	for {
		// read a chunk
		n, err := config.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		nTotal = nTotal + n
		if n == 0 {
			break
		}
	}

	if nTotal != 0 {
		// we have previous state, lets encode it
		app.Lock()
		err = app.decode(buf)
		app.Unlock()
		if err != nil {
			return err
		}
	} else {
		//write base structure
		app.SaveState(configFile)
	}

	return nil // success
}

func (app *App) SaveState(configFile string) error {
	config, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer config.Close()

	writeBuf, err := app.encode()
	if err != nil {
		return err
	}

	_, err = config.Write(writeBuf)
	if err != nil {
		return err
	}

	// success
	return nil
}

func (app *App) GetActiveAuth() {

}


func (app *App) Lock () {
	app.mutex.Lock()
}

func (app *App) Unlock () {
	app.mutex.Unlock()
}

func (app *App) encode() ([]byte, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(&app)
	if err != nil {
		return []byte{}, err
	}
	return b.Bytes(), nil
}

func (app *App) decode(data []byte) error {
	b := bytes.Buffer{}
	b.Write(data)
	d := gob.NewDecoder(&b)
	err := d.Decode(app)
	if err != nil {
		return err
	}

	// success
	return nil
}

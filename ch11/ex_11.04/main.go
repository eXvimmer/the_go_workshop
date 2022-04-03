package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
)

// client-side user model
type UserClient struct {
	ID   string
	Name string
}

// client-side transaction
type TxClient struct {
	ID          string
	User        *UserClient
	AccountFrom string
	AccountTo   string
	Amount      float64
}

// server-side user model
type UserServer struct {
	ID string
}

// server-side transaction
type TxServer struct {
	ID          string
	User        UserServer
	AccountFrom string
	AccountTo   string
	Amount      *float32
}

func main() {
	var net bytes.Buffer // dummy network

	clientTx := &TxClient{
		ID: "123456789",
		User: &UserClient{
			ID:   "ABCDEF",
			Name: "Mustafa",
		},
		AccountFrom: "Malena",
		AccountTo:   "Maya",
		Amount:      99.9,
	}

	// encode data using client types
	enc := gob.NewEncoder(&net)
	if err := enc.Encode(clientTx); err != nil {
		log.Fatal("error encoding: ", err)
	}

	// send data to server
	serverTx, err := sendToServer(&net)
	if err != nil {
		log.Fatal("server error: ", err)
	}

	fmt.Printf("%#v\n", serverTx)
}

func sendToServer(net io.Reader) (*TxServer, error) {
	tx := &TxServer{}

	// decode the coming data
	dec := gob.NewDecoder(net)
	err := dec.Decode(tx)

	return tx, err
}

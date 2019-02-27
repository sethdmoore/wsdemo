package main

import (
	"fmt"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func handleWs() {
	defer conn.Close()
	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			fmt.Printf("ERR: %v\n", err)
			continue
		}
		err = wsutil.WriteServerMessage(conn, op, msg)
		if err != nil {
			fmt.Printf("ERR: %v\n", err)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}

	go handleWs()
}

func main() {
	//http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}

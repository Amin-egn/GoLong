package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"io"
	"net/http"
)

func handle_server_connection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}

func server() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle_server_connection(c)
	}
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello, World!"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func hello_world(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
		  <head>
			  <title>Hello World</title>
		  </head>
		  <body>
			  Hello World!
		  </body>
		</html>`,
	)
}

// func main() {
// 	http.HandleFunc("/hello", hello_world)
// 	http.ListenAndServe(":9000", nil)
// }

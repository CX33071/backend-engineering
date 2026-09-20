package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Request struct {
	Method string
	A      int
	B      int
}

type Response struct {
	Result int
}

func Add(a, b int) int {
	return a + b
}

func handle(conn net.Conn) {
	defer conn.Close()

	var req Request

	err := json.NewDecoder(conn).Decode(&req)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	fmt.Println("收到请求:", req.Method, req.A, req.B)

	var result int

	if req.Method == "Add" {
		result = Add(req.A, req.B)
	}

	resp := Response{
		Result: result,
	}

	err = json.NewEncoder(conn).Encode(&resp)
	if err != nil {
		fmt.Println("encode error:", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("RPC Server listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}

		go handle(conn)
	}
}

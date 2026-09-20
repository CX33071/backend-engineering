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

func Call(method string, a, b int) int {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	req := Request{
		Method: method,
		A:      a,
		B:      b,
	}

	err = json.NewEncoder(conn).Encode(&req)
	if err != nil {
		panic(err)
	}

	var resp Response

	err = json.NewDecoder(conn).Decode(&resp)
	if err != nil {
		panic(err)
	}

	return resp.Result
}

func main() {
	result := Call("Add", 10, 20)

	fmt.Println("RPC result:", result)
}

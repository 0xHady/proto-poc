package main

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"

	"example.com/client/poc/client/demo" // path from the generated code
)

func main() {
	u := &demo.User{
		Name: "Ahmed",
		Age:  29,
	}

	// Serialize the message
	data, err := proto.Marshal(u)
	if err != nil {
		panic(err)
	}

	// Write it to a file (simulates sending to another service)
	err = os.WriteFile("../message.bin", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Client wrote message.bin with serialized User")
}

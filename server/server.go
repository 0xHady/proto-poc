package main

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"

	"example.com/server/poc/server/demo"
)

func main() {
	// Read the binary file created by the client
	data, err := os.ReadFile("../message.bin")
	if err != nil {
		panic(err)
	}

	var p demo.Lalaland
	if err := proto.Unmarshal(data, &p); err != nil {
		panic(err)
	}

	fmt.Printf("✅ Server successfully decoded Person: %+v\n", p)
}

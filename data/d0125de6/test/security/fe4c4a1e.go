package main

import (
	"log"
	"main/protos/query"
)

func testProtobuf() {
	password := "N@syy3ay" // $ Source

	query := &query.Query{}
	query.Description = password

	log.Println(query.GetDescription()) // $ Alert
	log.Println(query.GetId())          // OK
}

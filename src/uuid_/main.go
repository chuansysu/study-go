package main

import (
	"fmt"

	"uuid_/uuid"
)

func main() {
	id := uuid.GenUUID()
	fmt.Println(id)
}

package main

import (
	"chat/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}
}

package db

import (
	"log"

	"github.com/sdomino/scribble"
)

func Open() *scribble.Driver {
	db, err := scribble.New("./.data", nil)

	if err != nil {
		log.Fatal(err)
		return Open()
	}

	return db
}

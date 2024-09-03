package main

import (
	"context"
	"log"

	"github.com/iaPlotnikovv/New-breath/storage"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Printf("Server started...")

	db := storage.InitDB()
	defer db.Close()

}

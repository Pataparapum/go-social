package main

import (
	"log"

	"github.com/pataparapum/social/internal/env"
	"github.com/pataparapum/social/internal/env/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewStorage(nil)

	app := &application{
		cfg,
		store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}

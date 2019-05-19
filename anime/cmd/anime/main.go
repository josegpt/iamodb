package main

import (
	"log"
	"time"

	"github.com/josegpt/iamodb/anime"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURL string
}

func main() {
	cfg := Config{"root@/animes"}

	var r anime.Repository
	retry.ForeverSleep(time.Second*2, func(_ int) (err error) {
		r, err = anime.NewMySQLRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 5000...")
	s := anime.NewService(r)
	log.Fatal(anime.ListenGRPC(s, 5000))
}

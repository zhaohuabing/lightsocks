package schedule

import (
	"time"
	"log"
)

var (
	ticker = time.NewTicker(10 * time.Minute)
)

func Start() {
	log.Println("schedule", "start")
	for range ticker.C {
		CheckAllServersAlive()
	}
}

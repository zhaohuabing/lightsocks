package schedule

import (
	"time"
	"log"
)

var (
	ticker = time.NewTicker(10 * time.Minute)
)

func Start() {
	for range ticker.C {
		log.Println("schedule", "start")
		CheckAllServersAlive()
	}
}

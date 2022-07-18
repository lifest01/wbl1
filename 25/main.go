// Реализовать собственную функцию sleep.

package main

import (
	"log"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	log.Println("Do something")
	sleep(5 * time.Second)
	log.Println("Something else")
}

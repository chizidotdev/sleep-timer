package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	durationArg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Error parsing duration:", err)
	}

	duration := time.Duration(durationArg) * time.Minute

	log.Println("Automatic suspend in ", duration)

    startTime := time.Now()
    exitTime := startTime.Add(duration)
    fmt.Println("Start time: ", startTime.Format(time.RFC822), "\nExit time: ", exitTime.Format(time.RFC822))

	ticker := time.NewTicker(duration / 10)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
                timeLeft := exitTime.Sub(time.Now())
				log.Println("Automatic suspend in ", timeLeft.Round(time.Second))
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(duration)
	log.Println("Suspending system...")

	err = exec.Command("systemctl", "suspend").Run()
	if err != nil {
		log.Fatal("Error suspending system:", err)
	}

    done <- true
}

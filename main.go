package main

import (
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
	time.Sleep(duration)
	log.Println("Suspending system...")

	err = exec.Command("systemctl", "suspend").Run()
	if err != nil {
		log.Fatal("Error suspending system:", err)
	}
}

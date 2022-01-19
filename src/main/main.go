package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var znum = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "\u218A", "\u218B"}

func main() {
	poolCollection := new([]Pool)
	cin := bufio.NewReader(os.Stdin)

	running := true
	for running {
		// accept input
		fmt.Printf("> ")
		cmd, err := cin.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		cmd = strings.Replace(cmd, "\n", "", -1)

		tokens := strings.Split(cmd, " ")

		// switch on input
		switch tokens[0] {
		case "exit":
			running = false
		case "ls": // TODO: improve formatting
			for _, v := range *poolCollection {
				fmt.Println(v.name, v.running, v.ztime.Format("%d;%b%t%l"))
			}
		case "new":
			if len(tokens) >= 3 {
				newPool(poolCollection, tokens[1], tokens[2])
			}
		case "start":
			if len(tokens) >= 2 {
				temp, err := strconv.Atoi(tokens[1])
				if err == nil {
					if (temp >= 0) && (temp < len(*poolCollection)) {
						runPool(poolCollection, temp)
					} else {
						fmt.Println("Invalid Index")
					}
				}
			}
		case "pause":
			pauseAll(poolCollection)
		default:
			fmt.Println("Invalid Argument")
		}
	}
}

// runPool pauses all pools, and then starts the one selected from poolCollection[p]
func runPool(pools *[]Pool, p int) {
	cin := bufio.NewReader(os.Stdin)

	pauseAll(pools)

	(*pools)[p].SetRunning(true)
	(*pools)[p].ztime.Dec()

	// Create ticker to Dec() every lull
	poolManager := time.NewTicker(4166 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-poolManager.C:
				if (*pools)[p].ztime.Dec() == false {
					fmt.Println("Time Expired")
				}
				// TODO: Display time remaining automatically
			}
		}
	}()

	// Stop Ticker when user input is received
	_, err := cin.ReadString('\n')
	if err != nil {
		log.Fatal()
	}

	poolManager.Stop()
	done <- true

	if (*pools)[p].ztime.Dec() == true {
		(*pools)[p].ztime.Inc()
		(*pools)[p].ztime.Inc()
	}

	pauseAll(pools)
}

// pauseAll pauses all pools
func pauseAll(pools *[]Pool) {
	for i := range *pools {
		(*pools)[i].SetRunning(false)
	}
}

func newPool(pools *[]Pool, name string, time string) {
	p := NewPool(name, time)

	*pools = append(*pools, *p)
}

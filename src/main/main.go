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
			printPools(*poolCollection, []int{})
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
		case "siphon":
			if len(tokens) == 2 {
				// with 1 argument, move remaining time to overflow
				// if no overflow marked, do nothing
				source, err := strconv.Atoi(tokens[1])
				if (err == nil) && (source >= 0) && (source < len(*poolCollection)) {
					for target := range *poolCollection { // find overflow pool
						if ((*poolCollection)[target].overflow == true) &&
							(target != source) { // can't siphon overflow to overflow
							// set overflow
							(*poolCollection)[target].ztime, _ =
								Sum((*poolCollection)[target].ztime,
									(*poolCollection)[source].ztime,
								)
							// clear siphoned pool
							(*poolCollection)[source].ztime.Set("0000")
						}
					}
				} else {
					fmt.Println("Invalid Index")
				}
			} else if len(tokens) >= 3 {
				// with 2 arguments, move remaining time to other pool
				source, errs := strconv.Atoi(tokens[1])
				target, errt := strconv.Atoi(tokens[2])
				if (errs == nil) && (errt == nil) && (source != target) &&
					(source >= 0) && (source < len(*poolCollection)) &&
					(target >= 0) && (target < len(*poolCollection)) {
					// set target pool
					(*poolCollection)[target].ztime, _ =
						Sum((*poolCollection)[target].ztime,
							(*poolCollection)[source].ztime,
						)
					// clear siphoned pool
					(*poolCollection)[source].ztime.Set("0000")
				} else {
					fmt.Println("Invalid Index")
				}
			}
		case "overflow":
			// set previous overflow (if any) to false
			for i := range *poolCollection {
				(*poolCollection)[i].overflow = false
			}

			if len(tokens) >= 2 {
				// set new overflow (if valid)
				temp, err := strconv.Atoi(tokens[1])
				if err == nil {
					if (temp >= 0) && (temp < len(*poolCollection)) {
						(*poolCollection)[temp].overflow = true
					}
				}
			}
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

	// hacky `exec.Command("clear").Run()`
	for i := 0; i < 0x80; i++ {
		fmt.Println()
	}
	printPools(*pools, []int{p})

	// Create ticker to Dec() every lull
	poolManager := time.NewTicker(4166 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-poolManager.C:
				// if time has expired
				if (*pools)[p].ztime.Dec() == false {
					if (*pools)[p].overflow {
						fmt.Println("Time Expired")
					} else {
						// find overflow
						overflow := -1
						for i := range *pools {
							if (*pools)[i].overflow {
								overflow = i
							}
						}
						// run runPool on it
						if overflow != -1 {
							runPool(pools, overflow)
						} else {
							fmt.Println("Time Expired; No Overflow Found")
						}
					}
				} else {
					// display time remaining
					// hacky `exec.Command("clear").Run()`
					for i := 0; i < 0x80; i++ {
						fmt.Println()
					}
					printPools(*pools, []int{p})
				}
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

func printPools(pools []Pool, selected []int) {
	if len(selected) == 0 {
		if len(pools) >= 1 {
			// print all pools
			guideSlice := []int{}
			for i := range pools {
				guideSlice = append(guideSlice, i)
			}
			printPools(pools, guideSlice)
		}
	} else if len(selected) == 1 {
		// print with no index
		fmt.Printf("    %v  %v\n",
			pools[selected[0]].ztime.Format("%d;%b%t%l"),
			pools[selected[0]].name,
		)
	} else {
		// print selected pools with indices
		// TODO: dozenalize indices
		for i := range selected {
			info := ""
			if pools[selected[i]].running {
				info = "  [running]"
			}
			if pools[selected[i]].overflow {
				info += "  [overflow]"
			}

			fmt.Printf("[%v]  %v  %v%v\n",
				selected[i],
				pools[selected[i]].ztime.Format("%d;%b%t%l"),
				pools[selected[i]].name,
				info,
			)
		}
	}
}

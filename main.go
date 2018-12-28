package main

import (
	"fmt"
	"sync"
	"time"

	flag "github.com/ogier/pflag"
)

//flags
var (
	user      string
	numFights int
)

func main() {
	flag.Parse()
	numFights = 10

	for i := 0; i < 10; i++ {
		numFights = numFights * 10
		fmt.Println("durations for numfights =%s", numFights)
		fmt.Println("with go routines %s", runWithGo())
		fmt.Println("without go routines %s", runWithoutGo())
	}
}

func runWithGo() *time.Duration {
	var wg sync.WaitGroup
	wg.Add(numFights)
	fmt.Println("begin the brawl...")
	startTotal := time.Now()
	for i := 0; i < numFights; i++ {

		go func() {
			defer wg.Done()
			s0 := Soldier{10 * i, 1}
			s1 := Soldier{10 * i, 2}
			SimpleFight(&s0, &s1)
		}()
	}
	wg.Wait()
	elapsedTotal := time.Since(startTotal)
	return &elapsedTotal
}

func runWithoutGo() *time.Duration {
	fmt.Println("begin the brawl...")
	startTotal := time.Now()
	for i := 0; i < numFights; i++ {

		s0 := Soldier{10 * i, 1}
		s1 := Soldier{10 * i, 2}

		SimpleFight(&s0, &s1)
	}
	elapsedTotal := time.Since(startTotal)
	return &elapsedTotal
}
func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

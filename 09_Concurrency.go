package main

import (
	"fmt"
	"sync"
	"time"
)

func send_email(indx int) {
	time.Sleep(5 * time.Second)
	var my_name = "Gab"
	// Sprintf is the same as Printf but returning the formated text instead of print it
	var formattedmsg = fmt.Sprintf("Gey %v", my_name)

	fmt.Println("Call number:", indx, "Sending message to email", formattedmsg)
	wg.Done()
}

var wg = sync.WaitGroup{}

func main() {

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go send_email(i)

		fmt.Println("Moving to the next loop #######")
	}

	wg.Wait()
}

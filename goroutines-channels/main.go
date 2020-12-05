package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // crteate a channel c with type string

	for _, link := range links {
		go checkLink(link, c)
	}

	// waiting for a message from a channel is  blocking call!
	//fmt.Println(<-c) // instantly print out the message received in the channel c

	// this for loop keeps the program alive until it has logged a message for each link!
	// shortened for loop is an infinite loop (essentially, while true/alive) => for { loop logic }
	// for {
	// call the checkLink function with the message that comes through the channel and the channel c itself!
	// 	go checkLink(<-c, c) // a blocking call!
	// }

	//alternative loop syntax: (THESE ARE IDENTICAL IN EXECUTION!)
	// for l := range c { // wait for channel c to return a value and assign it to variable l
	// 	go checkLink(l, c) // instantly spawn a new goroutine with the link that came through the channel and the channel itself
	// }

	// creating an arbitrary pause between calls to the same URL
	// FUNCTION LITERALS (or, lambdas/anonymous functions)
	for l := range c {
		go func(link string) { // receive a link as a string
			time.Sleep(5 * time.Second) // sleep the coroutine for 5 secs before the new checkLink call
			checkLink(link, c)
		}(l) // pass the link from the message to the function literal to use as the link! > else it won't keep getting an updated link!
		// we have to put down these parentheses to invoke the function literal because it's a goroutine!
		// a go routine requires a FUNCTION INVOCATION!!!

		//NEVER TRY TO ACCESS THE SAME VARIABLE IN DIFFERENT ROUTINES!
		//ALWAYS PASS THEM AS ARGUMENTS!
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) //BLOCKING CALL! This freeze's the programme until this is completed
	if err != nil {
		fmt.Print(link, "might be down!")
		c <- link // send a message in the channel passed to the function
		return
	}

	fmt.Println(link, "is up!")
	c <- link // send a message in the channel passed to the function
}

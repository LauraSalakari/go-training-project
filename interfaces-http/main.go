package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// *response, error := http.Get("full url") || note the POINTER!
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// where js/ts would just return a response body inside the response return value, Go is not quite a s direct
	// here plain response represents the data around the response, not the (html)body!
	// fmt.Println(resp)

	// actually reading the response:

	// declare a []byte with a guaranteed amount of empty spaces of 99,999 (assumed sufficient length)
	// we do this because the read function is not set up to resize the slice if it is already full!
	// bs := make([]byte, 99999)

	// // use the reader interface to read the response body and pass it the empty [byte]
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))// this prints out the whole html of http://google.com

	// reading the body of the response, upgraded:
	// io.Copy(os.Stdout, resp.Body)
	// we read the incoming data in resp.Body and we write it to os.Stdout (terminal)
	// pull information from a source that implements a Reader interface
	// and passing it to a destination that implements the Writer interface

	// reading the body of the reponse with a custom Writer
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// by creating this function, the logWriter type now implements the Writer interface!
// because associating this function with the type satisfies the Writer interface
// HOWEVER! just by satisfying the interface, the type doesn't necessarily function as the interface expects it
// re: the meta of the function is sufficient to satisfy the interface even if the logic of the method is crap
func (logWriter) Write(bs []byte) (int, error) {
	// return 1, nil // this line will make the program compile without errors, but obviously it won't work the way it's expected!

	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	// the actual return is expected to be number of bytes written and an error message
	return len(bs), nil
}

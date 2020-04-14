package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

var exit1 = flag.Bool("exit1", false, "exit the script with exit code 1")
var sleep = flag.Bool("sleep", false, "sleep for an hour, basically never return unless interrupted")
var fillBuffer = flag.Bool("fill-buffer", false, "print a string with 1,000,000 characters")
var stutter = flag.Bool("stutter", false, "print 5 messages with 50 ms delays")

func main() {
	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, os.Interrupt)

	flag.Parse()

	fmt.Println("an expected string")

	if *sleep {
		/* This will listen to a ctrl-c event for up to two hours
		 * Notice: That it is *only* necessary to watch for an interrupt
		 * signal on Windows.  On Linux&MacOS the interrupt signal would
		 * always break the control flow, whereas on Windows it is not really
		 * clear when and how this is happening.
		 */
		select {
		case <-time.After(1 * time.Hour):
			fmt.Println("returning after an hour, this will never happen")
		case sig := <-c:
			fmt.Printf("received %v\n", sig)
			os.Exit(123)
		}
	}

	if *fillBuffer {
		for i := 0; i < 1e5; i++ {
			os.Stdout.Write([]byte("a"))
		}
		os.Stdout.Write([]byte("\n"))
	}

	if *stutter {
		for i := 0; i < 5; i++ {
			fmt.Printf("stuttered %d times\n", i+1)
			time.Sleep(50 * time.Millisecond)
		}
	}

	if *exit1 {
		os.Exit(1)
	}
}

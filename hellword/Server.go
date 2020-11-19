package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("%s\n", "helloword")
	for {
		time.Sleep(time.Second)
	}

	fmt.Printf("%s\n", "helloword end")
}

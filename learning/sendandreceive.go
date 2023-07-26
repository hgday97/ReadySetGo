package learning

import (
	"fmt"
	"time"
)

func SendAndReceive() {

	c1 := make(chan string) // chan is just a variable name (it's a channel in this case)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500) // Sleep for 500 ms
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2) // Sleep for 2 seconds
		}
	}()

	// Bad way
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2) // This blocks c1, so we will only receive from c1 every 2 seconds
	// }

	// Good way
	for {
		select {
		case <-c1:
			fmt.Println(<-c1)
		case <-c2:
			fmt.Println(<-c2)
		}
	}

}

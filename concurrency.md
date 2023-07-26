# Concurrency in Go 

Concurrency does not mean parallelism (exactly the same time) <br>
Concurrency instead means breaking an issue into smaller tasks that could _potentially_ be run at the same time. So a concurrent program is one that _could_ be parallelized.

To create a go routine (one that allows for concurrency) the function is called with 'go' as a prefix (eg. go count('Sheep')). This won't wait for the function to finish before going to the next line

An example of a go routine used with a wait group (to wait for the count to reach 5 for all counts then terminate) is shown in concurrency.go

For go, sending and receiving (c <- "input", output := <- c) are blockers, meaning that they have to synchronize up across the multiple go routines. This means that you can't have a corresponding send and receive in the same go routine, as the send will wait (and block) until the receive is ready, which it never will be.
package main

import (
	"fmt"
	// "net/http"
	// "github.com/gorilla/mux"
	"example.com/packages/util"
)

func main() {
	greeting := fmt.Sprintf("Hello, %s", "Harry")
	fmt.Println(greeting)

	fmt.Printf("Length of greeting: %d\n", util.StringLength(greeting))
	fmt.Println(util.GetGreeting())

	// r := mux.NewRouter()

	// http.ListenAndServe(":9000", r)
}

package main

import (
	handler "cashtransfer/dev/srvcs/server/handler"
	"fmt"
)

func main() {
	fmt.Printf("Listening on port 8080...\n")
	handler.HandleRequests()
}

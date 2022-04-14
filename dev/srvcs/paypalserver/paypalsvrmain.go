package main

import (
	handler "cashtransfer/dev/srvcs/paypalserver/handler"
	"fmt"
)

func main() {
	fmt.Printf("Listening on port 8082...\n")
	handler.HandleRequests()
}

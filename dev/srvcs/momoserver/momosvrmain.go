package main

import (
	handler "cashtransfer/dev/srvcs/momoserver/handler"
	"fmt"
)

func main() {
	fmt.Printf("Listening on port 8081...\n")
	handler.HandleRequests()
}

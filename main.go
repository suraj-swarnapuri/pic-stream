package main

import "github.com/suraj-swarnapuri/pic-stream/backend"

func main() {
	router := backend.setupServer()
	router.Run("localhost:8080")
}

package main

import (
	"log"
	"runtime"
)

func main() {
	log.Printf("Hello world from %s/%s", runtime.GOOS, runtime.GOARCH)
}

package main

import (
	"fmt"
	"github.com/Tarasusrus/url-shortener/internal/app/server"
	"os"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

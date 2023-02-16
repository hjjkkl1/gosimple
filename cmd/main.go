package main

import (
	"os"

	"github.com/valyevo/gosimple/internal"
)

func main() {
	if err := internal.NewApp().Execute(); err != nil {
		os.Exit(1)
	}
}

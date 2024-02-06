package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start")

	application := app.Generate()

	exception := application.Run(os.Args)
	if exception != nil {
		log.Fatal(exception)
	}
}

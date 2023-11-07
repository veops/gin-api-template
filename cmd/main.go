package main

import (
	"fmt"
	"os"

	app "app/cmd/apps"
)

func main() {
	command := app.NewServerCommand()

	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

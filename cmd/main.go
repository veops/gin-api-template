package main

import (
	"fmt"
	"os"

	apps "app/cmd/apps"
)

func main() {
	command := apps.NewServerCommand()
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

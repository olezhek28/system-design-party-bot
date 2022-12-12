package main

import (
	"log"

	"github.com/olezhek28/system-design-party-bot/internal/app"
)

func main() {
	a := app.New()
	err := a.Run()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

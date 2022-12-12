package main

import (
	"context"
	"log"

	"github.com/olezhek28/system-design-party-bot/internal/app"
)

func main() {
	ctx := context.Background()

	a := app.New()
	err := a.Run(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/olezhek28/system-design-party-bot/internal/app"
)

func main() {
	ctx := context.Background()

	isStgEnvPtr := flag.Bool("stg-config-enabled", false, "use staging environment")
	flag.Parse()

	var isStgEnv bool
	if isStgEnvPtr != nil {
		isStgEnv = *isStgEnvPtr
	}

	a, err := app.New(ctx, isStgEnv)
	if err != nil {
		log.Fatalln(err)
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func init() {
	//	all time now will be in UTC timezone
	time.Local = time.UTC
}

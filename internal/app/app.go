package app

import (
	"context"
	"log"
	"sync"
)

type App struct {
	serviceProvider *serviceProvider
}

func New() *App {
	return &App{
		serviceProvider: NewServiceProvider(),
	}
}

func (a *App) Run(ctx context.Context) error {
	processorService := a.serviceProvider.GetProcessorService(ctx)
	schedulerService := a.serviceProvider.GetSchedulerService(ctx)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := processorService.Run(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		defer wg.Done()
		err := schedulerService.Run(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	wg.Wait()

	return nil
}

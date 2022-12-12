package app

import (
	"context"
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

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		processorService.Run(ctx)
	}()

	wg.Wait()

	return nil
}

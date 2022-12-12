package app

import "sync"

type App struct {
	serviceProvider *serviceProvider
}

func New() *App {
	return &App{
		serviceProvider: NewServiceProvider(),
	}
}

func (a *App) Run() error {
	processorService := a.serviceProvider.GetProcessorService()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		processorService.Run()
	}()

	wg.Wait()

	return nil
}

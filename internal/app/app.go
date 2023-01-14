package app

import (
	"context"
	"log"
	"sync"

	"github.com/olezhek28/system-design-party-bot/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
}

func New(ctx context.Context, isStgEnv bool) (*App, error) {
	a := &App{
		serviceProvider: NewServiceProvider(isStgEnv),
	}

	err := a.initDeps(ctx)

	return a, err
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		config.Init,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
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

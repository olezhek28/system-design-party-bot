package scheduler

import (
	"context"
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

const schedulerPeriod = "0 */10 * * * *"

func (s *Service) Run(ctx context.Context) error {
	parser := cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	scheduler := cron.New(cron.WithParser(parser))

	_, err := scheduler.AddFunc(schedulerPeriod, func() {
		fmt.Printf("cron: start 'notification' task\n")

		errJob := s.sendNotification(ctx)
		if errJob != nil {
			fmt.Printf("cron: 'notification' task failed: %s\n", errJob.Error())
		}

		fmt.Printf("cron: end 'notification' task\n")

	})
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		scheduler.Run()
	}()

	wg.Wait()
	return nil
}

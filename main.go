package main

import (
	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron/v2"
	"time"
)

const timeBetweenNotifications = 20 * time.Minute
const endTime = 48 * time.Hour

func main() {

	schedule, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	_, err = schedule.NewJob(
		gocron.DurationJob(
			timeBetweenNotifications,
		),
		gocron.NewTask(
			func() {
				notify()
			},
		),
	)

	if err != nil {
		panic(err)
	}

	schedule.Start()

	select {
	case <-time.After(endTime):
	}
}

func notify() {
	err := beeep.Notify("Pause Reminder!", "Take a break and make tea!", "")
	if err != nil {
		panic(err)
	}
}

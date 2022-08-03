package main

import (
	"github.com/nitrictech/go-sdk/faas"
	"github.com/nitrictech/go-sdk/resources"
	"github.com/sirupsen/logrus"
)

func run() error {
	err := resources.NewSchedule("hourly-job", "1 hours",
		func(ec *faas.EventContext, eh faas.EventHandler) (*faas.EventContext, error) {
			log := logrus.WithFields(logrus.Fields{
				"schedule": "hourly-job",
			})

			log.Info("running")

			return eh(ec)
		})
	if err != nil {
		return err
	}

	return resources.Run()
}

func main() {
	if err := run(); err != nil {
		logrus.Info("Shutting down with error:", err.Error())
	}

}

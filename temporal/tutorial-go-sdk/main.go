package main

import (
	"go.uber.org/zap"

	"go.temporal.io/temporal/client"
	"go.temporal.io/temporal/worker"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("Zap logger created")

	// The client is a heavyweight object that should be created once
	serviceClient, err := client.NewClient(client.Options{
		Logger: logger,
	})

	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}

	worker := worker.New(serviceClient, "tutorial_tq", worker.Options{})

	worker.RegisterWorkflow(Greetings)
	worker.RegisterActivity(GetUser)
	worker.RegisterActivity(SendGreeting)

	err = worker.Start()
	if err != nil {
		logger.Fatal("Unable to start worker", zap.Error(err))
	}

	select {}
}

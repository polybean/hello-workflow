package util

import (
	"context"

	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"
)

// StartWorker starts workflow worker and activity worker based on configured options.
func StartWorker(env *WorkerServiceEnv, wo worker.Options) {
	worker := worker.New(env.ClientStub, env.Domain, env.TaskList, wo)
	err := worker.Start()
	if err != nil {
		env.Logger.Error("Failed to start workers.", zap.Error(err))
		panic("Failed to start workers")
	}
}

// StartWorkflow starts a workflow
func StartWorkflow(env *WorkerServiceEnv, options client.StartWorkflowOptions, workflow interface{}, args ...interface{}) {
	cc, err := env.Builder.BuildCadenceClient()
	if err != nil {
		env.Logger.Error("Failed to build cadence client.", zap.Error(err))
		panic(err)
	}

	we, err := cc.StartWorkflow(context.Background(), options, workflow, args...)
	if err != nil {
		env.Logger.Error("Failed to create workflow", zap.Error(err))
		panic("Failed to create workflow.")

	} else {
		env.Logger.Info("Started Workflow", zap.String("WorkflowID", we.ID), zap.String("RunID", we.RunID))
	}
}

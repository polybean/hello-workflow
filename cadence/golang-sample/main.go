package main

import (
	"flag"
	"time"

	"polybean/cadence-sample/util"

	"github.com/google/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
)

type flags struct {
	mode     string
	flow     string
	domain   string
	taskList string
}

func setFlags(f *flags) {
	flag.StringVar(&f.mode, "m", "trigger", "Mode: worker or trigger")
	flag.StringVar(&f.flow, "f", "", "Flow definition YAML file")
	flag.StringVar(&f.domain, "d", "sample", "Domain")
	flag.StringVar(&f.taskList, "tl", "sample-queue", "Task list")
	flag.Parse()
}

func startWorker(env *util.WorkerServiceEnv) {
	wo := worker.Options{
		MetricsScope: env.Scope,
		Logger:       env.Logger,
	}
	util.StartWorker(env, wo)
}

func startTrigger(env *util.WorkerServiceEnv) {
	wo := client.StartWorkflowOptions{
		ID:                              "polybean_" + uuid.New().String(),
		TaskList:                        env.TaskList,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	util.StartWorkflow(env, wo, SampleWorkflow, "polybean")
}

func main() {
	var f flags
	setFlags(&f)

	var env util.WorkerServiceEnv
	env.SetupWorkerServiceEnv("cadence.yaml", f.domain, f.taskList)

	switch f.mode {
	case "worker":
		startWorker(&env)
		// The workers are supposed to be long running process that should not exit.
		// Use select{} to block indefinitely for samples, you can quit by CMD+C.
		select {}
	case "trigger":
		startTrigger(&env)
	}
}

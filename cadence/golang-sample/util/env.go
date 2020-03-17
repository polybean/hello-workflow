package util

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/uber-go/tally"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

type (
	// WorkerServiceEnv holds the environment details for a worker service
	WorkerServiceEnv struct {
		ClientStub           workflowserviceclient.Interface
		CadenceServiceConfig CadenceServiceConfig
		Logger               *zap.Logger
		Scope                tally.Scope
		Builder              *WorkflowClientBuilder
		Domain               string
		TaskList             string
	}

	// CadenceServiceConfig holds the detailed information to connect to the Cadence service
	CadenceServiceConfig struct {
		DomainName      string `yaml:"domain"`
		ServiceName     string `yaml:"service"`
		HostNameAndPort string `yaml:"host"`
	}
)

func (env *WorkerServiceEnv) getCadenceServiceConfig(configFile string) {
	data, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf("Failed to log config file: %v, Error: %v", configFile, err))
	}

	if err := yaml.Unmarshal(data, &env.CadenceServiceConfig); err != nil {
		panic(fmt.Sprintf("Error initializing configuration: %v", err))
	}
}

func (env *WorkerServiceEnv) initializeLogger() {
	logger, err := zap.NewDevelopment()

	if err != nil {
		panic(err)
	}

	env.Logger = logger
}

func (env *WorkerServiceEnv) buildClientStub() {
	builder := NewClientBuilder(env.Logger).
		SetHostPort(env.CadenceServiceConfig.HostNameAndPort).
		SetDomain(env.CadenceServiceConfig.DomainName).
		SetMetricsScope(tally.NoopScope).
		SetDataConverter(nil).
		SetContextPropagators(nil)

	stub, err := builder.BuildServiceClient()

	if err != nil {
		panic(err)
	}

	env.ClientStub = stub
	env.Builder = builder
}

func (env *WorkerServiceEnv) verifyOrCreateDomain(domain string) {
	dc, err := env.Builder.BuildCadenceDomainClient()

	if err != nil {
		panic(err)
	}

	_, err = dc.Describe(context.Background(), domain)

	if err == nil {
		return
	}

	switch err := err.(type) {
	case *shared.EntityNotExistsError:
		err2 := dc.Register(context.Background(), &shared.RegisterDomainRequest{Name: &domain})

		if err2 != nil {
			panic(err2)
		}
	default:
		panic(err)
	}
}

// SetupWorkerServiceEnv creates details needed to launch a worker service
func (env *WorkerServiceEnv) SetupWorkerServiceEnv(configFile string, domain, taskList string) {
	if env.ClientStub != nil {
		return
	}

	env.Scope = tally.NoopScope
	env.Domain = domain
	env.TaskList = taskList

	env.getCadenceServiceConfig(configFile)
	env.initializeLogger()
	env.buildClientStub()
	env.verifyOrCreateDomain(domain)
}

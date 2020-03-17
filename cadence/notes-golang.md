## 基本概念

Domain：类型名字空间
TaskList：队列/Topic
LocalActivity：不需要通过 Cadence Service 分发，即可在本节点运行 Activity

## 关于注册

示例工程：`helloworld`

- 跟踪源代码，该注册是向 Worker Service 的运行内存中注册；而非向 Cadence Service 注册
- 注册的要素是函数名和函数的参数

```go
func init() {
  workflow.Register(Workflow)
  activity.Register(helloworldActivity)
}
```

## 关于 Queue（TaskList）的匹配

示例工程：`dsl`

- Worker 和 trigger 是两个进程
- 分别都适用了`ApplicationName`作为 TaskList，从而完成了发送-接收的关联

```go
func startWorkers(h *common.SampleHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, ApplicationName, workerOptions)
}

func startWorkflow(h *common.SampleHelper, w Workflow) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "dsl_" + uuid.New(),
		TaskList:                        ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	h.StartWorkflow(workflowOptions, SimpleDSLWorkflow, w)
}
```

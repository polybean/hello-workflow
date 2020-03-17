## 1. Launch Cadence service

```sh
docker-compose up
```

## 2. Register domain

### 2.1 Using Shell

Register:

```sh
docker run \
  --network=host \
  --rm ubercadence/cli:master \
  --do sample \
  domain register \
  -rd 1
```

Verify:

```sh
docker run \
  --network=host \
  --rm ubercadence/cli:master \
  --do sample \
  domain describe
```

### 2.2 Using API

```sh
./gradlew -q execute -PmainClass=com.uber.cadence.samples.common.RegisterDomain
```

## 3. Define workflow & activity (`HelloActivity`)

```java
public class HelloActivity {
  static final String TASK_LIST = "HelloActivity";

  public interface GreetingWorkflow {
    /** @return greeting string */
    @WorkflowMethod(executionStartToCloseTimeoutSeconds = 10, taskList = TASK_LIST)
    String getGreeting(String name);
  }

  public interface GreetingActivities {
    @ActivityMethod(scheduleToCloseTimeoutSeconds = 2)
    String composeGreeting(String greeting, String name);
  }

  public static class GreetingWorkflowImpl implements GreetingWorkflow {
    // ...
  }

  static class GreetingActivitiesImpl implements GreetingActivities {
    // ...
  }
}
```

## 4. Register workflow & activity to worker

```java
public class HelloActivity {
  // ...
  public static void main(String[] args) {
    // Start a worker that hosts both workflow and activity implementations.
    Worker.Factory factory = new Worker.Factory(DOMAIN);
    Worker worker = factory.newWorker(TASK_LIST);
    // Workflows are stateful. So you need a type to create instances.
    worker.registerWorkflowImplementationTypes(GreetingWorkflowImpl.class);
    // Activities are stateless and thread safe. So a shared instance is used.
    worker.registerActivitiesImplementations(new GreetingActivitiesImpl());
    // Start listening to the workflow and activity task lists.
    factory.start();
  }
}
```

## 5. Start the worker

```sh
./gradlew -q execute -PmainClass=com.uber.cadence.samples.hello.HelloActivity
```

## 6. Start a workflow instance using CLI

### 6.1 Fire & forget

```sh
docker run \
  --network=host \
  --rm ubercadence/cli:master \
  --do sample \
  workflow start \
  --tasklist HelloActivity \
  --workflow_type GreetingWorkflow::getGreeting \
  --execution_timeout 3600 \
  --input \"Song\"
```

### 6.2 Wait for the result

```sh
docker run \
  --network=host \
  --rm ubercadence/cli:master \
  --do sample \
  workflow run \
  --tasklist HelloActivity \
  --workflow_type GreetingWorkflow::getGreeting \
  --execution_timeout 3600 \
  --input \"Nicole\"
```

## 7. Start a workflow instance using API

```java
public class HelloActivity {
  public static void main(String[] args) {
    // ... workflow & activity registration

    // Start a workflow execution. Usually this is done from another program.
    WorkflowClient workflowClient = WorkflowClient.newInstance(DOMAIN);
    // Get a workflow stub using the same task list the worker uses.
    GreetingWorkflow workflow = workflowClient.newWorkflowStub(GreetingWorkflow.class);
    // Execute a workflow waiting for it to complete.
    String greeting = workflow.getGreeting("World");
    System.out.println(greeting);
    System.exit(0);
  }
}
```

## 8. List workflows

```sh
docker run \
  --network=host \
  --rm ubercadence/cli:master \
  --do sample \
  workflow list
```

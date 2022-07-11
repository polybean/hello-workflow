## 0. Installation

```sh
# @H1
cd $HOME/repo/polybean/hello-workflow/temporal
curl -L https://github.com/temporalio/temporal/releases/latest/download/docker.tar.gz | tar -xz --strip-components 1 docker/docker-compose.yml
ls

docker-compose up
```

```sh
open http://localhost:8088/
```

## 1. Quick Example

```sh
# @H2
cd tutorial-go-sdk
go build

# Start the worker process
./tutorial-go-sdk
```

```sh
# @H3
# Start workflow execution
docker run --network=host --rm temporalio/tctl:latest wf start --tl tutorial_tq -w Greet_Temporal_1 --wt Greetings --et 3600 --dt 10
docker run --network=host --rm temporalio/tctl:latest wf start --help | grep "\-\-tl"
docker run --network=host --rm temporalio/tctl:latest wf start --help | grep "\-\-w"
docker run --network=host --rm temporalio/tctl:latest wf start --help | grep "\-\-wt"
docker run --network=host --rm temporalio/tctl:latest wf start --help | grep "\-\-dt"

# List namespaces
docker run --network=host --rm temporalio/tctl:latest namespace list
docker run --network=host --rm temporalio/tctl:latest n register default #--help #c default

```

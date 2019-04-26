{%- if cookiecutter.use_dynamodb == "y" -%}
DYNAMODB_CONTAINER_NAME := dynamodb-local

all: run

setup:
	./build/setup.sh $(DYNAMODB_CONTAINER_NAME) $(GOAWS_CONTAINER_NAME)

clean:
	. build/clean.sh $(DYNAMODB_CONTAINER_NAME)

run: setup gobuild
	sam local start-api -n configs/local.json -t deployments/template.yml

gobuild:
	GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o deployments/creation cmd/creation/main.go
{%- else -%}
all: run

gobuild:
	GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o deployments/creation cmd/creation/main.go

run: gobuild
	sam local start-api -n configs/local.json -t deployments/template.yml
{%- endif %}

test:
	go test ./... -cover

lint:
	golangci-lint run
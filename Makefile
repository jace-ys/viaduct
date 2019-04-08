TARGET = main
SOURCE = main.go
IMAGE = jaceys/viaduct
FILEPATH = config/config.sample.yml
DOCKERFILE = build/Dockerfile
COMPOSEFILE = build/docker-compose.yml

.PHONY: all test build execute docker container run format clean

all: test build execute

test:
	@echo "==> Running tests.."
	go test ./... -v

# Target to build and run executable locally
build:
	@echo "==> Building from source.."
	go build -o ${TARGET} ${SOURCE}

execute:
	@echo "==> Running executable.."
	FILE_PATH=${FILEPATH} ./${TARGET}

# Target to build and run Docker image
docker: container run

container:
	@echo "==> Building image.."
	docker build -f ${DOCKERFILE} -t ${IMAGE} .

run:
	@echo "==> Running container.."
	docker run --rm -p 8000:80 -v $(shell pwd)/${FILEPATH}:/config/config.yml ${IMAGE}

# Target to build and run using Docker Compose
compose:
	@echo "==> Starting Docker Compose.."
	docker-compose -f ${COMPOSEFILE} up

format:
	@echo "==> Formatting code.."
	gofmt -w .

clean:
	@echo "==> Cleaning up.."
	go mod tidy
	go clean

TARGET = viaduct
SOURCE = main.go
IMAGE = jaceys/viaduct
VERSION = latest
CONFIGFILE = config/config.sample.yaml
DOCKERFILE = build/Dockerfile
COMPOSEFILE = build/docker-compose.yaml

.PHONY: all test build run docker image container format clean

all: format build run

test:
	@echo "==> Running tests.."
	go test ./... -v

# Target to build and run executable locally
build:
	@echo "==> Building from source.."
	go build -o ${TARGET} ${SOURCE}

run:
	@echo "==> Running executable.."
	./viaduct start -p 3000 -f config/config.sample.yaml

# Target to build and run Docker image
docker: image container

image:
	@echo "==> Building image.."
	docker build -f ${DOCKERFILE} -t ${IMAGE}:${VERSION} .

container:
	@echo "==> Running container.."
	docker run --rm -p 8000:80 -v $(shell pwd)/${CONFIGFILE}:/config/config.yaml ${IMAGE}:${VERSION}

# Target to build and run using Docker Compose
compose:
	@echo "==> Starting Docker Compose.."
	docker-compose -f ${COMPOSEFILE} up

deploy:
	@echo "==> Publishing image to Docker Hub"
	docker push ${IMAGE}:${VERSION}

format:
	@echo "==> Formatting code.."
	gofmt -w .

clean:
	@echo "==> Cleaning up.."
	go mod tidy
	go clean

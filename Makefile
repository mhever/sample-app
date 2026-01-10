APP=sample-app
IMAGE=sample-app:local

.PHONY: build docker-build run

build:
	go build -o bin/$(APP) ./

docker-build:
	docker build -t $(IMAGE) .

run:
	SERVICE_NAME=sample-app VERSION=0.1.0 DB_DSN="postgres://postgres:password@localhost:5432/postgres?sslmode=disable" \
		./bin/$(APP)

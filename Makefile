.PHONY: docs

build-local:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

run-local: build-local
	docker-compose up --remove-orphans --build server

build-deploy:
	docker image build -t editt:0.1 -f ./deploy/Dockerfile .

run-container:
	export HOST=prod
	./run-container.sh

run-deploy: build-deploy run-container

docs:
	swag init -g pkg/server/app.go
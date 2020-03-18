build-local:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

run-local: build-local
	docker-compose up --remove-orphans --build server

build-deploy:
	docker image build -t editt:0.1 -f ./deploy/Dockerfile .

run-deploy:
	export HOST=prod && docker run -e HOST --rm --publish 8000:8000 --name editt-api editt:0.1

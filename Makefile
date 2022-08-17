build:
	go build -o ./.bin/apod ./cmd/main.go

runs: build
	./.bin/apod

build-image:
	docker build -t service_authorization:v1 .

start-container:
	docker run --name service-apod-api -p 8080:8080 --env-file .env apod_api:v1

swag-generate:
	swag init -g cmd/main.go

run:
	go run cmd/main.go
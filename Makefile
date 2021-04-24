

build:
	CGO_ENABLED=0 GOOS=linux go build -v ./cmd/app/main.go 
	sudo docker-compose build
	sudo docker-compose up 
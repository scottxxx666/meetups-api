build:
	go generate ./schema
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/main main.go

VERSION= $(shell git rev-parse --short HEAD)
NOW= $(shell date +'%Y-%m-%d_%T')
compile:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags "-X main.sha1ver=$(VERSION) -X main.buildTime=$(NOW)" -o goScheduler main.go

run:
	go run main.go
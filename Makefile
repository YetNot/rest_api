.PHONY: build
build:
		go build -o main cmd/main/app.go

.PHONY: run
run:
		go run cmd/main/app.go
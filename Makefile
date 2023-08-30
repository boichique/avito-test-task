service-up:
	docker compose up -d --build

service-down:
	docker compose down

before-push:
	go mod tidy &&\
	gofumpt -l -w . &&\
	go build ./...&&\
	golangci-lint run ./... &&\
	go test -v ./tests/...

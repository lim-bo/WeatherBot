@Phony: run
run:
	go run cmd/bot/main.go

@Phony: test
test:
	go test ./internal/weather
	
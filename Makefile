@Phony: runService
runService:
	bin/weather

@Phony: runBot
runBot:
	bin/bot

@Phony: test
test:
	go test ./internal/weather
	
@Phony: build
build:
	go build -o bin/bot cmd/bot/main.go
	go build -o bin/weather cmd/weatherApi/main.go
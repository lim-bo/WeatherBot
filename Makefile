@Phony: runService
runService:
	bin/weather.exe

@Phony: runBot
runBot:
	bin/bot.exe

@Phony: test
test:
	go test ./internal/weather
	
@Phony: build
build:
	go build -o bin/bot.exe cmd/bot/main.go
	go build -o bin/weather.exe cmd/weatherApi/main.go
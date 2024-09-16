FROM golang:1.22.2-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /main cmd/weatherApi/main.go

FROM alpine:latest
COPY --from=build /main /bin/main
COPY --from=build /app/configs/secret.yaml /configs/
COPY --from=build /app/certs /certs

## Wait for other service
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
RUN chmod +x /wait

CMD /wait && /bin/main
FROM golang:1.22.2-alpine as build
WORKDIR /app
COPY  go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /main cmd/bot/main.go

FROM alpine:latest
COPY --from=build /main /bin/main
COPY --from=build /app/configs/secret.yaml /configs/
ENTRYPOINT [ "/bin/main" ]
FROM --platform=linux/amd64 golang:1.22.4-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o main ./cmd

FROM --platform=linux/amd64 alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Копируем собранное приложение
COPY --from=builder /app/main .

COPY --from=builder /app/config .

# Добавляем права на выполнение
RUN chmod +x /app/main

EXPOSE 44044

# Устанавливаем точку входа
ENTRYPOINT ["./main"]

# Устанавливаем команду по умолчанию
CMD ["--config", "./local.yaml"]
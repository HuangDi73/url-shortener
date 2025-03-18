# Переменные
BINARY_NAME=shortener

# Сборка проекта
build:
	go build -o $(BINARY_NAME) ./cmd/url-shortener

# Запуск проекта
run: build
	./$(BINARY_NAME)

# Установка зависимостей
deps:
	go mod tidy

# Очистка временных файлов
clean:
	rm -f $(BINARY_NAME)

.PHONY: build run deps clean

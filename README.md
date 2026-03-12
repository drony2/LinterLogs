# loglint

Линтер для проверки лог-сообщений в Go (совместим с `golangci-lint`).

Поддерживаемые логгеры:
- `log/slog`
- `go.uber.org/zap`

## Правила

Линтер проверяет, что лог-сообщения:
1. начинаются со строчной буквы;
2. не содержат кириллицу (только английский);
3. не содержат спецсимволы или эмодзи (разрешены только буквы/цифры/пробелы);
4. не содержат потенциально чувствительные данные по ключевым словам: `password`, `token`, `secret`, `api_key`, `apikey`.

Примечание: анализируются только статические строки (строковые литералы) и простые конкатенации строк через `+`.

## Требования

- Go 1.22+

## Запуск (standalone)

```bash
go run ./cmd/loglint ./...
```

или

```bash
go build -o loglint ./cmd/loglint
./loglint ./...
```

## Интеграция с golangci-lint (Module Plugin System, рекомендуется, работает на Windows)

1. Установите `golangci-lint` v2.
2. Создайте `.custom-gcl.yml` рядом с этим репозиторием:

```yaml
version: v2.9.0 # или ваша версия golangci-lint v2
plugins:
  - module: 'LinterForLogs'
    import: 'LinterForLogs/golangci'
    path: .
```

3. Соберите кастомный бинарь:

```bash
golangci-lint custom
```

4. В проекте, который нужно проверять, добавьте `.golangci.yml`:

```yaml
version: "2"

linters:
  default: none
  enable:
    - loglint
  settings:
    custom:
      loglint:
        type: "module"
```

5. Запустите кастомный бинарь `golangci-lint`:

```bash
./custom-gcl run ./...
```

## Интеграция с golangci-lint (Go Plugin System, только Linux/macOS)

Go-плагины (`-buildmode=plugin`) не поддерживаются на Windows.

1. Соберите `.so`:

```bash
CGO_ENABLED=1 go build -buildmode=plugin -o loglint.so ./plugin
```

2. Подключите в `.golangci.yml`:

```yaml
version: "2"

linters:
  default: none
  enable:
    - loglint
  settings:
    custom:
      loglint:
        path: /path/to/loglint.so
        description: Linter for slog/zap log messages
```

## Тесты

```bash
go test ./...
```

## Пример

```go
package main

import "log/slog"

func main() {
	slog.Info("Starting server")
	slog.Info("запуск сервера")
	slog.Info("password is 123")
	slog.Info("server started!!!")
}
```

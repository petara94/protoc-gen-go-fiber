# protoc-gen-go-fiber

**protoc-gen-go-fiber** — это плагин для `protoc` или `buf`, который автоматически генерирует маршруты для [Fiber](https://github.com/gofiber/fiber) на основе gRPC-сервисов и аннотаций `google.api.http`.
protoc-gen-go-fiber
## Особенности
- Генерация HTTP-обработчиков для Fiber по gRPC-сервисам.
- Поддержка методов: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`.
- Автоматическая валидация запросов через `protoc-gen-validate`.
- Поддержка interceptor'ов для обработки ошибок и дополнительной логики.
- Автоматическая конвертация заголовков HTTP в gRPC metadata.
- **Важно** принимает только body

## Установка

```bash
go install github.com/petara94/protoc-gen-go-fiber@latest
```
или для go 1.24+
```bash
go get -tool github.com/petara94/protoc-gen-go-fiber@latest
```

## Использование

1. Убедись, что у тебя есть аннотации `google.api.http` в `.proto`-файлах:

```proto
import "google/api/annotations.proto";

service ExampleService {
  rpc GetExample (GetExampleRequest) returns (GetExampleResponse) {
    option (google.api.http) = {
      get: "/v1/example"
      body: "*"
    };
  }
}
```

2. Сгенерируй код:

`buf.gen.yaml`
```yaml
version: v2
plugins:
  - local: [ "go",  "tool", "protoc-gen-go-fiber" ]
    out: out_dir
    opt: paths=source_relative
```

3. В приложении Fiber зарегистрируй маршруты:

```go
app := fiber.New()
RegisterExampleServiceFiberRoutes(app, serverImpl, grpcInterceptor)
```

## Требования

- Go 1.20+
- Плагины:
    - `protoc-gen-validate`
    - `google/api/annotations.proto`
- Зависимости:
    - [Fiber](https://github.com/gofiber/fiber)
    - [gRPC](https://github.com/grpc/grpc-go)
    - [protoc-gen-validate](https://github.com/envoyproxy/protoc-gen-validate)

## Структура сгенерированного кода

- Обработчик каждого метода:
    - Парсит тело запроса (если требуется).
    - Валидирует входные данные (`Validate()`).
    - Пробрасывает HTTP-заголовки в контекст.
    - Вызывает метод gRPC-сервера через переданный interceptor.
    - Возвращает JSON-ответ или ошибку.

## Лицензия

[MIT License](LICENSE)

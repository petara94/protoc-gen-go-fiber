# protoc-gen-fiber

**protoc-gen-fiber** — это плагин для `protoc`, который автоматически генерирует маршруты для [Fiber](https://github.com/gofiber/fiber) на основе gRPC-сервисов и аннотаций `google.api.http`.

## Особенности
- Генерация HTTP-обработчиков для Fiber по gRPC-сервисам.
- Поддержка методов: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`.
- Автоматическая валидация запросов через `protoc-gen-validate`.
- Поддержка interceptor'ов для обработки ошибок и дополнительной логики.
- Автоматическая конвертация заголовков HTTP в gRPC metadata.

## Установка

```bash
go install path/to/protoc-gen-fiber
```

(Замени `path/to` на актуальный путь к твоему плагину.)

## Использование

1. Убедись, что у тебя есть аннотации `google.api.http` в `.proto`-файлах:

```proto
import "google/api/annotations.proto";

service ExampleService {
  rpc GetExample (GetExampleRequest) returns (GetExampleResponse) {
    option (google.api.http) = {
      get: "/v1/example/{id}"
    };
  }
}
```

2. Сгенерируй код:

```bash
protoc \
  --proto_path=. \
  --proto_path=third_party \
  --fiber_out=. \
  your_service.proto
```

(Добавь нужные опции поиска, в зависимости от расположения `google/api/annotations.proto`.)

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

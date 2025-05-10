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

Пример с `buf`

1. Убедись, что у тебя есть аннотации `google.api.http` в `.proto`-файлах:

```protobuf
import "google/api/annotations.proto";
import "validate/validate.proto";

option go_package = "gen/go/greeterpb;greeterpb";

service GreeterService {
  rpc SayHello(HelloRequest) returns (HelloResponse) {
    option (google.api.http) = {
      post: "api/v1/hello"
      body: "*"
    };
  }
}

message HelloRequest {
  string name = 1 [(validate.rules).string.min_len = 3];
  string email = 2 [(validate.rules).string.email = true];
}
```

2. Сгенерируй код:

`buf.gen.yaml`
```yaml
version: v2
plugins:
  - local: [ "go",  "tool", "protoc-gen-go-fiber" ]
    out: out_dir
    opt:
      - paths=source_relative
      - error_handlers_package=github.com/petara94/protoc-gen-go-fiber/utils
      - json_unmarshal_package=encoding/json
      - grpc_error_handle_func=HandleGRPCStatusError
      - unmarshal_error_handle_func=HandleUnmarshalError
      - validation_error_handle_func=HandleValidationError
      - disable_grpc_interceptor=false
```

3. В приложении Fiber зарегистрируй маршруты:

```go
app := fiber.New()
RegisterGreeterServiceFiberRoutes(app, serverImpl, grpcInterceptor)
```
## Поддерживаемые флаги

| Флаг                           | Описание                                             | Значение по умолчанию          |
| ------------------------------ | ---------------------------------------------------- | ------------------------------ |
| `error_handlers_package`       | Путь к пакету с функциями обработки ошибок           | `github.com/your-org/errors`   |
| `json_unmarshal_package`       | Путь к пакету с функциями JSON-маршалинга            | `github.com/your-org/jsonutil` |
| `grpc_error_handle_func`       | Имя функции для обработки gRPC-ошибок                | `HandleGRPCStatusError`        |
| `unmarshal_error_handle_func`  | Имя функции для обработки ошибок десериализации JSON | `HandleUnmarshalError`         |
| `validation_error_handle_func` | Имя функции для обработки ошибок валидации           | `HandleValidationError`        |
| `disable_grpc_interceptor`     | Отключить генерацию gRPC интерцептора                | `false`                        |


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
    - Вызывает метод gRPC-сервера через переданный interceptor(вкл по дефолту).
    - Возвращает JSON-ответ или ошибку.

## Лицензия

[MIT License](LICENSE)

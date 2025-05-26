# ProtobufSymbol

**ProtobufSymbol** — это утилита для парсинга `.proto`-файлов и извлечения символов верхнего уровня:  
`import`, `service`, `rpc (как method)`, `enum`, `message`.

## Сборка и запуск
Убедитесь, что у вас установлен Go 1.23.2 или выше.

### Makefile

В проекте есть `Makefile`, в котором доступны следующие команды:

#### `make build`

Собирает тулзу `protosym` под все основные платформы:

- `bin/protosym-linux` — Linux (amd64)
- `bin/protosym` — macOS (amd64)
- `bin/protosym.exe` — Windows (amd64)

```bash
make build-darwin   # для macOS
```
```bash
make build-linux    # для Linux
```
```bash
make build-windows  # для Windows
```

#### `make test_files` 

Собирает тулзу под текущую ОС, затем запускает её на всех .proto-файлах в папке mydir/.

```bash
make test_files
```

Пример вывода: 
→ mydir/example.proto
Example service 9:9-16
ExampleRPC method 10:7-16
ExampleEnum enum 13:6-17
ExampleRPCRequest message 19:9-26
ExampleRPCResponse message 26:9-27


#### Поддержка:
	•	import
	•	service
	•	rpc → method
	•	enum
	•	message
Пропускает вложенные message, enum, oneof

#### Структура проекта

.
.├── cmd/               # main.go
.├── internal/          # реализация Parser
.├── mydir/             # тестовые .proto-файлы
.├── bin/               # готовые бинарники
.├── Makefile
.├── README.md

<pre>
## 🗂 Структура проекта

protobufsymbol/                   # Корневая директория проекта
├── cmd/                          # Входная точка (main.go)
│   └── main.go
├── internal/                     # Логика парсинга .proto-файлов
│   └── parser.go
├── mydir/                        # Примеры .proto-файлов для тестов
│   ├── example.proto
│   └── file1.proto
├── bin/                          # Скомпилированные бинарники (после make build)
│   ├── protosym                  # macOS
│   ├── protosym-linux            # Linux
│   └── protosym.exe              # Windows
├── Makefile                      # Скрипты сборки и тестирования
└── README.md                     # Документация проекта
</pre>

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
make build-linux    # для Linux
make build-windows  # для Windows
```

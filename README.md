# lld-logger
# Logger System (LLD)

## Requirements
- Support log levels: DEBUG, INFO, WARN, ERROR
- Multiple appenders: Console, File (extensible)
- Configurable log format (timestamp, level, message)
- Thread-safe logging
- Should be easy to add new output sinks

## Assumptions
- In-memory config is fine (no DB)
- File appender writes to a local file path
- Async logging is optional (bonus)

## Design
### Core Components
- Logger (public API)
- Formatter
- Appender (interface): ConsoleAppender, FileAppender
- LogLevel
- Config

### Design Patterns
- Strategy Pattern for Appenders
- Optional Chain of Responsibility / Composite for multi-appenders

## APIs
- `Log(level, msg)`
- `Debug/Info/Warn/Error`

## How to run
```bash
go run ./cmd

package internal

type Appender interface {
	Append(event LogEvent) error
	Name() string
}

type Formatter interface {
	Format(event LogEvent) string
}

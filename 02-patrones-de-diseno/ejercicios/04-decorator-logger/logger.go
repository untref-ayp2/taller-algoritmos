package logger

type Logger interface {
	Log(mensaje string) string
}

type LoggerSimple struct{}

func (l LoggerSimple) Log(mensaje string) string {
	// TODO: implementar
	return ""
}

type TimestampDecorator struct {
	logger Logger
}

func (d TimestampDecorator) Log(mensaje string) string {
	// TODO: implementar (prefijar con timestamp fijo "[12:00:00]" para tests)
	return ""
}

type Nivel int

const (
	INFO Nivel = iota
	WARN
	ERROR
)

type LevelDecorator struct {
	logger Logger
	nivel  Nivel
}

func (d LevelDecorator) Log(mensaje string) string {
	// TODO: implementar (prefijar con [INFO], [WARN] o [ERROR])
	return ""
}

type UpperDecorator struct {
	logger Logger
}

func (d UpperDecorator) Log(mensaje string) string {
	// TODO: implementar (convertir el resultado a mayúsculas)
	return ""
}

func nivelString(n Nivel) string {
	switch n {
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

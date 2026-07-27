package logger

import "testing"

func TestLoggerSimple(t *testing.T) {
	l := LoggerSimple{}
	got := l.Log("hola")
	if got != "hola" {
		t.Errorf("LoggerSimple.Log(%q) = %q; esperado %q", "hola", got, "hola")
	}
}

func TestTimestampDecorator(t *testing.T) {
	l := TimestampDecorator{logger: LoggerSimple{}}
	got := l.Log("hola")
	esperado := "[12:00:00] hola"
	if got != esperado {
		t.Errorf("TimestampDecorator.Log(%q) = %q; esperado %q", "hola", got, esperado)
	}
}

func TestDecoradoresAnidados(t *testing.T) {
	l := UpperDecorator{
		logger: LevelDecorator{
			logger: LoggerSimple{},
			nivel:  INFO,
		},
	}
	got := l.Log("hola")
	esperado := "[INFO] HOLA"
	if got != esperado {
		t.Errorf("Anidados.Log(%q) = %q; esperado %q", "hola", got, esperado)
	}
}

package notificador

import "testing"

func TestEmailAdapter(t *testing.T) {
	adapter := EmailAdapter{mailer: Mailer{}}
	err := adapter.Enviar(Mensaje{Destino: "test@example.com", Cuerpo: "Hola"})
	if err != nil {
		t.Errorf("EmailAdapter.Enviar() error = %v", err)
	}
}

func TestSMSAdapter(t *testing.T) {
	adapter := SMSAdapter{gateway: SMSGateway{}}
	err := adapter.Enviar(Mensaje{Destino: "+123456789", Cuerpo: "Hola"})
	if err != nil {
		t.Errorf("SMSAdapter.Enviar() error = %v", err)
	}
}

func TestPushAdapter(t *testing.T) {
	adapter := PushAdapter{service: PushService{}}
	err := adapter.Enviar(Mensaje{Destino: "token-abc", Cuerpo: "Hola"})
	if err != nil {
		t.Errorf("PushAdapter.Enviar() error = %v", err)
	}
}

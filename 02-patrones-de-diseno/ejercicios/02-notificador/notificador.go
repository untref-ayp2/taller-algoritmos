package notificador

type Mensaje struct {
	Destino string
	Cuerpo  string
}

type Notificador interface {
	Enviar(m Mensaje) error
}

// --- Canales existentes (interfaces incompatibles) ---

type Mailer struct{}

func (m Mailer) EnviarEmail(to, subject, body string) error {
	// TODO: implementar
	return nil
}

type SMSGateway struct{}

func (s SMSGateway) EnviarSMS(numero, texto string) error {
	// TODO: implementar
	return nil
}

type PushService struct{}

func (p PushService) EnviarPush(deviceToken, titulo, mensaje string) error {
	// TODO: implementar
	return nil
}

// --- Adapters ---

type EmailAdapter struct {
	mailer Mailer
}

func (a EmailAdapter) Enviar(m Mensaje) error {
	// TODO: implementar usando a.mailer.EnviarEmail
	return nil
}

type SMSAdapter struct {
	gateway SMSGateway
}

func (a SMSAdapter) Enviar(m Mensaje) error {
	// TODO: implementar usando a.gateway.EnviarSMS
	return nil
}

type PushAdapter struct {
	service PushService
}

func (a PushAdapter) Enviar(m Mensaje) error {
	// TODO: implementar usando a.service.EnviarPush
	return nil
}

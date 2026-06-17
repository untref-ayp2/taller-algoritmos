package adapter

// Medible es la interfaz que espera el cliente.
type Medible interface {
	Medir() float64
}

// Robot es el Adaptado: un robot que mide en metros y centímetros.
type Robot struct{}

// Medir devuelve la distancia en metros y centímetros.
func (r *Robot) Medir() (int, int) {
	// TODO: devolver metros y centímetros de una medición
	return 0, 0
}

// RobotAdapter envuelve un Robot y adapta su interfaz a Medible.
// Convierte metros y centímetros a pulgadas.
type RobotAdapter struct {
	robot *Robot
}

// NewRobotAdapter crea un nuevo adaptador para el robot dado.
func NewRobotAdapter(robot *Robot) *RobotAdapter {
	return &RobotAdapter{robot: robot}
}

// Medir convierte la medición del robot a pulgadas.
// 1 pulgada = 2.54 centímetros.
func (a *RobotAdapter) Medir() float64 {
	// TODO: obtener metros y centímetros del robot
	// y convertirlos a pulgadas
	return 0
}

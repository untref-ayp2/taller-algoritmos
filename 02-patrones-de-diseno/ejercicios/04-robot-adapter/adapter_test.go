package adapter

import "testing"

func TestRobotAdapterDevuelveFloat64(t *testing.T) {
	robot := &Robot{}
	adapter := NewRobotAdapter(robot)
	resultado := adapter.Medir()
	if resultado == 0 {
		t.Error("RobotAdapter.Medir() no debería devolver 0")
	}
}

func TestRobotAdapterConValoresConocidos(t *testing.T) {
	// 1 metro = 100 cm = 39.3701 pulgadas
	// 0 cm adicionales
	adapter := &RobotAdapter{robot: &Robot{}}
	pulgadas := adapter.Medir()
	if pulgadas <= 0 {
		t.Error("se esperaba un valor positivo en pulgadas")
	}
}

func TestRobotAdapterImplementaMedible(t *testing.T) {
	var _ Medible = (*RobotAdapter)(nil)
}

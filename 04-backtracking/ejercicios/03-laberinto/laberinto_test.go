package laberinto

import "testing"

func TestResolver(t *testing.T) {
	lab := [][]int{
		{0, 0, 0, 1},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{1, 1, 1, 0},
	}
	camino := Resolver(lab)
	if camino == nil || len(camino) == 0 {
		t.Fatal("Resolver devolvió nil o vacío, debería encontrar camino")
	}
	// Verificar que empieza y termina bien
	primero := camino[0]
	if primero.Fila != 0 || primero.Col != 0 {
		t.Errorf("el camino debe empezar en (0,0), empieza en %v", primero)
	}
	ultimo := camino[len(camino)-1]
	n := len(lab)
	if ultimo.Fila != n-1 || ultimo.Col != n-1 {
		t.Errorf("el camino debe terminar en (%d,%d), termina en %v", n-1, n-1, ultimo)
	}
	// Verificar que cada paso es válido
	for i := 1; i < len(camino); i++ {
		ant := camino[i-1]
		act := camino[i]
		if act.Fila < ant.Fila || act.Col < ant.Col {
			t.Errorf("paso hacia atrás en %v → %v", ant, act)
		}
		dist := (act.Fila - ant.Fila) + (act.Col - ant.Col)
		if dist != 1 {
			t.Errorf("salto inválido de %v a %v", ant, act)
		}
		if lab[act.Fila][act.Col] == 1 {
			t.Errorf("el camino pasa por obstáculo en %v", act)
		}
	}
}

func TestResolverSinCamino(t *testing.T) {
	lab := [][]int{
		{0, 1},
		{1, 0},
	}
	camino := Resolver(lab)
	if camino != nil && len(camino) > 0 {
		t.Error("debería devolver nil, no hay camino")
	}
}

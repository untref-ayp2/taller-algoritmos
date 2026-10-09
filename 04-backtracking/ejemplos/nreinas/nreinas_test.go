package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestNReinasGeneraSolucion(t *testing.T) {
	for _, n := range []int{1, 4, 8} {
		sol := NReinas(n)
		if len(sol) != n {
			t.Fatalf("NReinas(%d) = %v, se esperaban %d filas", n, sol, n)
		}
		if !esSolucionValida(sol) {
			t.Errorf("NReinas(%d) = %v no es una solución válida", n, sol)
		}
	}
}

func TestNReinasSinSolucion(t *testing.T) {
	for _, n := range []int{2, 3} {
		if sol := NReinas(n); sol != nil {
			t.Errorf("NReinas(%d) = %v, se esperaba nil", n, sol)
		}
	}
}

func TestNReinasTodas(t *testing.T) {
	casos := []struct {
		n      int
		quiero int
	}{
		{1, 1}, {2, 0}, {3, 0}, {4, 2}, {5, 10}, {6, 4},
	}

	for _, caso := range casos {
		soluciones := NReinasTodas(caso.n)
		if len(soluciones) != caso.quiero {
			t.Errorf("NReinasTodas(%d) dio %d soluciones, se esperaban %d",
				caso.n, len(soluciones), caso.quiero)
		}
		for _, sol := range soluciones {
			if len(sol) != caso.n || !esSolucionValida(sol) {
				t.Errorf("NReinasTodas(%d) incluye %v, que no es válida", caso.n, sol)
			}
		}
	}
}

func TestNReinasPaso(t *testing.T) {
	intentos := 0
	factibles := 0
	sol := NReinasConPaso(4, ObservadorReinas{
		Intento: func(fila, columna int, factible bool, parcial []int) {
			intentos++
			if len(parcial) != fila {
				t.Errorf("en el intento (fila=%d, columna=%d) la parcial tiene %d reinas",
					fila, columna, len(parcial))
			}
			if factible {
				factibles++
			}
		},
	})

	if intentos == 0 || factibles == 0 {
		t.Fatalf("el callback recibió %d intentos (%d factibles), se esperaban más de 0", intentos, factibles)
	}
	if len(sol) != 4 || !esSolucionValida(sol) {
		t.Errorf("NReinasConPaso(4) = %v, se esperaba una solución válida", sol)
	}
}

func TestNReinasPasoSeDetiene(t *testing.T) {
	intentos := 0
	NReinasConPaso(6, ObservadorReinas{
		Intento: func(fila, columna int, factible bool, parcial []int) {
			intentos++
		},
		Solucion: func(solucion []int) bool {
			return false // no seguir buscando
		},
	})

	// La primera solución de 6 reinas se alcanza bastante antes de terminar
	// de explorar todo el árbol.
	if intentos > 200 {
		t.Errorf("la búsqueda siguió tras la primera solución: %d intentos", intentos)
	}
}

func TestNReinasPasoContinua(t *testing.T) {
	encontradas := 0
	NReinasConPaso(4, ObservadorReinas{
		Solucion: func(solucion []int) bool {
			encontradas++
			return true // seguir buscando
		},
	})

	if encontradas != 2 {
		t.Errorf("NReinasConPaso(4) encontró %d soluciones, se esperaban 2", encontradas)
	}
}

func TestTableroNReinas(t *testing.T) {
	tablero := TableroNReinas(4, []int{2, 0, 3, 1}, -1, -1, true)

	if got := strings.Count(tablero, "♛"); got != 4 {
		t.Errorf("el tablero debería tener 4 reinas, tiene %d:\n%s", got, tablero)
	}
}

func TestTableroNReinasMuestraIntento(t *testing.T) {
	tablero := TableroNReinas(4, []int{2, 0}, 2, 1, false)

	if !strings.Contains(tablero, "×") {
		t.Errorf("el intento no factible debería marcarse con '×':\n%s", tablero)
	}
}

func TestLeerCantidadReinas(t *testing.T) {
	casos := []struct {
		entrada string
		quiero  int
	}{
		{"8\n", 8},
		{"\n", 6},             // Enter vacío: 6
		{"3\n5\n", 5},         // rechaza 3 (menor que 4) y toma 5
		{"99\n2\n11\n7\n", 7}, // rechaza 99, 2 y 11
		{"cuatro\n6\n", 6},    // rechaza lo no numérico
		{"", 6},               // sin entrada: 6
		{"10\n", 10},          // borde superior
		{"4\n", 4},            // borde inferior
	}

	for _, caso := range casos {
		var salida strings.Builder
		got := LeerCantidadReinas(bufio.NewReader(strings.NewReader(caso.entrada)), &salida)
		if got != caso.quiero {
			t.Errorf("LeerCantidadReinas(%q) = %d, se esperaba %d", caso.entrada, got, caso.quiero)
		}
	}
}

func TestAnimarReinasAgotaLasSoluciones(t *testing.T) {
	var salida strings.Builder

	// 4 reinas tienen 2 soluciones: respondemos que sí a las dos.
	AnimarReinas(4, bufio.NewReader(strings.NewReader("s\ns\n")), &salida, 0)

	texto := salida.String()
	if !strings.Contains(texto, "Solución 1. Tablero de 4x4") ||
		!strings.Contains(texto, "Solución 2. Tablero de 4x4") {
		t.Errorf("la animación debería mostrar las dos soluciones:\n%s", texto)
	}
	if !strings.Contains(texto, "No hay más soluciones") {
		t.Errorf("la animación debería avisar que no hay más soluciones:\n%s", texto)
	}
}

func TestAnimarReinasSeDetiene(t *testing.T) {
	var salida strings.Builder

	AnimarReinas(4, bufio.NewReader(strings.NewReader("n\n")), &salida, 0)

	texto := salida.String()
	if !strings.Contains(texto, "Solución 1") {
		t.Errorf("la animación debería mostrar la primera solución:\n%s", texto)
	}
	if strings.Contains(texto, "Solución 2") {
		t.Errorf("la animación no debería seguir buscando si se responde que no:\n%s", texto)
	}
	if !strings.Contains(texto, "detenida") {
		t.Errorf("la animación debería avisar que la búsqueda se detuvo:\n%s", texto)
	}
}

// esSolucionValida comprueba que ninguna reina ataque a otra: no comparten
// columna ni ninguna de las dos diagonales.
func esSolucionValida(sol []int) bool {
	for fila := range sol {
		for otra := fila + 1; otra < len(sol); otra++ {
			if sol[fila] == sol[otra] ||
				fila+sol[fila] == otra+sol[otra] ||
				fila-sol[fila] == otra-sol[otra] {
				return false
			}
		}
	}
	return true
}

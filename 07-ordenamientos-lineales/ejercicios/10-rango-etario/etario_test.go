package etario

import "testing"

func TestAgruparPorEdad(t *testing.T) {
	personas := []Persona{
		{"Ana", 8},
		{"Luis", 25},
		{"Pedro", 70},
	}
	got := AgruparPorEdad(personas)
	if len(got) != 5 {
		t.Fatalf("esperados 5 grupos, obtenidos %d", len(got))
	}
	if got[0].Rango != "0-12" || len(got[0].Personas) != 1 {
		t.Errorf("grupo infantil incorrecto: %v", got[0])
	}
	if got[1].Rango != "13-17" || len(got[1].Personas) != 0 {
		t.Errorf("grupo adolescente incorrecto: %v", got[1])
	}
	if got[2].Rango != "18-35" || len(got[2].Personas) != 1 {
		t.Errorf("grupo joven incorrecto: %v", got[2])
	}
	if got[4].Rango != "61+" || len(got[4].Personas) != 1 {
		t.Errorf("grupo mayor incorrecto: %v", got[4])
	}
}

func TestAgruparMismoRango(t *testing.T) {
	personas := []Persona{
		{"Ana", 30},
		{"Luis", 28},
		{"Pedro", 35},
	}
	got := AgruparPorEdad(personas)
	grupo := got[2]
	if len(grupo.Personas) != 3 {
		t.Fatalf("esperadas 3 personas en grupo 18-35, obtenidas %d", len(grupo.Personas))
	}
	if grupo.Personas[0].Edad < grupo.Personas[1].Edad {
		t.Error("no ordenado por edad descendente")
	}
}

func TestAgruparVacio(t *testing.T) {
	got := AgruparPorEdad([]Persona{})
	if len(got) != 5 {
		t.Errorf("esperados 5 grupos, obtenidos %d", len(got))
	}
	for _, g := range got {
		if len(g.Personas) != 0 {
			t.Errorf("grupo %s no vacío", g.Rango)
		}
	}
}

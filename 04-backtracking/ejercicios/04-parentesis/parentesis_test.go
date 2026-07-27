package parentesis

import (
	"reflect"
	"testing"
)

func TestN1(t *testing.T) {
	got := GenerarParentesis(1)
	esperado := []string{"()"}
	if !reflect.DeepEqual(got, esperado) {
		t.Errorf("GenerarParentesis(1) = %v; esperado %v", got, esperado)
	}
}

func TestN3(t *testing.T) {
	got := GenerarParentesis(3)
	esperado := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	if !reflect.DeepEqual(got, esperado) {
		t.Errorf("GenerarParentesis(3) = %v; esperado %v", got, esperado)
	}
}

func TestN0(t *testing.T) {
	got := GenerarParentesis(0)
	if len(got) != 0 {
		t.Errorf("GenerarParentesis(0) = %v; esperado []", got)
	}
}

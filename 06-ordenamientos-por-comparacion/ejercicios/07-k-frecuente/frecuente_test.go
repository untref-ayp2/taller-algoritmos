package frecuente

import (
	"reflect"
	"testing"
)

func TestKMasFrecuente(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	got := KMasFrecuente(nums, 2)
	esperado := []int{1, 2}
	if !reflect.DeepEqual(got, esperado) {
		t.Errorf("KMasFrecuente = %v; esperado %v", got, esperado)
	}
}

func TestKMasFrecuenteUnico(t *testing.T) {
	got := KMasFrecuente([]int{1}, 1)
	if !reflect.DeepEqual(got, []int{1}) {
		t.Errorf("KMasFrecuente = %v; esperado [1]", got)
	}
}

func TestKMasFrecuenteMaxFreq(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3, 3}
	got := KMasFrecuente(nums, 1)
	if !reflect.DeepEqual(got, []int{3}) {
		t.Errorf("KMasFrecuente = %v; esperado [3]", got)
	}
}

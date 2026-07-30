# Iterator — Lista Enlazada

Ejemplo del patrón *Iterator* con una API de dos métodos. `Siguiente()` avanza
y devuelve `true` si hay elemento; `Valor()` retorna el valor actual. El patrón
de uso es `for it.Siguiente() { fmt.Println(it.Valor()) }`.

```bash
go run main.go
```

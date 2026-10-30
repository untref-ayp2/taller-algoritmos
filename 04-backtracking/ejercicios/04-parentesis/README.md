# Generar Paréntesis Balanceados

**Función a implementar:** `GenerarParentesis(n int) []string`

## El problema

Generar **todas** las secuencias de `n` pares de paréntesis balanceados
válidas. Para `n = 3` hay 5: `((()))`, `(()())`, `(())()`, `()(())` y
`()()()`. La función las devuelve en un slice de strings; para `n = 0`
devuelve un slice vacío.

## Conceptos clave

- **Secuencia balanceada:** misma cantidad de `(` y `)`, y en **ningún
  punto** del prefijo hay más `)` que `(`. Por ejemplo `())(` tiene 2
  y 2, pero es inválida: al tercer carácter ya cerró de más.

- **Longitud fija:** `n` pares = `2n` caracteres. Ese es el
  `esSolucion`: la secuencia está completa.

- **Candidatos:** en cada paso sólo caben dos opciones, agregar `(` o
  agregar `)`.

- **Podas:**
  - Sólo se puede abrir si `abiertas < n`: si ya se pusieron los `n`
    paréntesis de apertura, no entran más.
  - Sólo se puede cerrar si `cerrados < abiertas`: si no, el prefijo
    quedaría con más `)` que `(` y la secuencia ya no se puede arreglar.

## Estrategia de backtracking

| Elemento del esquema | Definición en este ejercicio |
|---|---|
| `solucionParcial` | un string con los caracteres elegidos hasta ahora |
| `esSolucion` | `len(actual) == 2*n` |
| `extender` | dos opciones: `(` si `abiertas < n`, `)` si `cerrados < abiertas` |
| `esFactible` | las dos condiciones de arriba |
| `registrar` | concatenar el carácter a la secuencia |
| `borrar` | la vuelta atrás es implícita: la recursión recibe el string    nuevo como parámetro y, al volver, la llamada anterior sigue con    su propio `actual` |

**El orden de exploración importa:** los tests comparan el resultado
con `reflect.DeepEqual`, o sea que exigen el orden exacto. Hay que
probar `(` **antes** que `)` en cada nodo.

## Ejemplo a mano

`n = 2`. Árbol de búsqueda completo (cada nodo muestra la secuencia
parcial y entre paréntesis los contadores `(abiertas, cerrados)`):

```
"" (0,0)
└── ( → "(" (1,0)
    ├── (  →  "((" (2,0)
    │        └── )  →  "(()" (2,1)
    │                 └── )  →  "(())" ✓
    └── )  →  "()" (1,1)
             └── (  →  "()(" (2,1)
                      └── )  →  "()()" ✓
```

De `"("` se puede abrir o cerrar; de `"(("` ya no se puede abrir
(`abiertas = n`), sólo cerrar; de `"()"` no se puede cerrar
(`cerrados = abiertas`), sólo abrir. Las podas son las que dan forma al
árbol: sin ellas habría ramas inválidas que explorar.

Resultado para `n = 2`: `["(())", "()()"]`.

Para `n = 3`, la misma exploración da el orden exacto que exige el
test: `["((()))", "(()())", "(())()", "()(())", "()()()"]`.

Casos borde:

- `n = 0` → slice vacío (no hay caracteres que poner).
- `n = 1` → `["()"]`, la única posible.

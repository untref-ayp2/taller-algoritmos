# Código Huffman

**Funciones a implementar:**

- `Construir(frecuencias map[rune]int) *ArbolHuffman`
- `(a *ArbolHuffman) Codigos() map[rune]string`
- `(a *ArbolHuffman) Comprimir(texto string) string`
- `(a *ArbolHuffman) Descomprimir(codigo string) string`

## El problema

**¿Cómo se guarda un texto en la computadora?** Cada carácter (cada letra,
cada dígito, el espacio) se guarda como un número, y ese número se escribe
en binario. En la representación estándar de texto (ASCII), cada carácter
del alfabeto ocupa **1 byte, y un byte son 8 bits**, o sea 8 dígitos
binarios: la letra `a` se guarda como `01100001` (que es el 97 en binario,
porque en la tabla ASCII la `a` es el número 97), la `b` como `01100010`,
y así con todo. O sea que:

- `"hola"` = 4 caracteres = **4 bytes = 32 bits**;
- un texto de 1000 caracteres = 1000 bytes (1 kilobyte);
- un libro entero = millones de bytes.

Eso está bien, pero es muy "caro": siempre paga lo mismo (8 bits) un
carácter que aparece mil veces que uno que aparece una sola vez.

**Idea de la compresión:** en vez de darle a todos los caracteres los
mismos 8 bits, darle a cada uno una cadena binaria de largo distinto:

- los caracteres **frecuentes** → cadena **corta** (1, 2 bits);
- los caracteres **raros** → cadena **larga** (puede tener varios bits).

Como el total de bits del texto comprimido es
`Σ (frecuencia × longitud del código)`, achicarle bits a los que más
aparecen es donde está el ahorro. Ejemplo: en `"aaabbc"` la `a` aparece 4
veces, la `b` 2 y la `c` 1. Sin comprimir: 7 caracteres × 8 bits = 56
bits. Si a `a` le damos el código `0`, a `b` el `10` y a `c` el `11`,
el texto comprimido queda `0 0 0 0 10 10 11` = **10 bits**. Una locura de
ahorro; después vemos cómo llegar a esa asignación de códigos.

**La trampa: poder descomprimir sin ambigüedad.** Comprimir no sirve de
nada si después no se puede recuperar el texto original exactamente. Y
acá aparece la condición central del ejercicio. Primero, un concepto:

- **Prefijo:** cadena A es prefijo de cadena B si B *empieza con* A. Por
  ejemplo, `"hol"` es prefijo de `"holanda"`, y `"anda"` no es prefijo de
  `"holanda"`.

La regla es: **el código de una letra no puede ser prefijo del código de
otra** (y ni hablemos de que tengan exactamente el mismo código: eso sería
peor todavía). Si esa regla se viola, la descompresión es ambigua. Ejemplo
con dos letras: si a `a` le damos `0` y a `b` le damos `01`, y el texto
comprimido resultante es `01`, ¿qué es? ¿Una `b` completa? ¿O una `a`
(`0`) seguida de un `1` suelto que no significa nada? No hay forma de
saberlo: el sistema no sirve.

Un caso que **sí** está bien: `b = 10` y `c = 11`. Las dos empiezan con
`1`, comparten ese primer bit, pero enseguida se diferencian: al leer,
después del `1` basta con mirar el siguiente bit para saber a qué letra
pertenece. Lo que no puede pasar es que una cadena esté completa "en el
medio" de la otra.

A esta propiedad se la llama **códigos prefijo-libres** (o códigos
instantáneos): ningún código es prefijo de ningún otro. Gracias a eso, la
descompresión puede leer los bits de a uno, de izquierda a derecha, y en
cuanto complete un código ya sabe qué letra era, sin necesidad de
separadores ni de saber de antemano cuánto mide cada código.

**Qué hay que implementar:** a partir de las *frecuencias* de cada
carácter (no del texto en sí), construir un árbol de Huffman que define
los códigos, generar el mapa carácter → código, comprimir un texto
(reemplazando cada carácter por su código) y descomprimirlo de vuelta.

## Conceptos clave

- **Bit / byte:** un bit es un dígito binario: 0 o 1. Un byte son 8 bits.
  Cada carácter de texto sin comprimir ocupa 1 byte = 8 bits.

- **Frecuencia:** cantidad de veces que aparece cada carácter en el texto.
  Es la única entrada del ejercicio: `map[rune]int` (carácter → cantidad).

- **Código binario:** una cadena hecha sólo con los caracteres `'0'` y
  `'1'`. Ojo: en este ejercicio `Comprimir` devuelve una `string` donde
  cada carácter es `'0'` o `'1'` (texto binario), no un archivo con bits
  empaquetados.

- **Prefijo:** definido arriba. La propiedad **prefijo-libre** es que
  ningún código de la tabla sea prefijo de otro (ni igual a otro).

- **Árbol binario:** estructura hecha de *nodos*, donde cada nodo tiene a
  lo sumo dos hijos, uno izquierdo (`Izq`) y uno derecho (`Der`). El nodo
  de arriba de todo es la **raíz**; los nodos sin hijos son las **hojas**.
  Para llegar de la raíz a un nodo hay que bajar por un **camino**, y ese
  camino se puede describir con binario: bajar a la izquierda = `'0'`,
  bajar a la derecha = `'1'`.

- **Árbol de Huffman:** el árbol binario que construye este ejercicio. Las
  **hojas** guardan los caracteres (`Car`) y sus frecuencias (`Freq`);
  los **nodos internos** no representan ningún carácter y guardan en
  `Freq` la suma de las frecuencias de sus dos hijos. El código de un
  carácter es el camino desde la raíz hasta su hoja.

  Y acá está la magia: como cada carácter vive en una **hoja** (y una hoja
  no tiene hijos), ningún camino termina "en el medio" del camino de otro
  → los códigos salen prefijo-libres **automáticamente**, sin que haya que
  verificarlo a mano.

- **Longitud promedio de código:** `Σ(frecuencia × longitud del código) /
  Σ(frecuencias)`, es decir, el promedio de bits por carácter. Es lo que
  hay que minimizar: un código fijo de 8 bits para todo, o de 3 bits si
  hay 6 símbolos, son el punto de partida a vencer.

**Nota sobre la implementación:** el comentario del esqueleto menciona una
*cola de prioridad*: es simplemente la idea de una colección de la que
siempre extraemos el elemento de menor frecuencia (el de mayor prioridad
para este algoritmo). Más adelante se verá una estructura adecuada para
ese trabajo; por ahora puede resolverse recorriendo la lista de nodos en
cada paso y buscando los dos de menor frecuencia.

## Estrategia greedy

1. Crear una hoja por cada carácter con su frecuencia.
2. Mientras haya más de un nodo: tomar los dos de **menor frecuencia**,
   crear un nodo nuevo con frecuencia igual a la suma de ambos y con esos
   dos como hijos, y devolver el nodo nuevo al conjunto.
3. El único nodo que queda es la raíz del árbol.
4. `Codigos()`: recorrer el árbol desde la raíz; al llegar a una hoja,
   registrar `Car` → el prefijo acumulado (`0` a la izquierda, `1` a la
   derecha).
5. `Comprimir()`: concatenar el código de cada carácter del texto.
6. `Descomprimir()`: recorrer la raíz bit a bit (`'0'` → izquierda, `'1'` →
   derecha); al caer en una hoja, emitir su carácter y volver a la raíz.

**¿Por qué funciona?** Combinar siempre los dos menos frecuentes "hunde"
a los caracteres raros en el árbol: quedan en hojas profundas, con códigos
largos. Los frecuentes quedan cerca de la raíz, con códigos cortos. Como
los caracteres raros aparecen pocas veces, que sus códigos sean largos
apenas pesa en el total; en cambio, ahorrarle bits a los frecuentes rinde
en cada aparición. Es la decisión local del greedy: en cada paso empeoro a
los que menos importan para beneficiar al resto.

Si hay **empates** en frecuencia, el orden en que se combinen puede dar un
árbol distinto (y por lo tanto códigos distintos); el costo total es
óptimo igualmente. Por eso los tests sólo verifican la cantidad de códigos,
que ninguno esté vacío y que el ciclo compresión/descompresión sea fiel.

## Ejemplo a mano

### Ejemplo corto: `"aaabbc"`

Frecuencias: `a:4, b:2, c:1` (7 en total).

| Paso | Se combinan  | Nuevo nodo | Conjunto restante  |
|------|--------------|------------|--------------------|
| —    | —            | hojas      | a(4) b(2) c(1)    |
| 1    | c(1) + b(2)  | n(3)       | a(4) n(3)         |
| 2    | a(4) + n(3)  | raíz(7)    | queda la raíz     |

```
        (7)
       /    \
    a(4)    (3)
           /   \
        b(2)   c(1)
```

Códigos: `a → 0` (1 bit), `b → 10` (2 bits), `c → 11` (2 bits). Son
prefijo-libres: ninguno está dentro de otro.

- Sin comprimir: 7 caracteres × 8 bits = **56 bits**.
- Comprimido: `0 0 0 0 10 10 11` = 4×1 + 2×2 + 1×2 = **10 bits**.

### Ejemplo del test: frecuencias `a:5, b:9, c:12, d:13, e:16, f:45`

100 apariciones en total.

**Paso 1:** combinar los dos de menor frecuencia hasta que quede uno solo.

| Paso | Se combinan     | Nuevo nodo | Conjunto restante                  |
|------|-----------------|------------|------------------------------------|
| —    | —               | hojas      | a(5) b(9) c(12) d(13) e(16) f(45) |
| 1    | a(5) + b(9)     | n1(14)     | n1(14) c(12) d(13) e(16) f(45)    |
| 2    | c(12) + d(13)   | n2(25)     | n1(14) e(16) n2(25) f(45)         |
| 3    | n1(14) + e(16)  | n3(30)     | n2(25) n3(30) f(45)               |
| 4    | n2(25) + n3(30) | n4(55)     | f(45) n4(55)                      |
| 5    | f(45) + n4(55)  | raíz(100)  | queda la raíz                     |

**Árbol resultante** (izquierda = `'0'`, derecha = `'1'`):

```
           (100)
          /     \
       f(45)    (55)
               /    \
            (25)    (30)
            / \     /  \
         c(12) d(13) (14) e(16)
                  / \
              a(5)  b(9)
```

**Paso 2:** leer los códigos recorriendo la raíz hasta cada hoja.

| Car | Frecuencia | Código | Bits |
|-----|------------|--------|------|
| f   | 45         | `0`      | 1 |
| c   | 12         | `100`    | 3 |
| d   | 13         | `101`    | 3 |
| e   | 16         | `111`    | 3 |
| a   | 5          | `1100`   | 4 |
| b   | 9          | `1101`   | 4 |

Ningún código es prefijo de otro: todos terminan en una hoja distinta del
árbol, así que la descompresión es inambigua.

**Compresión** de `"abcdef"`: concatenar los códigos →

```
1100 1101 100 101 111 0   →   "110011011001011110"   (18 bits)
```

(Sin comprimir serían 6 caracteres × 8 bits = 48 bits.)

**Descompresión** de `"110011011001011110"`:

| Bits leídos | Camino en el árbol                     | Se emite | Se vuelve a la raíz |
|-------------|----------------------------------------|----------|---------------------|
| `1 1 0 0`   | der → der → izq → izq                 | `a`      | sí                  |
| `1 1 0 1`   | der → der → izq → der                 | `b`      | sí                  |
| `1 0 0`     | der → izq → izq                       | `c`      | sí                  |
| `1 0 1`     | der → izq → der                       | `d`      | sí                  |
| `1 1 1`     | der → der → der                       | `e`      | sí                  |
| `0`         | izq                                    | `f`      | sí                  |

Resultado: `"abcdef"` — el ciclo compresión/descompresión cierra. Fijate
cómo en ningún momento hay que "adivinar" dónde corta una letra: al llegar
a una hoja, el camino terminó.

**Ahorro:** la longitud promedio es `(5×4 + 9×4 + 12×3 + 13×3 + 16×3 +
45×1) / 100 = 2,24` bits por carácter, contra 3 bits que pediría un código
fijo de 3 bits para 6 símbolos: 25% menos para cualquier texto con esas
frecuencias.

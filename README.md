# Taller de Algoritmos

Repositorio complementario de la sección **Diseño de Algoritmos** de los apuntes de Algoritmos y Programación II.

## Estructura

```
01-recursividad/                   # ← capítulo 4-1 (Recursividad y DyC)
├── ejemplos/
│   ├── factorial/                 # factorial recursivo
│   ├── par-impar/                 # recursión indirecta
│   └── busqueda-binaria/          # división y conquista
└── ejercicios/
    ├── 01-palindromo/             # esqueleto con tests
    ├── 02-contar-caracter/        # esqueleto con tests
    ├── 03-hanoi/                  # esqueleto con tests
    ├── 04-busqueda-ternaria/      # esqueleto con tests
    └── 05-pico/                   # esqueleto con tests
```

Cada directorio `ejercicios/` contiene un `README.md` con los enunciados y esqueletos con tests.

## Cómo usar

```bash
git clone https://github.com/untref-ayp2/taller-algoritmos.git
cd taller-algoritmos
```

Para ejecutar un ejemplo:

```bash
go run ./01-recursividad/ejemplos/factorial
```

Para correr los tests de un ejercicio:

```bash
go test ./01-recursividad/ejercicios/01-palindromo/...
```

Para correr todos los tests:

```bash
go test ./...
```

## Requisitos

- Go 1.22 o superior
- Opcional: [golangci-lint](https://golangci-lint.run/) para linting local

## Licencia

MIT

# Cómo trabajar en este repositorio

## Ejecutar los tests

Ejecutá todos los tests:

    make test

Ejecutar los tests con detalle de cada caso:

    make test-v

Ejecutar los tests de un ejercicio específico:

    go test -v ./01-recursividad/ejercicios/01-palindromo/...

## Feedback automático (CI)

Cada vez que hacés `git push`, GitHub Actions ejecuta los tests automáticamente.
Vas a ver el resultado en la pestaña **Actions** de tu repositorio.

Si los tests pasan: se muestra un tick verde.
Si algún test falla: se muestra una cruz roja — revisá el error, corregí el código y volvé a pushear.

## Tu repositorio tiene un PR "Feedback"

Cuando empezaste la assignment, GitHub Classroom creó automáticamente un Pull Request
llamado **"Feedback"** en tu repositorio. Este PR es tu canal de comunicación con los docentes.

**No cierres este PR. Si lo cerrás, perdés el canal de ayuda.**

### Cómo funcionan los commits en el PR "Feedback"

Cada vez que hacés `git push`, tus commits aparecen automáticamente en el PR "Feedback".
No necesitás hacer nada especial — el PR se actualiza solo.

### Cómo pedir ayuda

1. **Commitear y pushear** tu intento (aunque no esté completo):

       git add .
       git commit -m "intento ejercicio palindromo"
       git push

2. **Ir al PR "Feedback"** — en la pestaña Pull Requests de tu repositorio,
   click en el PR llamado "Feedback".

3. **Agregar un comentario** en el PR explicando:
   - Qué ejercicio estás haciendo.
   - Qué intentaste.
   - Qué error estás viendo.

4. **Mencionar al docente** — si un docente te dio su usuario de GitHub,
   podés arrobarlo en el comentario para que reciba una notificación:

       @nombredeusuario ayuda con el ejercicio de palindromo

### Normas para pedir ayuda

- **Un solo PR por repositorio** — no necesitás abrir más PRs; el PR "Feedback"
  ya existe y es el canal correcto.
- **Antes de pedir ayuda** — intentá resolver el ejercicio por tu cuenta al
  menos 3 veces. En el comentario explicá qué intentaste.
- **No pedir que te lo resuelvan** — el objetivo es que aprendas, no que te
  den la solución.
- **No cerrar el PR** — el PR "Feedback" debe mantenerse abierto para que los
  docentes puedan verte y responderte.

## Convenciones de código

- **Ejercicios**: cada uno es un package nombrado (no `package main`).
  Cada paquete tiene su propio `*_test.go`.
- **Indentación**: tabs (ya configurado en `.editorconfig`).

## Estructura de capítulos

```
01-recursividad/              # capítulo 4-1
02-patrones-de-diseno/        # capítulo 4-2
03-algoritmos-avidos/         # capítulo 4-3
04-backtracking/              # capítulo 4-4
05-programacion-dinamica/     # capítulo 4-5
06-ordenamientos-recursivos/  # capítulo 4-6
07-ordenamientos-lineales/     # capítulo 4-7
```

Cada capítulo tiene `ejercicios/` (esqueletos con tests) y los primeros dos
también tienen `ejemplos/` (código funcional).

## Comandos útiles

    make fmt    # formatear todo el código
    make lint   # verificar estilo con linter
    make build  # compilar todo sin ejecutar
    make clean  # limpiar archivos generados

    go run ./01-recursividad/ejemplos/factorial  # ejecutar un ejemplo

## Requisitos

- Go 1.22 o superior.
- Opcional: golangci-lint (https://golangci-lint.run/) para verificar estilo
  localmente.

# Adapter — Robot

Ejemplo del patrón *Adapter*. Un `Robot` mide distancias en metros y
centímetros (`Medir() (int, int)`), pero el cliente espera recibir
la distancia en pulgadas (`Medir() float64`). El `RobotAdapter`
envuelve al robot y convierte la medición.

```bash
go run main.go
```

# Ejercicios: Patrones de Diseño

Los ejercicios cubren los tres patrones vistos en el apunte: **Composite**,
**Adapter** e **Iterator**. Los ejemplos de referencia están en `../ejemplos/`.

---

1. **Sistema de Figuras.** Implementar una jerarquía de figuras geométricas
   usando el patrón *Composite* para que un `Dibujo` pueda estar compuesto
   tanto por figuras simples como por otros dibujos. Cada figura debe
   saber calcular su área y perímetro.
   → `01-sistema-figuras/`

2. **Notificador.** Implementar un sistema de notificaciones usando el
   patrón *Adapter* que permita enviar mensajes a través de diferentes
   canales (email, SMS, push). Los canales existentes tienen interfaces
   incompatibles; el adapter debe unificarlas.
   → `02-notificador/`

3. **Iterador de Lista Doble.** Implementar un iterador externo para una
   lista doblemente enlazada usando el patrón *Iterator*. La interfaz del
   iterador tiene dos métodos: `Siguiente()` (avanza hacia adelante) y
   `Anterior()` (retrocede). El patrón de uso es `for it.Siguiente() { ... }`.
   → `03-iterador-lista-doble/`

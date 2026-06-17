# Ejercicios: Patrones de Diseño

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

3. **Iterador de Listas.** Implementar un iterador externo para una lista
   simplemente enlazada usando el patrón *Iterator*. Debe permitir
   recorrer la lista sin exponer su estructura interna.
   → `03-iterador-listas/`

4. **Robot Adapter.** Un robot mide distancias en metros y centímetros,
   pero el sistema del cliente espera recibir la distancia en pulgadas.
   Implementar un *Adapter* que convierta la interfaz del robot a la
   interfaz esperada por el cliente (1 pulgada = 2.54 cm).
   → `04-robot-adapter/`

5. **Tren de Figuras.** Usando el patrón *Composite*, modelar figuras
   geométricas (rectángulo, círculo, triángulo) que puedan agruparse
   en un `Grupo` para calcular el área total. El grupo debe poder
   contener tanto figuras simples como otros grupos.
   → `05-tren-figuras/`

6. **Iterador Clásico.** Implementar el patrón *Iterator* con el estilo
   clásico GoF: un iterador con métodos `Primero()`, `Siguiente()`,
   `HaySiguiente()` y `Actual()` para recorrer una lista enlazada.
   → `06-iterador-clasico/`

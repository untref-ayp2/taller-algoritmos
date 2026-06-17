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

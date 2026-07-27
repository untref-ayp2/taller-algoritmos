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

4. **Decorator para Logging.** Implementar el patrón *Decorator* para un
   sistema de logging componible. Un logger simple puede ser envuelto
   con decoradores que agregan timestamp, nivel (INFO/WARN/ERROR)
   o convierten a mayúsculas, de forma anidada.
   → `04-decorator-logger/`

5. **Factory Method para Figuras.** Implementar una fábrica que cree
   figuras geométricas (círculo, rectángulo, cuadrado) a partir de
   descripciones textuales como `"circulo 5"`, devolviendo error si
   el formato no es reconocido.
   → `05-factory-figuras/`

6. **Command con Undo/Redo.** Implementar el patrón *Command* para un
   mini editor de texto con soporte de deshacer y rehacer. Los comandos
   `Insertar` y `Borrar` implementan una interfaz común y un `Historial`
   administra las pilas de undo/redo.
   → `06-command-undo/`

Los ejemplos de referencia están en el directorio `../ejemplos/`.

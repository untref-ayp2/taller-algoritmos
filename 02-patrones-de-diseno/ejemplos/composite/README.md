# Composite — Sistema de Archivos

Ejemplo del patrón *Composite*. `Archivo` es un elemento simple (hoja) que
reporta su tamaño en bytes. `Carpeta` es un compuesto que contiene otros
`Componente` (archivos o carpetas) y calcula su tamaño total sumando
recursivamente. Se construye un directorio de proyecto con archivos y
subcarpetas, y se calcula su tamaño total.

```bash
go run main.go
```

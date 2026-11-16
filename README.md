# Sistema de Gestión de Streaming

Proyecto del **Aprendizaje Autónomo 1** — Programación Orientada a Objetos (UIDE).
Etapa 1: planeación del software y desarrollo mediante **programación funcional** en Go.

## Objetivo

Gestionar las operaciones básicas de una plataforma de streaming: registro e inicio
de sesión de usuarios, administración del catálogo de contenidos, y reproducción
con historial y favoritos.

## Estructura de paquetes

```
sistema-streaming/
├── go.mod                      → módulo: github.com/pinky96/sistema-streaming
├── main.go                     → punto de entrada: demuestra el flujo completo
├── usuarios/
│   └── usuarios.go             → RegistrarUsuario, IniciarSesion, ListarUsuarios
├── catalogo/
│   └── catalogo.go             → AgregarContenido, BuscarContenido, FiltrarPorGenero, ListarCatalogo
├── reproduccion/
│   └── reproduccion.go         → ReproducirContenido, AgregarAFavoritos, VerHistorial
└── docs/
    └── el video                → Url del video 
```

Cada paquete corresponde a un módulo del diseño, las funciones exportadas
(con mayúscula inicial) son las declaradas en la tabla de funciones del PDF.

## Paquetes utilizados

Solo se usan paquetes de la **biblioteca estándar** de Go (se importan, no se descargan):

| Paquete   | Uso                                              |
|-----------|--------------------------------------------------|
| `fmt`     | Formatear e imprimir mensajes en consola         |
| `errors`  | Crear errores simples de validación              |
| `strings` | Normalizar texto (espacios, mayúsculas)          |
| `sort`    | Ordenar alfabéticamente el listado de usuarios   |

No se incorporan paquetes de terceros en esta etapa.

## Cómo compilar y ejecutar

```bash
go build ./...   # compila todos los paquetes
go run .         # ejecuta la demostración
```

## Enfoque funcional

- El programa está organizado en **funciones dentro de paquetes** (sin structs).
- Las funciones **no modifican sus argumentos**: reciben el estado (mapas y slices),
  y retornan una **copia nueva** con el cambio aplicado.
- Los errores se retornan como segundo valor (`(resultado, error)`), siguiendo la
  convención de Go.

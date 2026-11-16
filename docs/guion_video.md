# Guion del video explicativo (3–5 minutos)

## 1. Presentación del sistema (±45 s)
- Saludo y nombre.
- "Elegí un Sistema de Gestión de Streaming: una plataforma que registra usuarios,
  administra un catálogo de películas y series, y gestiona la reproducción con
  historial y favoritos."
- Objetivo y alcance en una frase: qué hace y qué queda fuera (sin pagos, sin
  reproducción real de video — se simula en consola).

## 2. Recorrido por los módulos (±2 min) — compartir pantalla con VS Code
- Mostrar el árbol de carpetas: un **paquete por módulo** (`usuarios/`, `catalogo/`, `reproduccion/`).
- Abrir `usuarios/usuarios.go` y explicar **una firma completa**:
  `func RegistrarUsuario(usuarios map[string]string, nombre string, correo string) (map[string]string, error)`
  — la palabra clave `func`, el tipo va después del nombre del parámetro,
  el retorno múltiple `(map[string]string, error)`, y que la mayúscula inicial la exporta.
- Mostrar brevemente `catalogo/` (búsqueda y filtro) y `reproduccion/` (historial y favoritos).
- Destacar el enfoque funcional: las funciones retornan copias nuevas, no modifican los argumentos.

## 3. El repositorio y la compilación (±1 min) — terminal
- Mostrar `go.mod` y explicar: "el repositorio empieza con `go mod init` y el nombre
  del sistema; sin eso no puedo importar mis paquetes".
- Qué se importa: solo biblioteca estándar (`fmt`, `errors`, `strings`, `sort`) —
  nada que descargar con `go get` en esta etapa.
- Ejecutar en vivo: `go build ./...` (compila sin errores) y `go run .` (demostración).

## 4. Cierre (±30 s)
- Resumen: 10 funcionalidades en 3 módulos, repositorio en GitHub con README.
- Mostrar la página del repositorio en GitHub en el navegador.

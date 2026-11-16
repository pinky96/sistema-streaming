// Package reproduccion gestiona la reproducción de contenidos, el
// historial de visualización y la lista de favoritos.
//
// El historial y los favoritos se representan como slices de títulos
// ([]string), en orden de llegada.
package reproduccion

import (
	"errors"
	"fmt"
	"strings"
)

// ReproducirContenido simula la reproducción de un contenido y lo agrega
// al historial. Retorna un nuevo historial (no modifica el original),
// un mensaje para mostrar en pantalla y un error si el título está vacío.
func ReproducirContenido(historial []string, titulo string) ([]string, string, error) {
	titulo = strings.TrimSpace(titulo)
	if titulo == "" {
		return historial, "", errors.New("debe indicar el título a reproducir")
	}

	resultado := make([]string, len(historial), len(historial)+1)
	copy(resultado, historial)
	resultado = append(resultado, titulo)

	mensaje := fmt.Sprintf("▶ Reproduciendo: %s", titulo)
	return resultado, mensaje, nil
}

// AgregarAFavoritos agrega un título a la lista de favoritos.
// Retorna una nueva lista y un error si el título ya estaba en favoritos.
func AgregarAFavoritos(favoritos []string, titulo string) ([]string, error) {
	titulo = strings.TrimSpace(titulo)
	for _, f := range favoritos {
		if strings.EqualFold(f, titulo) {
			return favoritos, fmt.Errorf("%q ya está en favoritos", titulo)
		}
	}

	resultado := make([]string, len(favoritos), len(favoritos)+1)
	copy(resultado, favoritos)
	return append(resultado, titulo), nil
}

// VerHistorial retorna el historial de reproducción como texto numerado,
// listo para imprimir en consola.
func VerHistorial(historial []string) string {
	if len(historial) == 0 {
		return "Aún no se ha reproducido ningún contenido."
	}
	lineas := make([]string, 0, len(historial))
	for i, titulo := range historial {
		lineas = append(lineas, fmt.Sprintf("%d. %s", i+1, titulo))
	}
	return strings.Join(lineas, "\n")
}

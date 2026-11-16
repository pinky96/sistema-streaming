// Package catalogo gestiona el catálogo de contenidos del sistema de
// streaming: altas, búsquedas, filtros y listado.
//
// Cada contenido se representa con un map[string]string con las claves
// "titulo", "genero" y "anio". El catálogo es un slice de esos mapas.
package catalogo

import (
	"errors"
	"fmt"
	"strings"
)

// AgregarContenido agrega una película o serie al catálogo.
// Retorna un nuevo catálogo (no modifica el original) y un error si el
// título está vacío o ya existe.
func AgregarContenido(cat []map[string]string, titulo string, genero string, anio string) ([]map[string]string, error) {
	titulo = strings.TrimSpace(titulo)
	if titulo == "" {
		return cat, errors.New("el título no puede estar vacío")
	}
	for _, c := range cat {
		if strings.EqualFold(c["titulo"], titulo) {
			return cat, fmt.Errorf("%q ya existe en el catálogo", titulo)
		}
	}

	nuevo := map[string]string{
		"titulo": titulo,
		"genero": strings.TrimSpace(genero),
		"anio":   strings.TrimSpace(anio),
	}
	// append sobre una copia para no alterar el slice original.
	resultado := make([]map[string]string, len(cat), len(cat)+1)
	copy(resultado, cat)
	return append(resultado, nuevo), nil
}

// BuscarContenido busca un contenido por título (sin distinguir
// mayúsculas). Retorna el contenido y un error si no se encuentra.
func BuscarContenido(cat []map[string]string, titulo string) (map[string]string, error) {
	for _, c := range cat {
		if strings.EqualFold(c["titulo"], strings.TrimSpace(titulo)) {
			return c, nil
		}
	}
	return nil, fmt.Errorf("no se encontró %q en el catálogo", titulo)
}

// FiltrarPorGenero retorna un nuevo slice solo con los contenidos del
// género indicado.
func FiltrarPorGenero(cat []map[string]string, genero string) []map[string]string {
	resultado := []map[string]string{}
	for _, c := range cat {
		if strings.EqualFold(c["genero"], strings.TrimSpace(genero)) {
			resultado = append(resultado, c)
		}
	}
	return resultado
}

// ListarCatalogo retorna una representación en texto de todo el catálogo,
// lista para imprimir en consola.
func ListarCatalogo(cat []map[string]string) string {
	if len(cat) == 0 {
		return "El catálogo está vacío."
	}
	lineas := make([]string, 0, len(cat))
	for i, c := range cat {
		lineas = append(lineas, fmt.Sprintf("%d. %s (%s, %s)", i+1, c["titulo"], c["genero"], c["anio"]))
	}
	return strings.Join(lineas, "\n")
}

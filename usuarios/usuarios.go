// Package usuarios gestiona el registro, la autenticación y el listado
// de las personas usuarias del sistema de streaming.
//
// Los usuarios se representan con un map[string]string donde la clave es
// el correo electrónico y el valor es el nombre de la persona.
package usuarios

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// RegistrarUsuario agrega una nueva persona usuaria al sistema.
// Retorna un nuevo mapa de usuarios (no modifica el original) y un error
// si el correo ya está registrado o si algún dato está vacío.
func RegistrarUsuario(usuarios map[string]string, nombre string, correo string) (map[string]string, error) {
	nombre = strings.TrimSpace(nombre)
	correo = strings.TrimSpace(strings.ToLower(correo))

	if nombre == "" || correo == "" {
		return usuarios, errors.New("el nombre y el correo no pueden estar vacíos")
	}
	if _, existe := usuarios[correo]; existe {
		return usuarios, fmt.Errorf("el correo %s ya está registrado", correo)
	}

	// Copia del mapa original: la función no produce efectos secundarios.
	resultado := make(map[string]string, len(usuarios)+1)
	for k, v := range usuarios {
		resultado[k] = v
	}
	resultado[correo] = nombre
	return resultado, nil
}

// IniciarSesion verifica que el correo exista en el sistema.
// Retorna el nombre de la persona usuaria y un error si no está registrada.
func IniciarSesion(usuarios map[string]string, correo string) (string, error) {
	correo = strings.TrimSpace(strings.ToLower(correo))
	nombre, existe := usuarios[correo]
	if !existe {
		return "", fmt.Errorf("no existe una cuenta con el correo %s", correo)
	}
	return nombre, nil
}

// ListarUsuarios retorna los nombres registrados, ordenados alfabéticamente.
func ListarUsuarios(usuarios map[string]string) []string {
	nombres := make([]string, 0, len(usuarios))
	for _, nombre := range usuarios {
		nombres = append(nombres, nombre)
	}
	sort.Strings(nombres)
	return nombres
}

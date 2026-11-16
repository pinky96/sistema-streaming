// Sistema de Gestión de Streaming — Aprendizaje Autónomo 1
// Programación Orientada a Objetos (Etapa 1: programación funcional)
//
// Punto de entrada del sistema: demuestra el flujo completo llamando a
// las funciones exportadas de cada paquete (usuarios, catalogo y
// reproduccion).
package main

import (
	"fmt"

	"github.com/pinky96/sistema-streaming/catalogo"
	"github.com/pinky96/sistema-streaming/reproduccion"
	"github.com/pinky96/sistema-streaming/usuarios"
)

func main() {
	fmt.Println("=== Sistema de Gestión de Streaming ===")

	// ---- Módulo de usuarios -------------------------------------------
	cuentas := map[string]string{}

	cuentas, err := usuarios.RegistrarUsuario(cuentas, "Ana Quito", "ana@correo.com")
	if err != nil {
		fmt.Println("Error:", err)
	}
	cuentas, err = usuarios.RegistrarUsuario(cuentas, "Luis Ríos", "luis@correo.com")
	if err != nil {
		fmt.Println("Error:", err)
	}

	nombre, err := usuarios.IniciarSesion(cuentas, "ana@correo.com")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nSesión iniciada: bienvenida, %s\n", nombre)
	}

	fmt.Println("\nUsuarios registrados:")
	for _, n := range usuarios.ListarUsuarios(cuentas) {
		fmt.Println(" -", n)
	}

	// ---- Módulo de catálogo -------------------------------------------
	cat := []map[string]string{}

	cat, _ = catalogo.AgregarContenido(cat, "La Nube Andina", "Documental", "2024")
	cat, _ = catalogo.AgregarContenido(cat, "Río Arriba", "Drama", "2023")
	cat, _ = catalogo.AgregarContenido(cat, "Código Abierto", "Documental", "2026")

	fmt.Println("\nCatálogo completo:")
	fmt.Println(catalogo.ListarCatalogo(cat))

	fmt.Println("\nDocumentales disponibles:")
	fmt.Println(catalogo.ListarCatalogo(catalogo.FiltrarPorGenero(cat, "Documental")))

	encontrado, err := catalogo.BuscarContenido(cat, "río arriba")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("\nBúsqueda: %s (%s, %s)\n", encontrado["titulo"], encontrado["genero"], encontrado["anio"])
	}

	// ---- Módulo de reproducción ---------------------------------------
	historial := []string{}
	favoritos := []string{}

	historial, mensaje, err := reproduccion.ReproducirContenido(historial, "La Nube Andina")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("\n" + mensaje)
	}
	historial, mensaje, _ = reproduccion.ReproducirContenido(historial, "Río Arriba")
	fmt.Println(mensaje)

	favoritos, err = reproduccion.AgregarAFavoritos(favoritos, "La Nube Andina")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("\nFavoritos:", favoritos)

	fmt.Println("\nHistorial de reproducción:")
	fmt.Println(reproduccion.VerHistorial(historial))
}

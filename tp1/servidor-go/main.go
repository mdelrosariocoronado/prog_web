package main

import (
	"fmt"
	"net/http"
)

// serveForm: Maneja GET / para mostrar el formulario físico de la carpeta static
func serveForm(w http.ResponseWriter, r *http.Request) {
	// Evitamos que la ruta raíz "/" actúe como comodín
	if r.URL.Path != "/" || r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	
	// SOLUCIÓN: ServeFile lee el archivo real y lo envía configurando el Content-Type correcto automáticamente
	http.ServeFile(w, r, "./static/formularios.html")
}

// handleLogin: Maneja el envío de datos por POST desde el formulario
func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Validamos que sea estrictamente POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// 1. Parsear los datos del formulario (¡Crucial!)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al parsear el formulario", http.StatusBadRequest)
		return
	}

	// 2. Obtener el valor del campo 'user' (verifica que tu HTML tenga <input name="user">)
	username := r.FormValue("nombre")

	// Si no enviaron nada, le damos un nombre por defecto
	if username == "" {
		username = "Invitado"
	}

	// 3. Generar y enviar la respuesta HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!DOCTYPE html>
		<html>
		<head>
			<title>Bienvenido</title>
		</head> 
		<body>
			<h1>¡Hola, %s!</h1>
			<p>Recibimos tus datos correctamente por método POST.</p> 
			<a href="/">Volver al formulario</a>
		</body>
		</html>`, username)
		}

// Único punto de entrada del programa
func main() {
	// Registramos los endpoints asociados a sus funciones manejadoras
	http.HandleFunc("/", serveForm)
	http.HandleFunc("/formulario", handleLogin)

	port := ":8080"
	fmt.Printf("Servidor ejecutándose en http://localhost%s\n", port)

	// Iniciamos el servidor de escucha
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}